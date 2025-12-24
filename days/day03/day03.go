package day03

import (
	"fmt"
	"strconv"
	"strings"
)

type Day03 struct{}

func (d Day03) Part1(input string) int {
	banks := strings.Split(input, "\n")

	totalOutputJoltage := 0

	for _, bank := range banks {
		// Find first battery with largest joltage before last battery in the bank
		var bat1 rune
		var bat1Pos int
		for i, bat := range bank[:len(bank) - 1] {
			if bat > bat1 {
				bat1 = bat
				bat1Pos = i
			} 
		}

		// Find next battery with largest joltage appearing after bat1
		var bat2 rune
		for _, bat := range bank[bat1Pos+1:] {
			if bat > bat2 {
				bat2 = bat
			}
		}

		largestJoltage, err := strconv.Atoi(fmt.Sprintf("%c%c", bat1, bat2))
		if err != nil {
			fmt.Println(err)
		}
		totalOutputJoltage += largestJoltage 	
	}

	return totalOutputJoltage
}


func (d Day03) Part2(input string) int {
	banks := strings.Split(input, "\n")

	// Solved using same approach as Part1

	totalOutputJoltage := 0
	// Max number of digits in outputJoltage of a bank
	maxDigit := 12

	for _, bank := range banks {
		bats := []rune{}
		sStart := 0
		for k := range maxDigit {
			var bat rune
			var batPos int 

			for i := sStart; i < len(bank); i++ {
				if i == len(bank) - maxDigit + k + 1 {
					break
				}
				if rune(bank[i]) > bat {
					bat = rune(bank[i])
					batPos = i
				} 

			}

			bats = append(bats, bat)
			sStart = batPos + 1
			// TODO: optimization
			// when remaining bats = maxDigit - sStart
		}

		largestJoltage, err := strconv.Atoi(string(bats))
		if err != nil {
			fmt.Println(err)
		}

		totalOutputJoltage += largestJoltage
	}

	return totalOutputJoltage
}
