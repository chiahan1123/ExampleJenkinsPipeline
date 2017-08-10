package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberToString(t *testing.T) {
	assert.Equal(t, "0", NumberToString(0, 2))
	assert.Equal(t, "2", NumberToString(2, 2))
	assert.Equal(t, "0.5", NumberToString(0.5, 2))
	assert.Equal(t, "1.34", NumberToString(1.34, 2))
	assert.Equal(t, "2.57", NumberToString(2.567, 2))
	assert.Equal(t, "-1", NumberToString(-1, 2))
	assert.Equal(t, "-1.23", NumberToString(-1.234, 2))

	assert.Equal(t, "2.5", NumberToString(2.500, 3))
	assert.Equal(t, "2.56", NumberToString(2.560, 3))

	assert.Equal(t, "3", NumberToString(3, 0))
	assert.Equal(t, "4", NumberToString(3.5, 0))

	assert.Equal(t, "", NumberToString(3, -1))
	assert.Equal(t, "", NumberToString(3.5, -1))
}
