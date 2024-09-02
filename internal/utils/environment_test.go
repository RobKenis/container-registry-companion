package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv_exists(t *testing.T) {
	t.Setenv("VALUE_TO_FIND", "hello")
	value := GetEnv("VALUE_TO_FIND", "the-default-value")
	assert.Equal(t, "hello", value)
}

func TestGetEnv_defaultValue(t *testing.T) {
	value := GetEnv("VALUE_TO_FIND", "the-default-value")
	assert.Equal(t, "the-default-value", value)
}
