package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type Day02 struct {}

func (d Day02) Part1(input string) string {
	idRanges := strings.Split(input, ",")

	// Sum of all invalid ids
	invalidIDSum := 0

	for _, idRange := range idRanges {
		r := strings.Split(idRange, "-")
		firstID, err := strconv.Atoi(r[0]); 
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
		lastID, err := strconv.Atoi(r[1]);
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}

		for id := firstID; id <= lastID; id++ {
			idStr := strconv.Itoa(id)
			l := len(idStr)
			if l % 2 == 0 {
				if idStr[:l/2] == idStr[l/2:] {
					invalidIDSum += id
				}
			}
		}
	}

	return fmt.Sprint(invalidIDSum)
}

func (d Day02) Part2(input string) string {
	idRanges := strings.Split(input, ",")

	// Sum of all invalid ids
	invalidIDSum := 0

	for _, idRange := range idRanges {
		r := strings.Split(idRange, "-")
		firstID, err := strconv.Atoi(r[0]); 
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
		lastID, err := strconv.Atoi(r[1]);
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}

		for id := firstID; id <= lastID; id++ {
			idStr := strconv.Itoa(id)
			maxChunkSize := len(idStr) / 2
			for chunkSize := maxChunkSize; chunkSize >= 1; chunkSize-- {
				// Ignore IDs which cannot be split into chunks of equal size 
				if len(idStr) % chunkSize != 0 {
					continue
				}

				invalid := true
				chunks := splitIntoChunk(idStr, chunkSize)
				firstChunk := chunks[0]

				for _, chunk := range chunks[1:] {
					if chunk != firstChunk {
						invalid = false
						break
					}
				}

				if invalid {
					invalidIDSum += id
					break
				}
			}
		}
	}

	return fmt.Sprint(invalidIDSum)
}

func splitIntoChunk(s string, chunkSize int) []string {
	out := []string{}
	for i := 0; i < len(s); i+= chunkSize {
		out = append(out, s[i:i+chunkSize])
	}
	return out
}
