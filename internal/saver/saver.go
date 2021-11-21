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
	capacity        int64
	bachSize        int64
	flusher         flusher.Flusher
	alarmer         alarmer.Alarmer
	notes           []api.Note
	notesChan       chan api.Note
	end             chan struct{}
	lossAllData     bool
	initInitialized bool
}

func NewSaver(capacity, bachSize int64, flusher flusher.Flusher, alarmer alarmer.Alarmer, lossAllData bool) Saver {
	return &saver{
		capacity:        capacity,
		bachSize:        bachSize,
		flusher:         flusher,
		alarmer:         alarmer,
		lossAllData:     lossAllData,
		notes:           make([]api.Note, 0),
		notesChan:       make(chan api.Note),
		end:             make(chan struct{}),
		initInitialized: false,
	}
}

func (s *saver) Init() error {
	if s.initInitialized {
		return errors.New("failed to repeated initialized saver")
	}
	err := s.alarmer.Init()
	if err != nil {
		return err
	}

	s.initInitialized = true

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

	return nil
}

func (s *saver) Save(note api.Note) error {
	if !s.initInitialized {
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
	res, err := s.flusher.Flush(s.notes, s.bachSize)

	if err != nil {
		log.Printf("failed to flushed data %s", err.Error())
	}
	s.notes = s.notes[:copy(s.notes, res)]
}
