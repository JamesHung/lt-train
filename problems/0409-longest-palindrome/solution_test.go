package longestpalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "example 1", s: "abccccdd", want: 7},
		{name: "example 2", s: "a", want: 1},
		{name: "all pairs", s: "aabbcc", want: 6},
		{name: "mixed odd counts", s: "aaabbbbcc", want: 9},
		{name: "case sensitive", s: "Aa", want: 1},
		{name: "uppercase paired", s: "AaBbCcDdEeFF", want: 3},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := longestPalindrome(tc.s)
			if got != tc.want {
				t.Fatalf("longestPalindrome(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}
