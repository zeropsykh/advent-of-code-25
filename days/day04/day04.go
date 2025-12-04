package day04

import (
	"fmt"
	"strings"
)

type Day04 struct {}

func (d Day04) Part1(input string) string {
	grid := strings.Split(input, "\n")

	// Number of paper roll that can be forklifted
	rollCount := 0

	for x, row := range grid {
		for y, roll := range row {
			if roll == '@' {
				adjacentRollCount := 0
				// Checking adjacent position contain fewer than 4 rolls
				for j := -1; j < 2 && adjacentRollCount < 4; j++ {
					for i := -1; i < 2 && adjacentRollCount < 4; i++ {
						if i == 0 && j == 0 {
							continue
						} 

						xPos := x + j
						yPos := y + i
						if xPos < 0 || yPos < 0 || xPos > len(row) - 1 || yPos > len(row) - 1 {
							continue
						}

						if grid[xPos][yPos] == '@' {
							adjacentRollCount++
						}
					}
				}

				if adjacentRollCount < 4 {
					rollCount++
				}
			}
		}
	}

	return fmt.Sprint(rollCount)
}

func (d Day04) Part2(input string) string {
	grid := strings.Split(input, "\n")

	totalRollCount := 0

	for {
		rollCount := 0
		rollIndices := []struct{x, y int}{}
		for x, row := range grid {
			for y, roll := range row {
				if roll == '@' {
					adjacentRollCount := 0
					// Checking adjacent position contain fewer than 4 rolls
					for j := -1; j < 2 && adjacentRollCount < 4; j++ {
						for i := -1; i < 2 && adjacentRollCount < 4; i++ {
							if i == 0 && j == 0 {
								continue
							}

							xPos := x + j
							yPos := y + i
							if xPos < 0 || yPos < 0 || xPos > len(row) - 1 || yPos > len(row) - 1 {
								continue
							}

							if grid[xPos][yPos] == '@' {
								adjacentRollCount++
							}
						}
					}

					if adjacentRollCount < 4 {
						rollIndices = append(rollIndices, struct{x, y int}{x, y})
						rollCount++
					}
				}
			}
		}

		// If no paper roll can be fork lifted - STOP
		if rollCount == 0 {
			break
		}

		for _, rollIndex := range rollIndices {
			runes := []rune(grid[rollIndex.x])
			runes[rollIndex.y] = '.'
			grid[rollIndex.x] = string(runes)
		}
		totalRollCount += rollCount
	}

	return fmt.Sprint(totalRollCount)
}

