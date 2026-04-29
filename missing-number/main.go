package main

import "fmt"

func main() {
	nums := []int{2, 4, 1, 3, 0}
	n := len(nums)
	totalSum := n * (n + 1) / 2
	sum := 0

	for i := 0; i < n; i++ {
		sum = sum + nums[i]

	}
	fmt.Println(totalSum - sum)
}
