package q

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	os.Setenv("I_EXIST", "value")
	os.Setenv("I_EXIST_INT", "42")
	os.Setenv("I_EXIST_BOOL", "true")
}

func TestReadEnvDefault(t *testing.T) {
	assert.Equal(t, "default", ReadEnv("I_DONT_EXIST", "default"))
}

func TestReadEnv(t *testing.T) {
	assert.Equal(t, "value", ReadEnv("I_EXIST", "default"))
}

func TestReadEnvIntDefault(t *testing.T) {
	assert.Equal(t, 1, ReadEnvInt("I_DONT_EXIST", 1))
}

func TestReadEnvInt(t *testing.T) {
	assert.Equal(t, 42, ReadEnvInt("I_EXIST_INT", 42))
}

func TestReadEnvBoolDefault(t *testing.T) {
	assert.Equal(t, true, ReadEnvBool("I_DONT_EXIST", true))
}

func TestReadEnvBool(t *testing.T) {
	assert.Equal(t, true, ReadEnvBool("I_EXIST_BOOL", false))
}
