package main

import "fmt"

func main() {
	fmt.Println(Median([]int{1, 2, 3, 4, 5}))
}

func Median(a []int) (med float64) {
	var sorted []int
	largest := a[0]
	for _, num := range a {
		if num >= largest {
			largest = num
			sorted = append(sorted, largest)
		}
	}
	var i int
	for i = range sorted {
		i++
	}
	if i%2 != 0 {
		med = float64(sorted[(i / 2)])
	} else {
		med = float64(sorted[i/2]+sorted[(i/2)-1]) / 2
	}
	return
}
