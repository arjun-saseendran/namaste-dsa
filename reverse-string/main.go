package main

import "fmt"

func main() {
	name := [5]string{"m", "a", "r", "i", "a"}

	reverse_string(name)

}

func reverse_string(arr [5]string) {
	temp := ""
	for i := 0; i < len(arr)/2; i++ {
		temp = arr[i]
		arr[i] = arr[len(arr)-1-i]
		arr[len(arr)-1-i] = temp
	}
	fmt.Println(arr)
}
