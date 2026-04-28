package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, 1, 0, 1, 1, 1}
	count := 0
	maxCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++

			if maxCount < count {
				maxCount = count
			}
		} else {

			count = 0
		}
	}
	fmt.Println("The max count of one is:", maxCount)
}
