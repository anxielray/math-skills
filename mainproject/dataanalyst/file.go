package dataanalyst

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func File(infile string) (data []int) {
	file, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	Error(infile)
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		data = append(data, num)
	}
	return
}
