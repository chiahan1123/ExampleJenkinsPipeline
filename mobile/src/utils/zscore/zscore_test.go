package zscore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZScore(t *testing.T) {
	delta := 0.001
	assert.InDelta(t, 3.408, zScore(-0.1600954, 9.476500305, 0.11218624, 14), delta)
	assert.InDelta(t, -4.07, zScore(-0.1600954, 9.476500305, 0.11218624, 6), delta)
	assert.InDelta(t, 0.207, zScore(-0.1600954, 9.476500305, 0.11218624, 9.7), delta)
}

func TestSearchZScore(t *testing.T) {
	assert.EqualValues(t, 0, searchZScore(-2.327))
	assert.EqualValues(t, 1, searchZScore(-2.326))
	assert.EqualValues(t, 1, searchZScore(-2.325))
	assert.EqualValues(t, 2, searchZScore(-2.055))
	assert.EqualValues(t, 2, searchZScore(-2.054))
	assert.EqualValues(t, 50, searchZScore(0.0))
	assert.EqualValues(t, 50, searchZScore(0.001))
	assert.EqualValues(t, 51, searchZScore(0.024))
	assert.EqualValues(t, 98, searchZScore(2.054))
	assert.EqualValues(t, 98, searchZScore(2.055))
	assert.EqualValues(t, 99, searchZScore(2.325))
	assert.EqualValues(t, 99, searchZScore(2.326))
	assert.EqualValues(t, 100, searchZScore(2.327))
}
