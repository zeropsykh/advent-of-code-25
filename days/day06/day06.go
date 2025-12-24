package day06

import (
	"fmt"
	"strconv"
	"strings"

	helper "github.com/zeropsykh/advent-of-code-2025/internal"
)

type Day06 struct {}

func (d Day06) Part1(input string) int {
	lines := strings.SplitAfter(input, "\n")

	var problems [][]int

	for _, line := range lines[:len(lines) - 1] {
		var numStart, k int	
		for i, c := range line {
			if (c == ' ' || c == '\n') {
				if numStart == i {
					numStart++
					continue
				}

				num, err := strconv.Atoi(line[numStart:i])
				if err != nil {
					fmt.Println(err)
				}

				if len(problems) == k {
					problems = append(problems, []int{num})
				} else {
					problems[k] = append(problems[k], num)
				}

				k++
				numStart = i + 1
			}
		}
	}

	grandTotal := calculateGrandTotal(problems, lines[len(lines) - 1])

	return grandTotal
}

func (d Day06) Part2(input string) int {
	lines := strings.SplitAfter(input, "\n")
	
	var (
		problems [][]int
		nums []int
	)

	for i := range len(lines[0]) {
		var numStr []rune
		for _, line := range lines[:len(lines) - 1] {
			c := line[i]
			if c == ' ' || c == '\n' {
				continue
			}
			numStr = append(numStr, rune(c))
		}
		
		if len(numStr) == 0 {
			problems = append(problems, nums)
			nums = []int{}
			continue
		}

		num, err := strconv.Atoi(string(numStr))
		if err != nil {
			fmt.Println(err)
		}
		nums = append(nums, num)
	}

	grandTotal := calculateGrandTotal(problems, lines[len(lines) - 1])

	return grandTotal
}

func calculateGrandTotal(numbers [][]int, operations string) int {
	var answers []int
	var k int

	// Solves the problems
	for _, c := range operations {
		switch c {
		case '+':
			answers = append(answers, helper.SumArray(numbers[k]))
			k++
		case '*':
			answers = append(answers, helper.ProductArray(numbers[k]))
			k++
		default:
			continue
		}
	}

	return helper.SumArray(answers)
}
