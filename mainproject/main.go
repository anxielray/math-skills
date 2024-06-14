package main

import (
	"fmt"
	"os"

	ray "raymath/dataanalyst"
)

func main() {
	ray.Errors()
	data := ray.File(os.Args[1])
	fmt.Printf("Average: %d\n", ray.Average(data))
	fmt.Printf("Median: %d\n", ray.Median(data))
	fmt.Printf("Variance: %d\n", ray.Variance(data))
	fmt.Printf("Standard Deviation: %d\n", ray.StandardDeviation(data))
}
