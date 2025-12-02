package day01

import (
	"fmt"
	"strconv"
	"strings"
)

type Day01 struct {}

func mod(a, b int) int {
	return (a % b + b) % b
}

func (d Day01) Part1(input string) (output string) {
	rotations := strings.Split(input, "\n")

	// Initially the dial is pointed to 50
	var dial int = 50
	// (0 - 99) points around the dial 
	var points = 100
	// Password to safe is the number of times the dial is left pointed 
	// at 0 after any rotation in the sequence
	countZero := 0

	for _, r := range rotations {
		rotateDirection := r[0]
		distance, err := strconv.Atoi(r[1:])
		if err != nil {
			fmt.Println(err)
			return ""
		}
		
		if rotateDirection == byte('L') {
			dial = mod(dial- distance, points)
		} else if rotateDirection == byte('R') {
			dial = mod(dial + distance, points)
		}

		if dial == 0 {
			countZero++
		}
	}

	return fmt.Sprint(countZero)
}

func (d Day01) Part2(input string) (output string) {
	// rotations := strings.Split(input, "\n")

	countZero := 0

	fmt.Println("Day02 Part 2 Not Implemented")

	return fmt.Sprint(countZero)
}
