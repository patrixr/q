package q

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextIsDone_NotDone(t *testing.T) {
	ctx := context.Background()

	result := ContextIsDone(ctx)

	assert.False(t, result, "Background context should not be done")
}

func TestContextIsDone_CancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	result := ContextIsDone(ctx)

	assert.True(t, result, "Cancelled context should be done")
}

func TestContextIsDone_WithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), -1) // Already expired
	defer cancel()

	result := ContextIsDone(ctx)

	assert.True(t, result, "Expired timeout context should be done")
}
