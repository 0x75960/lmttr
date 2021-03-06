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

// Start of a process
func (l *Limitter) Start() {
	l.semaphore <- true
	l.wait.Add(1)
}

// End of a process
func (l *Limitter) End() {
	l.wait.Done()
	<-l.semaphore
}

// Wait to finish all of processes started
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

			for i := 0; i < len(s); i++ {
				select {
				case <-s:
					// take value
				default:
					break
				}
				<-time.After(3 * time.Millisecond)
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

// MaxLimitter
// be careful to deadlock
type MaxLimitter struct {
	semaphore chan bool
	wait      sync.WaitGroup
	max       uint
}

// NewMaxLimitter limit proc to max
func NewMaxLimitter(max uint) (ml MaxLimitter, err error) {
	return MaxLimitter{
		semaphore: make(chan bool, max),
		wait:      sync.WaitGroup{},
		max:       max,
	}, nil
}

// Start of a process
func (ml *MaxLimitter) Start() {
	ml.semaphore <- true
	ml.wait.Add(1)
}

// End of a process
func (ml *MaxLimitter) End() {
	ml.wait.Done()
}

// Flush limit once
func (ml *MaxLimitter) Flush() {
	for i := 0; i < len(ml.semaphore); i++ {
		select {
		case <-ml.semaphore:
		default:
			return
		}
		<-time.After(3 * time.Millisecond)
	}
}

// Mitigate limit once
func (ml *MaxLimitter) Mitigate(n int) {
	for i := 0; i < n; i++ {
		select {
		case <-ml.semaphore:
		default:
			return
		}
		<-time.After(3 * time.Millisecond)
	}
}

// IsMax check reaches to max
func (ml *MaxLimitter) IsMax() bool {
	return ml.max == uint(len(ml.semaphore))
}

// Wait to finish all of processes started
func (ml *MaxLimitter) Wait() {
	ml.wait.Wait()
}
