package lettercombinationsofaphonenumber

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{name: "example 23", digits: "23", want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{name: "single digit", digits: "2", want: []string{"a", "b", "c"}},
		{name: "length three", digits: "79", want: []string{"pw", "px", "py", "pz", "qw", "qx", "qy", "qz", "rw", "rx", "ry", "rz", "sw", "sx", "sy", "sz"}},
		{name: "empty", digits: "", want: nil},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := letterCombinations(tc.digits)
			sort.Strings(got)
			sort.Strings(tc.want)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("letterCombinations(%q) = %v, want %v", tc.digits, got, tc.want)
			}
		})
	}
}
