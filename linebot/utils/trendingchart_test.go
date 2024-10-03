package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTrendingChart(t *testing.T) {
	values := []map[string]float64{
		{"中文": 1, "b": 2, "c": 3},
		{"中文": 2, "c": 4},
		{"中文": 3, "b": 4, "c": 5},
	}
	legends := []string{"中文", "b", "c"}
	title := "test"
	path := "./test.html"

	err := CreateTrendingChart(values, legends, title, path)
	assert.NoError(t, err)
}
