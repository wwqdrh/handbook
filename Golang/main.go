package main

import "sync"

func main() {
	l := sync.RWMutex{}
	l.Lock()
	defer l.Unlock()

	a := func() {
		l.RLock()
		defer l.Unlock()
	}

	a()
}
