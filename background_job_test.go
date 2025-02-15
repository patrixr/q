package q

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestBackgroundJob_PerformLater(t *testing.T) {
	var processedCount atomic.Int32
	executor := func(job int) error {
		processedCount.Add(1)
		return nil
	}

	errorHandler := func(_ int, err error) {
		t.Errorf("Error handler should not be called: %v", err)
	}

	jobProcessor := NewBackgroundJob(context.Background(), executor, errorHandler, 3)
	defer jobProcessor.Close()

	for i := 0; i < 10; i++ {
		jobProcessor.PerformLater(i)
	}

	time.Sleep(1 * time.Second) // Allow some time for jobs to be processed

	if processedCount.Load() != 10 {
		t.Errorf("Expected 10 jobs to be processed, but got %d", processedCount.Load())
	}
}

func TestBackgroundJob_PerformNow(t *testing.T) {
	var processedCount atomic.Int32
	executor := func(job int) error {
		processedCount.Add(1)
		return nil
	}

	errorHandler := func(_ int, err error) {
		t.Errorf("Error handler should not be called: %v", err)
	}

	jobProcessor := NewBackgroundJob(context.Background(), executor, errorHandler, 3)
	defer jobProcessor.Close()

	for i := 0; i < 10; i++ {
		err := jobProcessor.PerformNow(i)
		if err != nil {
			t.Errorf("PerformNow returned an error: %v", err)
		}
	}

	if processedCount.Load() != 10 {
		t.Errorf("Expected 10 jobs to be processed, but got %d", processedCount.Load())
	}
}

func TestBackgroundJob_ErrorHandling(t *testing.T) {
	var processedCount atomic.Int32
	var errorCount atomic.Int32
	executor := func(job int) error {
		if job%2 == 0 {
			return errors.New("intentional error")
		}
		processedCount.Add(1)
		return nil
	}

	errorHandler := func(_ int, err error) {
		errorCount.Add(1)
	}

	jobProcessor := NewBackgroundJob(context.Background(), executor, errorHandler, 3)
	defer jobProcessor.Close()

	for i := 0; i < 10; i++ {
		jobProcessor.PerformLater(i)
	}

	time.Sleep(1 * time.Second) // Allow some time for jobs to be processed

	if processedCount.Load() != 5 {
		t.Errorf("Expected 5 successful jobs, but got %d", processedCount.Load())
	}

	if errorCount.Load() != 5 {
		t.Errorf("Expected 5 errors, but got %d", errorCount.Load())
	}
}

func TestBackgroundJob_Cancel_Context(t *testing.T) {
	t.Skip()

	var processedCount atomic.Int32
	executor := func(job int) error {
		processedCount.Add(1)
		return nil
	}

	errorHandler := func(_ int, err error) {
		t.Errorf("Error handler should not be called: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	jobProcessor := NewBackgroundJob(ctx, executor, errorHandler, 3)

	cancel() // Cancel the context

	for i := 0; i < 10; i++ {
		jobProcessor.PerformLater(i)
	}

	time.Sleep(1 * time.Second) // Allow some time for jobs to be processed

	if processedCount.Load() > 0 {
		t.Errorf("Expected no jobs to be processed, but got %d", processedCount.Load())
	}
}
