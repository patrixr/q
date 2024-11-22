package q

import (
	"context"
	"sync"
)

type EventHandler func(ev string, data any) error

type EventRef int

type Eventful interface {
	On(ev string, cb EventHandler) EventRef
	Once(ev string, cb EventHandler) EventRef
	Off(ev string, ref EventRef)
	Fire(ev string, data any) (bool, []error)
}

type callbackRef struct {
	ref      EventRef
	fn       EventHandler
	multiple bool
}

type EventEmitter struct {
	callbacks   map[string][]callbackRef
	concurrency int
	ctx         context.Context
	ref         EventRef
	mu          sync.Mutex
}

func NewEventEmitter(ctx context.Context, concurrency int) Eventful {
	return &EventEmitter{
		concurrency: concurrency,
		ctx:         ctx,
		callbacks:   map[string][]callbackRef{},
	}
}

func (emitter *EventEmitter) On(ev string, fn EventHandler) EventRef {
	return emitter.AddEventHandler(ev, fn, true)
}

func (emitter *EventEmitter) Once(ev string, fn EventHandler) EventRef {
	return emitter.AddEventHandler(ev, fn, false)
}

func (emitter *EventEmitter) AddEventHandler(ev string, fn EventHandler, multiple bool) EventRef {
	emitter.mu.Lock()

	defer emitter.mu.Unlock()

	ref := emitter.ref

	_, ok := emitter.callbacks[ev]

	if !ok {
		emitter.callbacks[ev] = []callbackRef{}
	}

	emitter.callbacks[ev] = append(emitter.callbacks[ev], callbackRef{ref, fn, multiple})
	emitter.ref += 1
	return ref
}

func (emitter *EventEmitter) Off(ev string, ref EventRef) {
	emitter.mu.Lock()

	defer emitter.mu.Unlock()

	_, ok := emitter.callbacks[ev]

	if !ok {
		return
	}

	emitter.callbacks[ev] = Filter(emitter.callbacks[ev], func(cbr callbackRef) bool {
		return cbr.ref != ref
	})
}

func (emitter *EventEmitter) Fire(ev string, data any) (bool, []error) {
	emitter.mu.Lock()

	if emitter.ctx.Err() != nil {
		// cancelled context
		return false, []error{emitter.ctx.Err()}
	}

	wg := sync.WaitGroup{}

	handlers, ok := emitter.callbacks[ev]

	if !ok || len(handlers) == 0 {
		return true, []error{}
	}

	results := make(chan error, len(handlers))
	jobs := make(chan EventHandler, len(handlers))

	emitter.callbacks[ev] = Filter(handlers, func(h callbackRef) bool {
		jobs <- h.fn // queue all the jobs
		return h.multiple
	})

	emitter.mu.Unlock()

	close(jobs)

	wg.Add(emitter.concurrency)

	for i := 0; i < emitter.concurrency; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-emitter.ctx.Done():
					return
				case fn, ok := <-jobs:
					if !ok {
						return
					}

					results <- fn(ev, data)
				}
			}
		}()
	}

	wg.Wait()

	close(results)

	errors := []error{}

	for err := range results {
		if err != nil {
			errors = append(errors, err)
		}
	}

	return len(errors) == 0, errors
}
