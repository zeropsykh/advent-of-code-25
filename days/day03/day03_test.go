package day03

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct{
		desc 	string
		input 	string
		output	string
	}{
		{
			desc: "day03 part1 given test",
			input: "987654321111111\n811111111111119\n234234234234278\n818181911112111",
			output: "357",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func (t * testing.T)  {
			d := Day03{}
			got := d.Part1(test.input)

			if got != test.output {
				t.Errorf("got %s, want %s", got, test.output)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct{
		desc 	string
		input 	string
		output	string
	}{
		{
			desc: "day03 part2 test 1",
			input: "987654321111111",
			output: "987654321111",
		},
		{
			desc: "day03 part2 test 2",
			input: "811111111111119",
			output: "811111111119",
		},
		{
			desc: "day03 part2 test 3",
			input: "234234234234278",
			output: "434234234278",
		},
		{
			desc: "day03 part2 test 4",
			input: "818181911112111",
			output: "888911112111",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func (t * testing.T)  {
			d := Day03{}
			got := d.Part2(test.input)

			if got != test.output {
				t.Errorf("got %s, want %s", got, test.output)
			}
		})
	}
}
