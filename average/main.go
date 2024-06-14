package main

import "fmt"

func main() {
	fmt.Println(Average([]int{1, 2, 3, 4, 5}))
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
