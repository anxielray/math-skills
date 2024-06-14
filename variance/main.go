package main

import "fmt"

func main() {
	fmt.Println(Variance([]int{1, 2, 3, 4, 5}))
}

func Variance(a []int) (variance float64) {
	mean := Average(a)
	var (
		deviations []float64
		squares    []float64
		squareSum  float64
		i          int
	)
	for _, num := range a {
		deviations = append(deviations, (float64(num) - mean))
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
	variance = (squareSum / float64(i))
	return
}

func Average(a []int) (av float64) {
	var num, sum, i int
	for i, num = range a {
		i++
		sum += num
	}
	av = float64(sum / i)
	return
}
