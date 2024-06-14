package dataanalyst

func Variance(a []int) (variance int) {
	mean := Average(a)
	var (
		deviations []float64
		squares    []float64
		squareSum  float64
		i          int
	)
	for _, num := range a {
		deviations = append(deviations, (float64(num) - float64(mean)))
	}
	for _, dev := range deviations {
		squares = append(squares, (dev * dev))
	}
	for _, sqrs := range squares {
		squareSum += sqrs
	}
	for i = range a {
		i++
	}
	variance = int(squareSum / float64(i))
	return
}
