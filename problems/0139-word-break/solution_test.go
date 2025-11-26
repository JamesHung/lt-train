package wordbreak

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{
			name:     "example1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			name:     "example2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			name:     "example3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
		{
			name:     "reuseWord",
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			want:     true,
		},
		{
			name:     "singleCharWord",
			s:        "ab",
			wordDict: []string{"a", "b"},
			want:     true,
		},
		{
			name:     "prefixMismatch",
			s:        "abcd",
			wordDict: []string{"ab", "abc", "cd"},
			want:     true,
		},
		{
			name:     "emptyString",
			s:        "",
			wordDict: []string{"a"},
			want:     true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := wordBreak(tc.s, tc.wordDict)
			if got != tc.want {
				t.Fatalf("wordBreak(%q, %v) = %v, want %v", tc.s, tc.wordDict, got, tc.want)
			}
		})
	}
}
