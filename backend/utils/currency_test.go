package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExchangeRate(t *testing.T) {
	_, err := GetExchangeAmount("TWD", "USD", 1)

	assert.NoError(t, err)
}
