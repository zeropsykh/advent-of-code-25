package helper

import (
	"fmt"
	"os"
)

func GetInput(input_file string) string {
	content, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	return string(content)
}

func SumArray(arr []int) int {
	var sum int
	for _, num := range arr {
		sum += num
	}
	return sum
}

func ProductArray(arr []int) int {
	var product int = 1
	for _, num := range arr {
		product *= num
	}
	return product
}
