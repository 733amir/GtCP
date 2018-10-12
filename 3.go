package main

import "log"

var coins = []int{4, 3, 1}

func exchange(money int) []int {
	exchanges := make([]int, money+1)

	for i := range exchanges {
		min := -1
		for j := range coins {
			if diff := i - coins[j]; diff >= 0 {
				if count := exchanges[diff] + 1; count < min || min == -1 {
					min = count
				}
			}
		}

		if min == -1 {
			exchanges[i] = 0
		} else {
			exchanges[i] = min
		}
	}

	return exchanges
}

func main() {
	log.Print(exchange(10))
}
