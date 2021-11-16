package flusher

import (
	"log"

	"github.com/Shemistan/note-service-api/internal/app/api"
	"github.com/Shemistan/note-service-api/internal/repo"
	mocksRepo "github.com/Shemistan/note-service-api/internal/repo/mocks"
	"github.com/Shemistan/note-service-api/internal/utills"
)

type Flusher interface {
	Flush(note []api.Note, batchSize int64) ([]api.Note, error)
}

type flusher struct {
	repo repo.Repo
}

func NewFlusher(repo *mocksRepo.MockRepo) Flusher {
	return &flusher{repo}
}

func (f *flusher) Flush(notes []api.Note, batchSize int64) ([]api.Note, error) {
	batches, err := utills.SplitSlice(notes, batchSize)
	if err != nil {
		log.Printf("failed to spliting slice: %s", err.Error())
		return nil, err
	}

	for i, batch := range batches {
		num, errAdd := f.repo.MultiAdd(batch)
		if errAdd != nil {
			log.Printf("failed to add slice: %s", errAdd.Error())

			var save = make([]api.Note, 0)
			for _, v := range batches[i:] {
				save = append(save, v...)
			}

			return save, errAdd
		}

		log.Printf("%d notes added", num)
	}

	return nil, nil
}
