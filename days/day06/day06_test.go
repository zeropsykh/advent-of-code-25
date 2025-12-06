package day06

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct{
		desc	string
		input	string
		output	string
	}{
		{
			desc: "given test",
			input: "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  ",
			output: "4277556",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := Day06{}
			got := d.Part1(test.input)

			if got != test.output {
				t.Errorf("got %s, want %s", got, test.output)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct{
		desc	string
		input	string
		output	string
	}{
		{
			desc: "given test",
			input: "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  ",
			output: "3263827",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := Day06{}
			got := d.Part2(test.input)

			if got != test.output {
				t.Errorf("got %s, want %s", got, test.output)
			}
		})
	}
}
