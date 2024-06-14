package dataanalyst

import "math"

func StandardDeviation(a []int) (std int) {
	std = int(math.Sqrt(float64(Variance(a))))
	return
}
