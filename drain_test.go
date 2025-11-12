package q

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDrainChannel_EmptyChannel(t *testing.T) {
	ch := make(chan int)
	close(ch)

	DrainChannel(ch)
}

func TestDrainChannel_ChannelWithValues(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	DrainChannel(ch)

	select {
	case _, ok := <-ch:
		assert.False(t, ok, "Channel should be closed and empty")
	default:
		t.Fatal("Channel should be closed, not blocking")
	}
}

func TestDrainChannel_BufferedChannelNotClosed(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	DrainChannel(ch)

	select {
	case <-ch:
		t.Fatal("Channel should be empty after draining")
	default:
	}

	select {
	case ch <- 4:
	default:
		t.Fatal("Should be able to send to open channel")
	}
}

func TestDrainChannel_UnbufferedChannelNotClosed(t *testing.T) {
	ch := make(chan int)

	DrainChannel(ch)

	go func() {
		ch <- 42
	}()

	select {
	case val := <-ch:
		assert.Equal(t, 42, val)
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Should have received value from channel")
	}
}

func TestDrainChannel_MixedScenario(t *testing.T) {
	ch := make(chan string, 2)
	ch <- "hello"

	DrainChannel(ch)

	select {
	case <-ch:
		t.Fatal("Channel should be empty")
	default:
	}

	ch <- "world"
	ch <- "test"
	close(ch)

	DrainChannel(ch)

	select {
	case _, ok := <-ch:
		assert.False(t, ok, "Channel should be closed and empty")
	default:
		t.Fatal("Channel should be closed, not blocking")
	}
}

func TestDrainChannel_GenericTypes(t *testing.T) {
	strCh := make(chan string, 2)
	strCh <- "test1"
	strCh <- "test2"
	close(strCh)
	DrainChannel(strCh)

	type testStruct struct {
		ID   int
		Name string
	}
	structCh := make(chan testStruct, 1)
	structCh <- testStruct{ID: 1, Name: "test"}
	close(structCh)
	DrainChannel(structCh)

	ptrCh := make(chan *int, 1)
	val := 42
	ptrCh <- &val
	close(ptrCh)
	DrainChannel(ptrCh)
}

func TestDrainChannel_LargeBuffer(t *testing.T) {
	const size = 1000
	ch := make(chan int, size)

	for i := 0; i < size; i++ {
		ch <- i
	}
	close(ch)

	start := time.Now()
	DrainChannel(ch)
	duration := time.Since(start)

	select {
	case _, ok := <-ch:
		assert.False(t, ok, "Channel should be closed and empty")
	default:
		t.Fatal("Channel should be closed, not blocking")
	}

	assert.Less(t, duration, time.Second, "DrainChannel should complete quickly for large buffers")
}

func TestDrainChannel_NilValues(t *testing.T) {
	ch := make(chan *string, 2)
	ch <- nil
	str := "test"
	ch <- &str
	close(ch)

	DrainChannel(ch)

	select {
	case _, ok := <-ch:
		assert.False(t, ok, "Channel should be closed and empty")
	default:
		t.Fatal("Channel should be closed, not blocking")
	}
}
