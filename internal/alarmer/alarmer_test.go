package alarmer

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAlarmer_Alarm(t *testing.T) {
	var alarm Alarmer
	maxExpRes := 6
	minExpRes := 4
	num := 0

	alarm = NewAlarmer(1 * time.Second)
	err := alarm.Init()
	timer := time.NewTimer(5 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-alarm.Alarm():
				num++
			case <-timer.C:

				return
			}
		}
	}()
	wg.Wait()

	res := false
	if num >= minExpRes && num <= maxExpRes {
		res = true
	}

	require.Nil(t, err)
	require.True(t, res)
}
