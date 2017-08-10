package zscore

import (
	"math"
	"sort"
)

var zScorePercentile = []float64{
	-2.326, -2.054, -1.881, -1.751, -1.645, -1.555, -1.476, -1.405, -1.341, -1.282,
	-1.227, -1.175, -1.126, -1.08, -1.036, -0.994, -0.954, -0.915, -0.878, -0.842,
	-0.806, -0.772, -0.739, -0.706, -0.674, -0.643, -0.613, -0.583, -0.553, -0.524,
	-0.496, -0.468, -0.44, -0.412, -0.385, -0.358, -0.332, -0.305, -0.279, -0.253,
	-0.228, -0.202, -0.176, -0.151, -0.126, -0.1, -0.075, -0.05, -0.025, 0.0,
	0.025, 0.05, 0.075, 0.1, 0.126, 0.151, 0.176, 0.202, 0.228, 0.253,
	0.279, 0.305, 0.332, 0.358, 0.385, 0.412, 0.44, 0.468, 0.496, 0.524,
	0.553, 0.583, 0.613, 0.643, 0.674, 0.706, 0.739, 0.772, 0.806, 0.842,
	0.878, 0.915, 0.954, 0.994, 1.036, 1.08, 1.126, 1.175, 1.227, 1.282,
	1.341, 1.405, 1.476, 1.555, 1.645, 1.751, 1.881, 2.054, 2.326,
}

// GetPercentile returns the percentile of the value based on the given l, mean, and sd
func GetPercentile(l float64, mean float64, sd float64, value float64) float64 {
	return float64(searchZScore(zScore(l, mean, sd, value))) / 100.0
}

func searchZScore(zScore float64) int {
	index := sort.SearchFloat64s(zScorePercentile, zScore)
	if index == 0 {
		if zScore < zScorePercentile[0] {
			return 0
		}
		return 1
	} else if index == len(zScorePercentile) {
		if zScore > zScorePercentile[len(zScorePercentile)-1] {
			return 100
		}
		return 99
	} else if math.Abs(zScore-zScorePercentile[index-1]) < math.Abs(zScore-zScorePercentile[index]) {
		return index
	}
	return index + 1
}

func zScore(l float64, mean float64, sd float64, value float64) float64 {
	zIndex := (math.Pow(value/mean, l) - 1) / (sd * l)
	if zIndex > 3 {
		return (3 + ((value - standardDeviation(l, mean, sd, 3)) / (standardDeviation(l, mean, sd, 3) - standardDeviation(l, mean, sd, 2))))
	} else if zIndex < -3 {
		return (-3 + ((value - standardDeviation(l, mean, sd, -3)) / (standardDeviation(l, mean, sd, -2) - standardDeviation(l, mean, sd, -3))))
	} else {
		return zIndex
	}
}

func standardDeviation(l float64, mean float64, sd float64, cutOff float64) float64 {
	return mean * math.Pow(1+(l*sd*cutOff), 1/l)
}
