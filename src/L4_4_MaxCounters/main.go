package main

import (
	"fmt"
)


/*
待优化
 */

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	ret := make([]int, N+1)
	maxa := 0
	mina := 0
	for _, a := range A {
		if a < N {
			 if ret[a] < mina {
			 	ret[a] = mina
			 }
			 ret[a] ++
			 if ret[a] > maxa {
			 	maxa = ret[a]
			 }
		} else {
			mina = maxa
		}
	}

	for i, r:= range ret {
		if r < mina {
			ret[i] = mina
		}
	}

	return ret[1:]


}

func main() {

	N := 5
	A := []int{3, 4, 4, 6, 1, 4, 4}

	ret := Solution(N, A)

	fmt.Println(ret)
}
