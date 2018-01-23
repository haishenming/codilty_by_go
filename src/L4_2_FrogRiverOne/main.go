package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, A []int) int {
	// write your code in Go 1.4
	B := make([]int, X)
	sumB := 0                     // 记录如果满员，总数是多少
	for i:=1;i<=X;i++{
		sumB += i
	}

	realSumB := 0
	for i, a := range A {
		if B[a-1] == 0{
			B[a-1] = a
			realSumB += a
		}
		if realSumB == sumB {
			return i
		}
	}

	return -1
}

func main() {
	var X int

	X = 5
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}

	ret := Solution(X, A)

	fmt.Printf("%d", ret)
}
