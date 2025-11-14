package ransomnote

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{name: "single false", ransomNote: "a", magazine: "b", want: false},
		{name: "double false", ransomNote: "aa", magazine: "ab", want: false},
		{name: "double true", ransomNote: "aa", magazine: "aab", want: true},
		{name: "identical strings", ransomNote: "abc", magazine: "abc", want: true},
		{name: "repeats require counts", ransomNote: "aabbc", magazine: "ababc", want: true},
		{name: "missing letter", ransomNote: "aabbcd", magazine: "ababc", want: false},
		{name: "magazine shorter", ransomNote: "abc", magazine: "ab", want: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := canConstruct(tc.ransomNote, tc.magazine)
			if got != tc.want {
				t.Fatalf("canConstruct(%q, %q) = %v, want %v", tc.ransomNote, tc.magazine, got, tc.want)
			}
		})
	}
}
