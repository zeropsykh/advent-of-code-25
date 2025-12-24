package day04

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct{
		desc	string
		input	string
		output 	int	
	}{
		{
			desc: "day04 part1 given test",
			input: "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@.",
			output: 13,	
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := Day04{}
			got := d.Part1(test.input)

			if got != test.output {
				t.Errorf("got %v, want %v", got, test.output)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct{
		desc	string
		input	string
		output	int	
	}{
		{
			desc: "day04 part2 given test",
			input: "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@.",
			output: 43,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := Day04{}
			got := d.Part2(test.input)

			if got != test.output {
				t.Errorf("got %v, want %v", got, test.output)
			}
		})
	}
}
