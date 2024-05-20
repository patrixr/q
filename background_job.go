package q

import (
	"sync"
	"sync/atomic"
)

type JobProcessor[T any] interface {
	PerformLater(T)
	PerformNow(T) error
	Close()
}

type BackgroundJob[J any] struct {
	queue        chan J
	errors       atomic.Int32
	executor     func(J) error
	errorHanlder func(error)
	workers      int
	waitGroup    sync.WaitGroup
}

// NewBackgroundJob creates a new background job processor
// with the given executor function, error handler, and concurrency.
// The executor function is called for each job in the queue.
// It uses Go routines to process the jobs concurrently.
func NewBackgroundJob[J any](
	executor func(J) error,
	errorHandler func(error),
	concurrency int,
) JobProcessor[J] {
	channel := make(chan J, 100)

	job := &BackgroundJob[J]{
		queue:        channel,
		executor:     executor,
		workers:      concurrency,
		errorHanlder: errorHandler,
	}

	job.waitGroup.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer job.waitGroup.Done()
			for j := range channel {
				err := job.executor(j)
				if err != nil {
					job.errors.Add(1)
					if errorHandler != nil {
						errorHandler(err)
					}
				}
			}
		}()
	}

	return job
}

func (job *BackgroundJob[J]) PerformLater(j J) {
	job.queue <- j
}

func (job *BackgroundJob[J]) PerformNow(j J) error {
	return job.executor(j)
}

func (job *BackgroundJob[J]) Close() {
	close(job.queue)
	job.waitGroup.Wait()
}
