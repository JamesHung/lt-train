package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "example1", s: "()", want: true},
		{name: "example2", s: "()[]{}", want: true},
		{name: "example3", s: "(]", want: false},
		{name: "example4", s: "([])", want: true},
		{name: "example5", s: "([)]", want: false},
		{name: "singleOpening", s: "(", want: false},
		{name: "onlyClosing", s: "}", want: false},
		{name: "nestedSameType", s: "{{{{}}}}", want: true},
		{name: "longFalse", s: "(((())))]", want: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := isValid(tc.s); got != tc.want {
				t.Fatalf("isValid(%q) = %t, want %t", tc.s, got, tc.want)
			}
		})
	}
}
