package flusher

import (
	"testing"

	"github.com/Shemistan/note-service-api/internal/app/api"
	mocksRepo "github.com/Shemistan/note-service-api/internal/repo/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestFlusher_Flush(t *testing.T) {
	var (
		mockCtrl = gomock.NewController(t)
	)

	var mockNoteRepo = mocksRepo.NewMockRepo(mockCtrl)
	noteFlusher := NewFlusher(mockNoteRepo)

	t.Run("input value equal nil", func(t *testing.T) {
		expectedRes := "error input values"

		_, err := noteFlusher.Flush(nil, 1)
		require.Error(t, err)
		require.Equal(t, expectedRes, err.Error())
	})

	t.Run("len input value equal zero", func(t *testing.T) {
		mockNoteRepo.EXPECT().MultiAdd(make([]api.Note, 0)).Return(int64(0), nil).Times(1)
		req := make([]api.Note, 0)
		var expectedRes []api.Note

		res, err := noteFlusher.Flush(req, 1)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("input one note", func(t *testing.T) {
		gomock.InOrder(mockNoteRepo.EXPECT().MultiAdd([]api.Note{
			{
				Id:          1,
				UserId:      2,
				ClassroomId: 3,
				DocumentId:  4,
			},
		},
		).Return(int64(0), nil).Times(1),
		)

		req := []api.Note{
			{
				Id:          1,
				UserId:      2,
				ClassroomId: 3,
				DocumentId:  4,
			},
		}
		var expectedRes []api.Note

		res, err := noteFlusher.Flush(req, 1)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("success case", func(t *testing.T) {
		gomock.InOrder(
			mockNoteRepo.EXPECT().MultiAdd([]api.Note{
				{
					Id:          1,
					UserId:      2,
					ClassroomId: 3,
					DocumentId:  4,
				},
			},
			).Return(int64(0), nil).Times(1),
			mockNoteRepo.EXPECT().MultiAdd([]api.Note{
				{
					Id:          5,
					UserId:      6,
					ClassroomId: 7,
					DocumentId:  8,
				},
			},
			).Return(int64(0), nil).Times(1),
		)

		req := []api.Note{
			{
				Id:          1,
				UserId:      2,
				ClassroomId: 3,
				DocumentId:  4,
			},
			{
				Id:          5,
				UserId:      6,
				ClassroomId: 7,
				DocumentId:  8,
			},
		}
		var expectedRes []api.Note

		res, err := noteFlusher.Flush(req, 1)
		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

}
