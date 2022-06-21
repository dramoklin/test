package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	arrayNumber := strings.Split(input, "")

	numberMax := arrayNumber[0]
	numberMin := arrayNumber[0]

	for i := 0; i < len(arrayNumber); i++ {

		if numberMin < arrayNumber[i] {
			numberMin = arrayNumber[i]
		}
		if numberMax < arrayNumber[i] {
			numberMax = arrayNumber[i]
		}

	}

	result = numberMin + " " + numberMax

	fmt.Println(result)
}
