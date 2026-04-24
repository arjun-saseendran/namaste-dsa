package main

import "fmt"

func main() {
	prices := []int{5, 5, 8, 3, 7}
	min := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i] - min > maxProfit {
			maxProfit = prices[i] - min
		}
	}

	fmt.Println(maxProfit)
}
