package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int
	arr = append(arr[:0], arr[1:]...)
	arr = append(arr[:6], arr[7:]...)
	result = arr

	fmt.Println(result)

}
