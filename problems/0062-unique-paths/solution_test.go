package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m, n int
		want int
	}{
		{name: "example wide", m: 3, n: 7, want: 28},
		{name: "example tall", m: 3, n: 2, want: 3},
		{name: "single cell", m: 1, n: 1, want: 1},
		{name: "single row", m: 1, n: 5, want: 1},
		{name: "single column", m: 5, n: 1, want: 1},
		{name: "symmetry", m: 7, n: 3, want: 28},
		{name: "square mid", m: 10, n: 10, want: 48620},
		{name: "larger square", m: 15, n: 15, want: 40116600},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := uniquePaths(tc.m, tc.n); got != tc.want {
				t.Fatalf("uniquePaths(%d, %d) = %d, want %d", tc.m, tc.n, got, tc.want)
			}
		})
	}
}
