package day02

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct{
		desc 	string
		input 	string
		output	int	
	}{
		{
			desc: "day02 part1 given test",
			input: "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
			output: 1227775554,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := Day02{}
			got := d.Part1(test.input)

			if got != test.output {
				t.Errorf("got %v, want %v", got, test.output)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct{
		desc 	string
		input 	string
		output 	int	
	}{
		{
			desc: "day02 part2 given test",
			input: "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
			output: 4174379265,
		},
		{
			desc: "day02 part2 test2",
			input: "824824821-824824827",
			output: 824824824,
		},
		{
			desc: "day02 part2 test3",
			input: "1188511880-1188511890",
			output: 1188511885,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := Day02{}
			got := d.Part2(test.input)

			if got != test.output {
				t.Errorf("got %v, want %v", got, test.output)
			}
		})
	}
}
