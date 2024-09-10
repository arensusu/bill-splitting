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
	dataset := [][4]string{
		{"1", "中文", "234", "123"},
		{"2", "中文", "2345", "123"},
		{"3", "c", "23456", "123"},
		{"1", "中文", "234", "123"},
		{"2", "c", "2345", "123"},
		{"3", "b", "23456", "123"},
		{"1", "中文", "234", "123"},
		{"2", "b", "2345", "123"},
		{"3", "3", "23456", "123"},
		{"1", "c", "234", "123"},
		{"2", "中文", "2345", "123"},
		{"3", "b", "23456", "123"},
		{"1", "123", "234", "123"},
		{"2", "中文", "2345", "123"},
		{"3", "b", "23456", "123"},
		{"1", "123", "234", "123"},
		{"2", "b", "2345", "123"},
		{"3", "中文", "23456", "123"},
	}
	path := "./test.html"

	err := CreatePieChart(values, legends, title, subtitle, dataset, path)
	assert.NoError(t, err)
}
