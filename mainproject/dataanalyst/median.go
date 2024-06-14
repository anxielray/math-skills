package dataanalyst

func Median(a []int) (med int) {
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
		med = int(sorted[(i / 2)])
	} else {
		med = int(sorted[i/2]+sorted[(i/2)-1]) / 2
	}
	return
}
