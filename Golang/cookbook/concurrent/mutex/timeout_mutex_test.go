package mutex

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetLock(t *testing.T) {
	lock := NewTimeoutLock()
	require.Nil(t, lock.Lock(1*time.Second))
	require.Equal(t, true, errors.Is(lock.Lock(1*time.Second), ErrTimeout))
	lock.Unlock()
	require.Nil(t, lock.Lock(1*time.Second))
	lock.Unlock()
}

func TestUnlockMultiGoroutine(t *testing.T) {
	wait := sync.WaitGroup{}
	wait.Add(2)

	lock := NewTimeoutLock()

	lockFlag := make(chan struct{}, 1)

	go func() {
		defer wait.Done()
		lock.Lock(1 * time.Second)
		lockFlag <- struct{}{}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
			wait.Done()
		}()
		<-lockFlag
		lock.Unlock()
	}()

	wait.Wait()
}
