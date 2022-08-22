package mutex

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrTimeout = errors.New("lock timeout")
)

type ITimeoutLock interface {
	Lock(duration time.Duration) error
	Unlock()
}

type timeoutLock struct {
	*sync.Mutex

	islocked bool
	// trueLock *sync.Mutex
}

func NewTimeoutLock() ITimeoutLock {
	return &timeoutLock{
		Mutex:    &sync.Mutex{},
		islocked: false,
		// trueLock: &sync.Mutex{},
	}
}

func (l *timeoutLock) Lock(duration time.Duration) error {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	step := duration.Milliseconds() / 100
	for {
		select {
		case <-timer.C:
			return ErrTimeout
		default:
			if l.tryLock() {
				return nil
			}
			time.Sleep(time.Duration(step) * time.Millisecond)
		}
	}
}

func (l *timeoutLock) tryLock() bool {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	if l.islocked {
		return false
	}

	l.islocked = true
	// l.trueLock.Lock()
	return true
}

func (l *timeoutLock) Unlock() {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	// l.trueLock.Unlock()
	l.islocked = false
}
