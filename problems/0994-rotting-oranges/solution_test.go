package rottingoranges

import "testing"

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example1",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			want: 4,
		},
		{
			name: "unreachable",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			want: -1,
		},
		{
			name: "no fresh",
			grid: [][]int{
				{0, 2},
			},
			want: 0,
		},
		{
			name: "single fresh",
			grid: [][]int{
				{1},
			},
			want: -1,
		},
		{
			name: "chain",
			grid: [][]int{
				{2, 1, 1, 1},
			},
			want: 3,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gridCopy := copyGrid(tc.grid)
			if got := orangesRotting(gridCopy); got != tc.want {
				t.Fatalf("orangesRotting() = %d, want %d", got, tc.want)
			}
		})
	}
}

func copyGrid(src [][]int) [][]int {
	dup := make([][]int, len(src))
	for i := range src {
		dup[i] = append([]int(nil), src[i]...)
	}
	return dup
}
