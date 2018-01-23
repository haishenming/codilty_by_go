package main

import (
	"fmt"
)

/*

有问题，待修改

 */

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	// 数组最大值和最小值的差
	maxA := 0
	minA := 0
	for _, a := range A {
		if a > maxA {
			maxA = a
		} else if a < minA {
			minA = a
		}
	}

	X := maxA - minA
	B := make([]int, X)
	sumB := 0
	for i:=1;i<=X;i++{
		sumB += i
	}

	realSumB := 0
	for _, a := range A {
		if a <= 0 {
			continue
		}
		if B[a-1] == 0{
			B[a-1] = a
			realSumB += a
		}

	}
	ret := sumB - realSumB
	if realSumB > 0 {
		if ret == 0 {
			return X + 1
		} else {
			return ret
		}
	}


	return 1

}

func main() {
	A := []int{1,2,3,4,5,8}

	ret := Solution(A)

	fmt.Printf("%d", ret)
}
