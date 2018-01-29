package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
	// write your code in Go 1.4
	if A > B {
		return 0
	}
	n1 := A / K
	if A % K != 0 {
		n1 ++
	}

	n2 := B / K

	return n2 - n1 + 1

}

func main() {

	A := 6
	B := 11
	K := 2

	ret := Solution(A, B, K)

	fmt.Println(ret)
}
