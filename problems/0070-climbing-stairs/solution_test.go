package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 1, want: 1},
		{n: 2, want: 2},
		{n: 3, want: 3},
		{n: 4, want: 5},
		{n: 5, want: 8},
		{n: 10, want: 89},
		{n: 45, want: 1836311903},
	}

	for _, tc := range tests {
		got := climbStairs(tc.n)
		if got != tc.want {
			t.Fatalf("climbStairs(%d) = %d, want %d", tc.n, got, tc.want)
		}
	}
}
