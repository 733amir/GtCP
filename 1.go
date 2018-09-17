package main

import (
	"fmt"
)

func main() {
	var num int64
	fmt.Scan(&num)
	fmt.Print(num)

	for {
		if num == 1 {
			break
		} else if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
		fmt.Print(" ", num)
	}

	fmt.Println()
}
