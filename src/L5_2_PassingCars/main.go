package main

import "fmt"

func Solution(A []int) int {
	// write your code in Go 1.4
	curZeroCnt := 0
	totalPairs := 0
	for _, a := range A {
		if a == 0 {
			curZeroCnt += 1
		} else {
			totalPairs += curZeroCnt

		}
	}


	if totalPairs > 1000000000 {
		return -1
	} else {
		return totalPairs
	}


}


func main() {
	A := []int {0,1,0,1,1}
	ret := Solution(A)
	fmt.Println(ret)
}