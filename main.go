package main

import (
	"flag"
	"fmt"

	"github.com/zeropsykh/advent-of-code-2025/days/day01"
	"github.com/zeropsykh/advent-of-code-2025/days/day02"
	"github.com/zeropsykh/advent-of-code-2025/days/day03"
	"github.com/zeropsykh/advent-of-code-2025/days/day04"
	"github.com/zeropsykh/advent-of-code-2025/days/day05"
	"github.com/zeropsykh/advent-of-code-2025/days/day06"
	helper "github.com/zeropsykh/advent-of-code-2025/internal"
)

type Day interface {
	Part1(string) int
	Part2(string) int 
}

var days = map[int]Day{
	1: day01.Day01{},
	2: day02.Day02{},
	3: day03.Day03{},
	4: day04.Day04{},
	5: day05.Day05{},
	6: day06.Day06{},
}

func main() {
	var day int
	var all bool
	var greetings = "###### Advent of Code '25 ######"
	flag.IntVar(&day, "day", 0, "Day number [1-12]")
	flag.BoolVar(&all, "all", false, "Print output of all days")
	flag.Parse()

	if all {
		fmt.Println(greetings)
		for k, _ := range days {
			DayOutput(k)
		}
		return
	}

	if day < 1 || day > 6 {
		fmt.Println("Usage: AoC25 -d[1-12] -all")
		return
	} 
	fmt.Println(greetings)
	DayOutput(day)
}

func DayOutput(day int) {
	inputFile := fmt.Sprintf("./inputs/day0%d_input.txt", day)
	input := helper.GetInput(inputFile)

	fmt.Printf("############ Day %02d ############\n", day)
	fmt.Println("Output of Part1:", days[day].Part1(input))	
	fmt.Println("Output of Part2:", days[day].Part2(input))	
}
