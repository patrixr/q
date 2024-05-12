package q

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	assert.Equal(t,
		[]int{2, 4},
		Filter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 }),
	)

	assert.Equal(t,
		[]string{"1", "3"},
		Filter([]string{"1", "2", "3"}, func(x string) bool { return x != "2" }),
	)
}
