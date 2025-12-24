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
			desc: "given test",
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
			desc: "given test",
			input: "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82",
			output: "6",
		},
		{ 
			desc: "test 1",
			input: "R49",
			output: "0",
		},
		{ 
			desc: "test 2",
			input: "R50",
			output: "1",
		},
		{ 
			desc: "test 3",
			input: "R100",
			output: "1",
		},
		{ 
			desc: "test 4",
			input: "R50\nR50",
			output: "1",
		},
		{ 
			desc: "test 5",
			input: "R50\nR100",
			output: "2",
		},
		{ 
			desc: "test 6",
			input: "R50\nR240",
			output: "3",
		},
		{ 
			desc: "test 7",
			input: "L49",
			output: "0",
		},
		{ 
			desc: "test 7",
			input: "L50",
			output: "1",
		},
		{ 
			desc: "test 8",
			input: "L55",
			output: "1",
		},
		{ 
			desc: "test 9",
			input: "L100",
			output: "1",
		},
		{ 
			desc: "test 10",
			input: "L149",
			output: "1",
		},
		{ 
			desc: "test 11",
			input: "L150",
			output: "2",
		},
		{ 
			desc: "test 12",
			input: "L50\nL140",
			output: "2",
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
