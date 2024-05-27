package q

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
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
	errorHanlder func(J, error)
	workers      int
	waitGroup    sync.WaitGroup
	cancel       context.CancelFunc
}

// NewBackgroundJob creates a new background job processor
// with the given executor function, error handler, and concurrency.
// The executor function is called for each job in the queue.
// It uses Go routines to process the jobs concurrently.
func NewBackgroundJob[J any](
	ctx context.Context,
	executor func(J) error,
	errorHandler func(J, error),
	concurrency int,
) JobProcessor[J] {
	channel := make(chan J, 100)

	localCtx, cancel := context.WithCancel(ctx)

	job := &BackgroundJob[J]{
		queue:        channel,
		executor:     executor,
		workers:      concurrency,
		errorHanlder: errorHandler,
		cancel:       cancel,
	}

	job.waitGroup.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer job.waitGroup.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case <-localCtx.Done():
					return
				case j, ok := <-channel:
					if !ok {
						return
					}
					err := job.executor(j)
					if err != nil {
						errorHandler(j, err)
					}
				default:
					time.Sleep(100 * time.Millisecond)
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
	job.cancel()
	job.waitGroup.Wait()
}
