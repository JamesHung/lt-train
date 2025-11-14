package lengthoflongestsubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "example2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "example3",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "empty",
			s:    "",
			want: 0,
		},
		{
			name: "unicode",
			s:    "åßå∂ƒ",
			want: 4,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tc.s); got != tc.want {
				t.Fatalf("lengthOfLongestSubstring(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}
