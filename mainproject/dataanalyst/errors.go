package dataanalyst

import (
	"fmt"
	"os"
)

func Errors() {
	if len(os.Args) != 2 {
		fmt.Println("wrong passing of arguments...")
		os.Exit(0)
	}
	
}
