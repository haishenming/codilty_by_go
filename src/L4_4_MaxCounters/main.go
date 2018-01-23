package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	ret := make([]int, N)

	for _, a := range A {
		if a-1 < N {
			ret[a-1] ++
		} else {
			maxA := make([]int, 2)
			for i, a := range ret {
				if a > maxA[1] {
					maxA[1] = a
					maxA[0] = i
				}
			}
			for i := range ret {
				ret[i] = maxA[1]
			}
		}

	}

	return ret


}

func main() {

	N := 5
	A := []int{3, 4, 4, 6, 1, 4, 4}

	ret := Solution(N, A)

	fmt.Println(ret)
}
