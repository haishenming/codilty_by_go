package main

import "fmt"

func Solution(A []int) int {
	// write your code in Go 1.4
	var current2, current3, min float32
	var minIndex int
	for i := range A {
		if i + 1 == len(A) {
			break
		}
		current2 = (float32(A[i]) + float32(A[i+1])) / float32(2)

		if current2 < min || i == 0 {
			min = current2
			minIndex = i
		}

		if i != len(A) - 2 {
			current3 = (float32(A[i]) + float32(A[i+1]) + float32(A[i+2])) / float32(3)
			if current3 < min {
				min = current3
				minIndex = i
			}
		}
	}

	return minIndex
}


func main() {

	A := make([]int, 7, 7)
	A[0] = 4
	A[1] = 2
	A[2] = 2
	A[3] = 5
	A[4] = 1
	A[5] = 5
	A[6] = 8

	ret := Solution(A)
	fmt.Println(ret)

}
