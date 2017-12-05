package day1_advent2017

import "testing"

func TestDay1(t *testing.T) {
	for _, test := range testCases {
		actual := Day1(test.input)
		if actual != test.expected {
			t.Fatalf("Day1(%d) = %d, want %d", test.input, actual, test.expected)
		}
	}
}

func TestDay1_Part2(t *testing.T) {
	for _, test := range testCasesPart2 {
		actual := Day1_Part2(test.input)
		if actual != test.expected {
			t.Fatalf("Day1(%d) = %d, want %d", test.input, actual, test.expected)
		}
	}
}
