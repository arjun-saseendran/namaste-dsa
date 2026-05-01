package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i value is %v and j\n value is %v", i, j)
		}
	}
}
