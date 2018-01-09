package main

import (
	"strconv"
	"fmt"
)

func Solution(N int) int {
	// write your code in Go 1.4
	b_N := strconv.FormatInt(int64(N), 2)


	o_numbers := 0
	l_numbers := 0
	max_o := 0
	for i:=0;i<len(b_N);i++ {
		if b_N[i] == 49 {
			l_numbers ++
		} else if b_N[i] == 48 {
			o_numbers ++
		}
		if l_numbers == 2 {
			if max_o > o_numbers {
				max_o = max_o
			} else {
				max_o = o_numbers
			}
			l_numbers = 1
			o_numbers = 0
		}
	}

	return max_o
}

func main() {
	N := 1041

	ret := Solution(N)

	fmt.Println(ret)
}
