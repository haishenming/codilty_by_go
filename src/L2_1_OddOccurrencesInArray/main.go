package main

import "fmt"

func Solution(A []int) int {
	// write your code in Go 1.4
	var ACounter map[int]int
	ACounter = make(map[int]int)

	for _, i := range A {
		ACounter[i] = 0
	}

	for k, _ := range ACounter {
		for _, i := range A {
			if i == k {
				ACounter[k] ++
			}
		}
	}

	fmt.Println(ACounter)
	for k, v := range ACounter {
		if v % 2 != 0 {
			return k
		}
	}

	return 0
}

func main() {

	var A []int
	A = []int{9, 3, 9, 3, 9, 7, 9}

	ret := Solution(A)

	fmt.Println("结果", ret)
}