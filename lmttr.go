package lmttr

import (
	"sync"
)

// Limitter object
type Limitter struct {
	semaphore chan bool
	wait      sync.WaitGroup
}

// NewLimitter object
func NewLimitter(concurrency uint) (lmttr Limitter, err error) {
	return Limitter{
		semaphore: make(chan bool, concurrency),
		wait:      sync.WaitGroup{},
	}, nil
}

// Start a process
func (l *Limitter) Start() {
	l.semaphore <- true
	l.wait.Add(1)
}

// End of a process
func (l *Limitter) End() {
	<-l.semaphore
	l.wait.Done()
}

// Wait all processes
func (l *Limitter) Wait() {
	l.wait.Wait()
}
