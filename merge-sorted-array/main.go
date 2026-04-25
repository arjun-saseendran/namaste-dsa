package main

import "fmt"

func main() {
	nums1 := [8]int{1, 3, 5, 7}
	nums2 := [4]int{3, 6, 7, 8}
	m := 4
	n := 4
	// nums1Copy := make([]int, m)
	// copy(nums1Copy, nums1[:m])

	// p1 := 0
	// p2 := 0
	p1 :=  m - 1
	p2 := n - 1

	// for i := 0; i < len(nums1); i++ {
		// if p2 >= n || p1 < m && nums1Copy[p1] < nums2[p2] {
		// 	nums1[i] = nums1Copy[p1]
		// 	p1++
		// } else {
		// 	nums1[i] = nums2[p2]
		// 	p2++
		// }
		
	// }
	
	for i := m + n - 1; i >= 0; i-- {
		if p2 < 0{
			break
		}
		if p1 >= 0 && nums1[p1] > nums2[p2]{
			nums1[i] = nums2[p1]
			p1--
		}else{
			nums1[i] = nums2[p2]
			p2--
		}
	}

	fmt.Println("The updated sorted array is:", nums1)
}
