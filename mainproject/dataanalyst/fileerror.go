package dataanalyst

import (
	"fmt"
	"log"
	"os"
)

func Error(file string) {
	info, err := os.Stat(file)
	if info.Size() == 0 {
		fmt.Println(file, " is empty")
		os.Exit(0)
	}
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
}
