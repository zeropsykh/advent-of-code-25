package day01

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct{
		desc string
		input string
		output string
	}{
		{ 
			desc: "day01 part1 given test",
			input: "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82",
			output: "3",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := Day01{}
			got := d.Part1(test.input)

			if got != test.output {
				t.Errorf("got %s, want %s", got, test.output)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct{
		desc string
		input string
		output string
	}{
		{ 
			desc: "Day01 Part2 given test",
			input: "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82",
			output: "6",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := Day01{}
			got := d.Part2(test.input)

			if got != test.output {
				t.Errorf("got %s, want %s", got, test.output)
			}
		})
	}
}
