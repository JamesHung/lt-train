package stringtointegeratoi

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "example1", s: "42", want: 42},
		{name: "example2", s: "   -042", want: -42},
		{name: "example3", s: "1337c0d3", want: 1337},
		{name: "example4", s: "0-1", want: 0},
		{name: "example5", s: "words and 987", want: 0},
		{name: "onlySpaces", s: "     ", want: 0},
		{name: "onlySign", s: "+", want: 0},
		{name: "leadingPlus", s: "+23", want: 23},
		{name: "leadingZeros", s: "   +000120", want: 120},
		{name: "stopAtSpace", s: "  -0012a42", want: -12},
		{name: "positiveOverflow", s: "91283472332", want: 2147483647},
		{name: "negativeOverflow", s: "-91283472332", want: -2147483648},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := myAtoi(tc.s); got != tc.want {
				t.Fatalf("myAtoi(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}
