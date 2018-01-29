package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	maxNum := N + 1
	base := 0
	currentMax := 0
	result := make([]int, N)

	for _, value := range(A) {
		if value == maxNum {
			base = currentMax
		} else {
			if result[value-1] < base {
				result[value-1] = base
			}

			result[value-1] += 1

			if currentMax < result[value-1] {
				currentMax = result[value-1]
			}
		}
	}

	for idx, value := range (result) {
		if value < base {
			result[idx] = base
		}
	}

	return result
}

func main() {

	N := 5
	A := []int{3, 4, 4, 6, 1, 4, 4}

	ret := Solution(N, A)

	fmt.Println(ret)
}
