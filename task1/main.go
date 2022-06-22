package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8} // 1, 4, -4, 6, 3, 8,
	var result []int
	result = make([]int, 0, len(arr))
	temp := map[int]struct{}{}
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}

	fmt.Println(result)

}
