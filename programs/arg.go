package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	for k := 1; k < 11; k++ {
		if k%2 == 0 {
			fmt.Printf("\n%v Even number", k)
		}
	}
}
