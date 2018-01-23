package main

import (
	"fmt"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	X := len(A)
	seens := make([]bool, X)
	fmt.Println(seens)

	for _, a := range A {
		if 0 < a && a <= X {
			seens[a-1] = true
		}
	}

	for i := range seens {
		if seens[i] == false {
			return i + 1
		}
	}

	return X + 1

}

func main() {
	A := []int{1,2,3,4,5,8}

	ret := Solution(A)

	fmt.Printf("%d", ret)
}
