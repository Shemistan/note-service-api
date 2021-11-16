package repo

//go:generate mockgen -destination=mocks/mock_repo.go -package=mocks . Repo

import (
	"github.com/Shemistan/note-service-api/internal/app/api"
)

type Repo interface {
	Add(note api.Note) error
	MultiAdd(notes []api.Note) (int64, error)
	Update(id int64, note api.Note) error
	Remove(id int64) error
	Describe(id int64) (api.Note, error)
}
