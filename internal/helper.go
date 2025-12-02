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
