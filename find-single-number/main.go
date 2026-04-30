package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 3, 1}
	hash := map[int]int{}
	hashMethod(nums, hash)
	xorMethod(nums)

}

func hashMethod(nums []int, hash map[int]int) {

	for _, num := range nums {

		hash[num]++
	}

	for _, num := range nums {
		if hash[num] == 1 {
			fmt.Println("The single number is:", num)
			break
		}
	}

}

func xorMethod(nums []int) {
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor = xor ^ nums[i]
	}
	fmt.Println("The single number is:", xor)
}
