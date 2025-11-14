package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "example1", s: "anagram", t: "nagaram", want: true},
		{name: "example2", s: "rat", t: "car", want: false},
		{name: "singleChar", s: "a", t: "a", want: true},
		{name: "differentLengths", s: "ab", t: "abc", want: false},
		{name: "repeatedLetters", s: "aabbcc", t: "abcabc", want: true},
		{name: "almostAnagram", s: "aabbccd", t: "abcabcc", want: false},
		{name: "leetcode", s: "ggii", t: "eekk", want: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := isAnagram(tc.s, tc.t); got != tc.want {
				t.Fatalf("isAnagram(%q, %q) = %t, want %t", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
