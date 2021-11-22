package saver

import (
	"errors"
	"github.com/Shemistan/note-service-api/internal/alarmer"
	"github.com/Shemistan/note-service-api/internal/app/api"
	"github.com/Shemistan/note-service-api/internal/flusher"
	"log"
)

// Saver ...
type Saver interface {
	Save(note api.Note) error
	Init() error
	Close()
}

type saver struct {
	capacity    int64
	batchSize   int64
	flusher     flusher.Flusher
	alarmer     alarmer.Alarmer
	notes       []api.Note
	notesChan   chan api.Note
	end         chan struct{}
	lossAllData bool
	initInit    bool
}

func NewSaver(capacity, batchSize int64, flusher flusher.Flusher, alarmer alarmer.Alarmer, lossAllData bool) (Saver, error) {
	if capacity <= 0 {
		return nil, errors.New("failed to capacity value")
	}

	if batchSize <= 0 {
		return nil, errors.New("failed to batch size value")
	}

	return &saver{
		capacity:    capacity,
		batchSize:   batchSize,
		flusher:     flusher,
		alarmer:     alarmer,
		lossAllData: lossAllData,
		notes:       []api.Note{},
		notesChan:   make(chan api.Note),
		end:         make(chan struct{}),
		initInit:    false,
	}, nil
}

func (s *saver) Init() error {
	if s.initInit {
		return errors.New("the saver has already been initialized")
	}

	go func() {
		for {
			select {
			case note := <-s.notesChan:
				s.saveToBuffer(note)
			case _, ok := <-s.alarmer.Alarm():
				if ok {
					s.flushData()
				}
			case <-s.end:
				return
			}
		}
	}()

	s.initInit = true

	return nil
}

func (s *saver) Save(note api.Note) error {
	if !s.initInit {
		return errors.New("failed to initialized saver")
	}

	s.notesChan <- note

	return nil
}

func (s *saver) Close() {
	s.end <- struct{}{}
	s.flushData()

	s.alarmer.Close()
	close(s.notesChan)
	close(s.end)
}

func (s *saver) saveToBuffer(note api.Note) {
	if int64(len(s.notes)) >= s.capacity {
		if s.lossAllData || s.capacity <= 1 {
			s.notes = s.notes[:0]
		} else {
			s.notes = s.notes[1:]
		}
	}

	s.notes = append(s.notes, note)
}

func (s *saver) flushData() {
	if len(s.notes) <= 0 {
		return
	}

	res, err := s.flusher.Flush(s.notes, s.batchSize)
	if err != nil {
		log.Printf("failed to flushed data %s", err.Error())
	}

	s.notes = s.notes[:copy(s.notes, res)]
}
