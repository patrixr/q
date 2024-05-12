package q

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert.Equal(t,
		[]int{2, 4, 6},
		Map([]int{1, 2, 3}, func(x int) int { return x * 2 }),
	)

	assert.Equal(t,
		[]string{"1", "2", "3"},
		Map([]int{1, 2, 3}, func(x int) string { return strconv.Itoa(x) }),
	)
}
