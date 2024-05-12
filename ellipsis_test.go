package q

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEllipsis(t *testing.T) {
	assert.Equal(t, "abc", Ellipsis("abc", 5))
	assert.Equal(t, "abc", Ellipsis("abc", 3))
	assert.Equal(t, "abc...", Ellipsis("abcd", 3))
	assert.Equal(t, "abcd...", Ellipsis("abcde", 4))
	assert.Equal(t, "abcde", Ellipsis("abcde", 5))
	assert.Equal(t, "ab...", Ellipsis("abcde", 2))
}
