package main

import "fmt"

func main() {
	x := 0
	
	nums := []int{1, 4, 0, 2, 4, 0, 9, 5, 0}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[x] = nums[i]
			x++
		}
		
	}
	for i := x; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println("Updated array is:", nums)
}
