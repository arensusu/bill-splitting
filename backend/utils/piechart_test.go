package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePieChart(t *testing.T) {
	values := []float64{1, 2, 3}
	legends := []string{"中文", "b", "c"}
	title := "test"
	subtitle := fmt.Sprintf("%.0f", 123.123456)
	path := "./test.html"

	err := CreatePieChart(values, legends, title, subtitle, path)
	assert.NoError(t, err)
}
