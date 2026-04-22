package main

import "fmt"

func main() {
	num := 3
	nums := [10]int{3, 4, 3, 6, 2, 3, 5, 1, 7, 3}
	result := remove_num(nums, num)
	fmt.Println("Removed element count is: ", result)
}

func remove_num(arr [10]int, n int) int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != n {
			arr[j] = arr[i]
			j += 1
		}
	}
	fmt.Println("Modified array is: ", arr)
	return len(arr)- j
}
