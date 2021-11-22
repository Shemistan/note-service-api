package saver

import (
	"github.com/Shemistan/note-service-api/internal/alarmer"
	"github.com/Shemistan/note-service-api/internal/app/api"
	"github.com/Shemistan/note-service-api/internal/flusher"
	mocksRepo "github.com/Shemistan/note-service-api/internal/repo/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

func TestSaver(t *testing.T) {
	var (
		mockCtrl = gomock.NewController(t)
	)
	var mockNoteRepo = mocksRepo.NewMockRepo(mockCtrl)

	req := []api.Note{
		{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
		{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
		{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		{Id: 11, UserId: 11, ClassroomId: 123, DocumentId: 16},
		{Id: 21, UserId: 21, ClassroomId: 124, DocumentId: 17},
		{Id: 31, UserId: 31, ClassroomId: 123, DocumentId: 16},
		{Id: 41, UserId: 41, ClassroomId: 124, DocumentId: 17},
	}

	t.Run("time alarmer equal time write in buffer", func(t *testing.T) {
		t.Run("input capacity and batch size equal one", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)

			noteFlusher := flusher.NewFlusher(mockNoteRepo)
			exmAlarmer, errNewAlarm := alarmer.NewAlarmer(10 * time.Millisecond)
			if errNewAlarm != nil {
				log.Printf("failed to crate new alarmer %s", errNewAlarm.Error())
			}

			exmSaver, errNewSaver := NewSaver(1, 1, noteFlusher, exmAlarmer, true)
			if errNewSaver != nil {
				log.Printf("failed to crate new saver %s", errNewSaver.Error())
			}

			errInitAlarm := exmAlarmer.Init()
			if errInitAlarm != nil {
				log.Printf("failed to initialized alarm: %s", errInitAlarm.Error())
				return
			}

			errInitSaver := exmSaver.Init()
			if errInitSaver != nil {
				log.Printf("failed to initialized saver: %s", errInitSaver.Error())
				return
			}
			defer exmSaver.Close()

			for _, val := range req {
				err := exmSaver.Save(val)
				if err != nil {
					log.Printf("failed to save %s", err.Error())
				}
				time.Sleep(10 * time.Millisecond)
			}

			require.Nil(t, errInitSaver)
			require.Nil(t, errNewSaver)
			require.Nil(t, errInitAlarm)
			require.Nil(t, errNewAlarm)
		})

		t.Run("input bach size less capacity", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)

			noteFlusher := flusher.NewFlusher(mockNoteRepo)
			exmAlarmer, errNewAlarm := alarmer.NewAlarmer(10 * time.Millisecond)
			if errNewAlarm != nil {
				log.Printf("failed to crate new alarmer %s", errNewAlarm.Error())
			}

			errInitAlarm := exmAlarmer.Init()
			if errInitAlarm != nil {
				log.Printf("failed to initialized alarm: %s", errInitAlarm.Error())
				return
			}

			exmSaver, errNewSaver := NewSaver(3, 1, noteFlusher, exmAlarmer, true)
			if errNewSaver != nil {
				log.Printf("failed to crate new saver %s", errNewSaver.Error())
			}

			errInitSaver := exmSaver.Init()
			if errInitSaver != nil {
				log.Printf("failed to initialized saver: %s", errInitSaver.Error())
				return
			}
			defer exmSaver.Close()

			for _, val := range req {
				err := exmSaver.Save(val)
				if err != nil {
					log.Printf("failed to save %s", err.Error())
				}
				time.Sleep(10 * time.Millisecond)
			}

			require.Nil(t, errInitSaver)
			require.Nil(t, errNewSaver)
			require.Nil(t, errInitAlarm)
			require.Nil(t, errNewAlarm)
		})

		t.Run("input capacity less bach size", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)

			noteFlusher := flusher.NewFlusher(mockNoteRepo)
			exmAlarmer, errNewAlarm := alarmer.NewAlarmer(10 * time.Millisecond)
			if errNewAlarm != nil {
				log.Printf("fail to crate new alarmer %s", errNewAlarm.Error())
			}

			exmSaver, errNewSaver := NewSaver(1, 3, noteFlusher, exmAlarmer, true)
			if errNewSaver != nil {
				log.Printf("fail to crate new saver %s", errNewSaver.Error())
			}

			errInitAlarm := exmAlarmer.Init()
			if errInitAlarm != nil {
				log.Printf("failed to initialized alarm: %s", errInitAlarm.Error())
				return
			}

			errInitSaver := exmSaver.Init()
			if errInitSaver != nil {
				log.Printf("failed to initialized saver: %s", errInitSaver.Error())
				return
			}
			defer exmSaver.Close()

			for _, val := range req {
				err := exmSaver.Save(val)
				if err != nil {
					log.Printf("failed to save %s", err.Error())
				}
				time.Sleep(10 * time.Millisecond)
			}

			require.Nil(t, errInitSaver)
			require.Nil(t, errNewSaver)
			require.Nil(t, errInitAlarm)
			require.Nil(t, errNewAlarm)
		})
	})

	t.Run("time alarmer less time write in buffer", func(t *testing.T) {
		t.Run("input capacity and batch size equal one", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)

			noteFlusher := flusher.NewFlusher(mockNoteRepo)
			exmAlarmer, errNewAlarm := alarmer.NewAlarmer(10 * time.Millisecond)
			if errNewAlarm != nil {
				log.Printf("fail to crate new alarmer %s", errNewAlarm.Error())
			}

			exmSaver, errNewSaver := NewSaver(1, 1, noteFlusher, exmAlarmer, true)
			if errNewSaver != nil {
				log.Printf("fail to crate new saver %s", errNewSaver.Error())
			}

			errInitAlarm := exmAlarmer.Init()
			if errInitAlarm != nil {
				log.Printf("failed to initialized alarm: %s", errInitAlarm.Error())
				return
			}

			errInitSaver := exmSaver.Init()
			if errInitSaver != nil {
				log.Printf("failed to initialized saver: %s", errInitSaver.Error())
				return
			}
			defer exmSaver.Close()

			for _, val := range req {
				err := exmSaver.Save(val)
				if err != nil {
					log.Printf("failed to save %s", err.Error())
				}
				time.Sleep(20 * time.Millisecond)
			}

			require.Nil(t, errInitSaver)
			require.Nil(t, errNewSaver)
			require.Nil(t, errInitAlarm)
			require.Nil(t, errNewAlarm)
		})

		t.Run("input bach size less capacity", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)

			noteFlusher := flusher.NewFlusher(mockNoteRepo)
			exmAlarmer, errNewAlarm := alarmer.NewAlarmer(10 * time.Millisecond)
			if errNewAlarm != nil {
				log.Printf("fail to crate new alarmer %s", errNewAlarm.Error())
			}

			exmSaver, errNewSaver := NewSaver(3, 1, noteFlusher, exmAlarmer, true)
			if errNewSaver != nil {
				log.Printf("fail to crate new saver %s", errNewSaver.Error())
			}

			errInitAlarm := exmAlarmer.Init()
			if errInitAlarm != nil {
				log.Printf("failed to initialized alarm: %s", errInitAlarm.Error())
				return
			}

			errInitSaver := exmSaver.Init()
			if errInitSaver != nil {
				log.Printf("failed to initialized saver: %s", errInitSaver.Error())
				return
			}
			defer exmSaver.Close()

			for _, val := range req {
				err := exmSaver.Save(val)
				if err != nil {
					log.Printf("failed to save %s", err.Error())
				}
				time.Sleep(20 * time.Millisecond)
			}

			require.Nil(t, errInitSaver)
			require.Nil(t, errNewSaver)
			require.Nil(t, errInitAlarm)
			require.Nil(t, errNewAlarm)
		})

		t.Run("input capacity less bach size", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)

			noteFlusher := flusher.NewFlusher(mockNoteRepo)
			exmAlarmer, errNewAlarm := alarmer.NewAlarmer(10 * time.Millisecond)
			if errNewAlarm != nil {
				log.Printf("fail to crate new alarmer %s", errNewAlarm.Error())
			}

			exmSaver, errNewSaver := NewSaver(1, 3, noteFlusher, exmAlarmer, true)
			if errNewSaver != nil {
				log.Printf("fail to crate new saver %s", errNewSaver.Error())
			}

			errInitAlarm := exmAlarmer.Init()
			if errInitAlarm != nil {
				log.Printf("failed to initialized alarm: %s", errInitAlarm.Error())
				return
			}

			errInitSaver := exmSaver.Init()
			if errInitSaver != nil {
				log.Printf("failed to initialized saver: %s", errInitSaver.Error())
				return
			}
			defer exmSaver.Close()

			for _, val := range req {
				err := exmSaver.Save(val)
				if err != nil {
					log.Printf("failed to save %s", err.Error())
				}
				time.Sleep(20 * time.Millisecond)
			}

			require.Nil(t, errInitSaver)
			require.Nil(t, errNewSaver)
			require.Nil(t, errInitAlarm)
			require.Nil(t, errNewAlarm)
		})
	})

	t.Run("time write in buffer less time alarmer", func(t *testing.T) {
		t.Run("input capacity and batch size equal one", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(5)

			noteFlusher := flusher.NewFlusher(mockNoteRepo)
			exmAlarmer, errNewAlarm := alarmer.NewAlarmer(20 * time.Millisecond)
			if errNewAlarm != nil {
				log.Printf("fail to crate new alarmer %s", errNewAlarm.Error())
			}

			exmSaver, errNewSaver := NewSaver(1, 1, noteFlusher, exmAlarmer, true)
			if errNewSaver != nil {
				log.Printf("fail to crate new saver %s", errNewSaver.Error())
			}

			errInitAlarm := exmAlarmer.Init()
			if errInitAlarm != nil {
				log.Printf("failed to initialized alarm: %s", errInitAlarm.Error())
				return
			}

			errInitSaver := exmSaver.Init()
			if errInitSaver != nil {
				log.Printf("failed to initialized saver: %s", errInitSaver.Error())
				return
			}
			defer exmSaver.Close()

			for _, val := range req {
				err := exmSaver.Save(val)
				if err != nil {
					log.Printf("failed to save %s", err.Error())
				}
				time.Sleep(10 * time.Millisecond)
			}

			require.Nil(t, errInitSaver)
			require.Nil(t, errNewSaver)
			require.Nil(t, errInitAlarm)
			require.Nil(t, errNewAlarm)
		})

		t.Run("input bach size less capacity", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(8)

			noteFlusher := flusher.NewFlusher(mockNoteRepo)
			exmAlarmer, errNewAlarm := alarmer.NewAlarmer(500 * time.Millisecond)
			if errNewAlarm != nil {
				log.Printf("fail to crate new alarmer %s", errNewAlarm.Error())
			}

			exmSaver, errNewSaver := NewSaver(3, 1, noteFlusher, exmAlarmer, true)
			if errNewSaver != nil {
				log.Printf("fail to crate new saver %s", errNewSaver.Error())
			}

			errInitAlarm := exmAlarmer.Init()
			if errInitAlarm != nil {
				log.Printf("failed to initialized alarm: %s", errInitAlarm.Error())
				return
			}

			errInitSaver := exmSaver.Init()
			if errInitAlarm != nil {
				log.Printf("failed to initialized saver: %s", errInitSaver.Error())
				return
			}
			defer exmSaver.Close()

			for _, val := range req {
				err := exmSaver.Save(val)
				if err != nil {
					log.Printf("failed to save %s", err.Error())
				}
				time.Sleep(250 * time.Millisecond)
			}

			require.Nil(t, errInitSaver)
			require.Nil(t, errNewSaver)
			require.Nil(t, errInitAlarm)
			require.Nil(t, errNewAlarm)
		})

		t.Run("input capacity less bach size", func(t *testing.T) {
			mockNoteRepo.EXPECT().MultiAdd(gomock.All()).Return(int64(0), nil).Times(2)

			noteFlusher := flusher.NewFlusher(mockNoteRepo)
			exmAlarmer, errNewAlarm := alarmer.NewAlarmer(500 * time.Millisecond)
			if errNewAlarm != nil {
				log.Printf("fail to crate new alarmer %s", errNewAlarm.Error())
			}

			exmSaver, errNewSaver := NewSaver(1, 3, noteFlusher, exmAlarmer, true)
			if errNewSaver != nil {
				log.Printf("fail to crate new saver %s", errNewSaver.Error())
			}

			errInitAlarm := exmAlarmer.Init()
			if errInitAlarm != nil {
				log.Printf("failed to initialized alarm: %s", errInitAlarm.Error())
				return
			}

			errInitSaver := exmSaver.Init()
			if errInitAlarm != nil {
				log.Printf("failed to initialized saver: %s", errInitSaver.Error())
				return
			}
			defer exmSaver.Close()

			for _, val := range req {
				err := exmSaver.Save(val)
				if err != nil {
					log.Printf("failed to save %s", err.Error())
				}
				time.Sleep(250 * time.Millisecond)
			}

			require.Nil(t, errInitSaver)
			require.Nil(t, errNewSaver)
			require.Nil(t, errInitAlarm)
			require.Nil(t, errNewAlarm)
		})
	})

	t.Run("input capacity equal zero", func(t *testing.T) {
		expectedError := "failed to capacity value"

		noteFlusher := flusher.NewFlusher(mockNoteRepo)
		exmAlarmer, errNewAlarm := alarmer.NewAlarmer(20 * time.Millisecond)
		if errNewAlarm != nil {
			log.Printf("fail to crate new alarmer %s", errNewAlarm.Error())
		}

		_, errNewSaver := NewSaver(0, 3, noteFlusher, exmAlarmer, true)
		if errNewSaver != nil {
			log.Printf("fail to crate new saver %s", errNewSaver.Error())
		}

		require.Nil(t, errNewAlarm)
		require.NotNil(t, errNewSaver)
		require.Equal(t, expectedError, errNewSaver.Error())
	})

	t.Run("input batch size equal zero", func(t *testing.T) {
		expectedError := "failed to batch size value"

		noteFlusher := flusher.NewFlusher(mockNoteRepo)
		exmAlarmer, errNewAlarm := alarmer.NewAlarmer(20 * time.Millisecond)
		if errNewAlarm != nil {
			log.Printf("fail to crate new alarmer %s", errNewAlarm.Error())
		}

		_, errNewSaver := NewSaver(3, 0, noteFlusher, exmAlarmer, true)
		if errNewSaver != nil {
			log.Printf("fail to crate new saver %s", errNewSaver.Error())
		}

		require.Nil(t, errNewAlarm)
		require.NotNil(t, errNewSaver)
		require.Equal(t, expectedError, errNewSaver.Error())
	})

}
