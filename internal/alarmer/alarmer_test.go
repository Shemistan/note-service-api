package alarmer

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAlarmer_Alarm(t *testing.T) {

	maxRes := 6
	minRes := 4
	num := 0

	alarmer, errNewAlarm := NewAlarmer(1 * time.Second)
	if errNewAlarm != nil {
		log.Printf("fail to crate new alarmer")
	}

	errInitAlarm := alarmer.Init()
	if errInitAlarm != nil {
		log.Printf("fail to initialized alarmer")
	}

	timer := time.NewTimer(5 * time.Second)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-alarmer.Alarm():
				num++
			case <-timer.C:
				return
			}
		}
	}()
	wg.Wait()

	success := false
	if (num >= minRes) && (num <= maxRes) {
		success = true
	}

	require.Nil(t, errInitAlarm)
	require.Nil(t, errNewAlarm)
	require.True(t, success)
}
