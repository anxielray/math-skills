package dataanalyst

func Average(a []int) (av int) {
	var num, sum, i int
	for i, num = range a {
		i++
		sum += num
	}
	av = int(sum / i)
	return
}
