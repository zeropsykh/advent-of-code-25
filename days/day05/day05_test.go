package day05

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct{
		desc	string
		input	string
		output	int	
	}{
		{
			desc: "given test",
			input: "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32",
			output: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := Day05{}
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
		{desc: "given test", input: "3-5\n10-14\n16-20\n12-18", output: 14},
		{desc: "test 1",  input: "3-5\n7-10\n15-17\n19-19\n1-2",  output: 13},
		{desc: "test 2",  input: "3-5\n7-10\n15-17\n19-19\n1-4",  output: 13},
		{desc: "test 3",  input: "3-5\n7-10\n15-17\n19-19\n1-6",  output: 14},
		{desc: "test 4",  input: "3-5\n7-10\n15-17\n19-19\n12-14", output: 14},
		{desc: "test 5",  input: "3-5\n7-10\n15-17\n19-19\n9-12",  output: 13},
		{desc: "test 6",  input: "3-5\n7-10\n15-17\n19-19\n11-16", output: 15},
		{desc: "test 7",  input: "3-5\n7-10\n15-17\n19-19\n8-9",   output: 11},
		{desc: "test 8",  input: "3-5\n7-10\n15-17\n19-19\n6-18",  output: 17},
		{desc: "test 9",  input: "3-5\n7-10\n15-17\n19-19\n20-22", output: 14},
		{desc: "test 10", input: "3-5\n7-10\n15-17\n19-24\n20-22", output: 16},
		{desc: "test 11", input: "3-5\n7-10\n15-17\n19-19\n7-10",  output: 11},
		{desc: "test 12", input: "3-5\n7-10\n15-17\n19-19\n7-11",  output: 12},
		{desc: "test 13", input: "3-5\n7-10\n15-17\n19-19\n8-10",  output: 11},
		{desc: "test 14", input: "3-5\n7-10\n15-17\n19-19\n8-16",  output: 15},
		{desc: "test 15", input: "3-5\n7-10\n15-17\n19-19\n6-10",  output: 12},
		{desc: "test 16", input: "3-5\n7-10\n15-17\n19-19\n19-19", output: 11},
		{desc: "test 17", input: "3-5\n7-10\n15-17\n19-19\n1-7",  output: 14},
		{desc: "test 18", input: "3-5\n7-10\n15-17\n19-19\n1-15", output: 18},
		{desc: "test 19", input: "3-5\n7-10\n15-17\n19-19\n10-16",output: 15},
		{desc: "test 20", input: "3-5\n7-10\n15-17\n19-19\n10-22",output: 19},
		{desc: "test 21", input: "3-5\n7-10\n15-17\n19-19\n10-17",output: 15},
		{desc: "test 22", input: "3-5\n7-10\n15-17\n19-19\n9-17", output: 15},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			d := Day05{}
			got := d.Part2(test.input)

			if got != test.output {
				t.Errorf("got %v, want %v", got, test.output)
			}
		})
	}
}
