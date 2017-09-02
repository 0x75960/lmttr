package lmttr

import (
	"sync"
	"time"
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

// PeriodLimitter provides limit in period
type PeriodLimitter struct {
	semaphore chan bool
	wait      sync.WaitGroup
	period    *time.Ticker
}

// NewPeriodLimitter from concurrency and span
func NewPeriodLimitter(concurrency uint, span time.Duration) (p PeriodLimitter, err error) {

	s := make(chan bool, concurrency)
	ticker := time.NewTicker(span)

	go func() {

		for range ticker.C {

			for {

				select {
				case <-s:
					// take value while it possible
				default:
					break
				}

			}

		}

	}()

	return PeriodLimitter{
		semaphore: s,
		wait:      sync.WaitGroup{},
		period:    ticker,
	}, nil
}

// Start of a process
func (pl *PeriodLimitter) Start() {
	pl.semaphore <- true
	pl.wait.Add(1)
}

// End of a process
func (pl *PeriodLimitter) End() {
	pl.wait.Done()
}

// Wait to finish all of processes started
func (pl *PeriodLimitter) Wait() {
	pl.wait.Wait()
}
