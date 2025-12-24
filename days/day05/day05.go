package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day05 struct {}

type idRange struct {
	start	int
	end		int
}

func (d Day05) Part1(input string) int {
	database := strings.Split(input, "\n\n")

	idRanges := strings.Split(database[0], "\n")
	availabeIDs := strings.Split(database[1], "\n")

	var freshIDRanges []idRange
	for _, idRng := range idRanges {
		t := strings.Split(string(idRng), "-")
		start, err := strconv.Atoi(t[0])
		if err != nil {
			fmt.Println(err)
		}
		end, err := strconv.Atoi(t[1])
		if err != nil {
			fmt.Println(err)
		}

		freshIDRanges = append(freshIDRanges, idRange{start, end})
	}

	// Count of fresh ingrediant IDs 
	var freshAvailbleIDCount int
	for _, availableID := range availabeIDs {
		id, err := strconv.Atoi(availableID)
		if err != nil {
			fmt.Println(err)
		}

		for _, idRng := range freshIDRanges {
			if id >= idRng.start && id <= idRng.end {
				freshAvailbleIDCount++
				break
			}
		}
	}

	return freshAvailbleIDCount
}

func (d Day05) Part2(input string) int {
	database := strings.Split(input, "\n\n")

	idRanges := strings.Split(database[0], "\n")

	var freshIDRanges []idRange
	for _, r := range idRanges {
		t := strings.Split(string(r), "-")
		start, err := strconv.Atoi(t[0])
		if err != nil {
			fmt.Println(err)
		}
		end, err := strconv.Atoi(t[1])
		if err != nil {
			fmt.Println(err)
		}

		// TODO: Make it more readable and optimized
		var isInserted bool
		for i, v := range freshIDRanges {
			if start < v.start {
				if end < v.start {
					freshIDRanges = slices.Insert(freshIDRanges, i, idRange{start, end})
					isInserted = true
					break
				} else if end > v.end {
					freshIDRanges[i].start = start
					start = v.end + 1
				} else if end >= v.start {
					freshIDRanges[i].start = start
					isInserted = true
					break
				}
			} else if start >= v.start {
				if end <= v.end {
					isInserted = true
					break
				} else if start < v.end {
					start = v.end + 1
				} else if start == v.end {
					start += 1
				}
			}
		}

		if !isInserted {
			freshIDRanges = append(freshIDRanges, idRange{start, end})
		}
	}

	freshAvailableIDCount := 0	
	for _, idRng := range freshIDRanges {
		freshAvailableIDCount += idRng.end - idRng.start + 1	
	}
	
	return freshAvailableIDCount
}

