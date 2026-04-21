package main

import "fmt"

func main() {
	var j int8 = 0
	nums := [11]int8{
		0, 0, 1, 1, 2, 2, 2, 4, 4, 4, 4,
	}
	result := removeDuplicate(nums, j)
	fmt.Println("The unique numbers count is: ", result)
}

// Remove duplicate from sorted array.
func removeDuplicate(nums [11]int8, j int8) int8 {

	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[j] {
			j += 1
			nums[j] = nums[i]
		}
	}
fmt.Println("updates array is: ",nums)
	return j + 1
}
