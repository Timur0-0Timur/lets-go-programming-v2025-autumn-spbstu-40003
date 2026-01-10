package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var left int = 15
		var right int = 30
		var K int
		fmt.Scan(&K)
		for i := 0; i < K; i++ {
			var symb string
			var val string
			fmt.Scan(&symb)
			fmt.Scan(&val)
			if symb == ">=" {
				left, _ = strconv.Atoi(val)
			} else {
				right, _ = strconv.Atoi(val)
			}
			if left > right {
				fmt.Println(-1)
			} else {
				fmt.Println(left)
			}
		}
	}
}
