package longestpalindromicsubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "example1",
			s:    "babad",
			want: "bab", // "aba" is also valid
		},
		{
			name: "example2",
			s:    "cbbd",
			want: "bb",
		},
		{
			name: "single char",
			s:    "a",
			want: "a",
		},
		{
			name: "all same",
			s:    "aaaa",
			want: "aaaa",
		},
		{
			name: "two palindromes",
			s:    "abacdfgdcaba",
			want: "aba",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := findLongestPalindromicString(tc.s)
			if len(got) != len(tc.want) {
				t.Fatalf("longestPalindrome(%q) = %q (len %d), want length %d", tc.s, got, len(got), len(tc.want))
			}
		})
	}
}
