package q

import (
	"context"
	// "errors"
	// "sync"
	"testing"
	// "time"

	"github.com/stretchr/testify/assert"
)

func TestEventEmitter_On(t *testing.T) {

	t.Run("test event refs are all different", func(t *testing.T) {
		emitter := NewEventEmitter(context.Background(), 1).(*EventEmitter)
		ref := emitter.On("test", func(ev string, data any) error { return nil })
		assert.Equal(t, EventRef(0), ref, "Expected ref 0")

		ref2 := emitter.On("test", func(ev string, data any) error { return nil })
		assert.Equal(t, EventRef(1), ref2, "Expected ref 1")

		assert.Len(t, emitter.callbacks["test"], 2, "Expected 2 callbacks")
	})

	t.Run("test handler is called", func(t *testing.T) {
		emitter := NewEventEmitter(context.Background(), 1).(*EventEmitter)
		called := false

		emitter.On("test", func(ev string, data any) error {
			called = true
			return nil
		})

		emitter.Fire("test", nil)
		assert.True(t, called, "Event did not fire as intended")
	})

	t.Run("test firing event with no handlers", func(t *testing.T) {
		emitter := NewEventEmitter(context.Background(), 1).(*EventEmitter)

		ok, errs := emitter.Fire("test", nil)
		assert.Empty(t, errs)
		assert.True(t, ok)
		ok, errs = emitter.Fire("test", nil)
		assert.Empty(t, errs)
		assert.True(t, ok)
	})

	t.Run("test removing handler", func(t *testing.T) {
		emitter := NewEventEmitter(context.Background(), 1).(*EventEmitter)
		count := 0

		ref := emitter.On("test", func(ev string, data any) error {
			count += 1
			return nil
		})

		emitter.Fire("test", nil)
		assert.Equal(t, 1, count, "Event did not fire as intended")

		emitter.Off("test", ref)

		emitter.Fire("test", nil)
		assert.Equal(t, 1, count, "Callback called after it was removed")
	})

	t.Run("test multiple calls", func(t *testing.T) {
		count := 0
		emitter := NewEventEmitter(context.Background(), 2).(*EventEmitter)

		emitter.On("test", func(ev string, data any) error {
			count += 1
			return nil
		})

		emitter.On("test", func(ev string, data any) error {
			count += 1
			return nil
		})

		emitter.On("test_other", func(ev string, data any) error {
			count += 1
			return nil
		})

		ok, errs := emitter.Fire("test", "some data")
		assert.Empty(t, errs)
		assert.True(t, ok)

		ok, errs = emitter.Fire("test_other", "some data")
		assert.Empty(t, errs)
		assert.True(t, ok)

		assert.Equal(t, 3, count, "Not all handlers were called")
	})

	t.Run("test data propagation", func(t *testing.T) {
		dataReceived := ""
		otherDataReceived := ""
		emitter := NewEventEmitter(context.Background(), 2).(*EventEmitter)

		emitter.On("test", func(ev string, data any) error {
			str, ok := data.(string)
			assert.True(t, ok, "Bad data type received")
			dataReceived = str
			return nil
		})

		emitter.On("other test", func(ev string, data any) error {
			str, ok := data.(string)
			assert.True(t, ok, "Bad data type received")
			otherDataReceived = str
			return nil
		})

		ok, errs := emitter.Fire("test", "some data")
		assert.Empty(t, errs)
		assert.True(t, ok)

		ok, errs = emitter.Fire("other test", "some other data")
		assert.Empty(t, errs)
		assert.True(t, ok)

		assert.Equal(t, "some data", dataReceived, "Data was not propagated properly")
		assert.Equal(t, "some other data", otherDataReceived, "Data was not propagated properly")
	})

	t.Run("test cancellation of context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		emitter := NewEventEmitter(ctx, 1).(*EventEmitter)
		count := 0

		emitter.On("test", func(ev string, data any) error {
			count += 1
			return nil
		})

		emitter.Fire("test", nil)
		assert.Equal(t, 1, count, "Event did not fire as intended")

		cancel()

		emitter.Fire("test", nil)
		assert.Equal(t, 1, count, "Callback called after the context was cancelled")
	})
}

// func TestEventEmitter_Off(t *testing.T) {
// 	emitter := NewEventEmitter(context.Background(), 1).(*EventEmitter)
// 	ref := emitter.On("test", func(ev string, data any) error { return nil })
// 	emitter.On("test", func(ev string, data any) error { return nil })

// 	emitter.Off("test", ref)

// 	assert.Len(t, emitter.callbacks["test"], 1, "Expected 1 callback")

// 	emitter.Off("nonexistent", 10) // Test removing from a non-existent event – no assertion needed here as it's just checking for no panic
// }

// // Test q.Filter function (used internally by EventEmitter.Off)
// func TestQFilter(t *testing.T) {
// 	nums := []int{1, 2, 3, 4, 5}
// 	even := Filter(nums, func(n int) bool { return n%2 == 0 })
// 	assert.Equal(t, []int{2, 4}, even)

// 	empty := Filter(nums, func(n int) bool { return n > 5 })
// 	assert.Empty(t, empty)
// }
