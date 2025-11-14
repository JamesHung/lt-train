package firstbadversion

import "testing"

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name string
		n    int
		bad  int
	}{
		{name: "example case", n: 5, bad: 4},
		{name: "single version", n: 1, bad: 1},
		{name: "bad at start", n: 7, bad: 1},
		{name: "bad at end", n: 10, bad: 10},
		{name: "large numbers", n: 2126753390, bad: 1702766719},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			prev := isBadVersion
			isBadVersion = func(version int) bool {
				return version >= tc.bad
			}
			defer func() { isBadVersion = prev }()

			got := firstBadVersion(tc.n)
			if got != tc.bad {
				t.Fatalf("firstBadVersion(%d) = %d, want %d", tc.n, got, tc.bad)
			}
		})
	}
}
