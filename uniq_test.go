package q

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	assert.Equal(t,
		[]int{1, 2, 3},
		Uniq([]int{1, 2, 3}),
	)

	assert.Equal(t,
		[]int{1, 2, 3},
		Uniq([]int{1, 2, 3, 1, 2, 3}),
	)

	assert.Equal(t,
		[]int{1, 2, 3},
		Uniq([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}),
	)

	assert.Equal(t,
		[]string{"a", "b", "c"},
		Uniq([]string{"a", "b", "c"}),
	)

	assert.Equal(t,
		[]string{"a", "b", "c"},
		Uniq([]string{"a", "b", "c", "a", "b", "c"}),
	)

	assert.Equal(t,
		[]string{"a", "b", "c"},
		Uniq([]string{"a", "a", "a", "b", "b", "b", "c", "c", "c"}),
	)
}
