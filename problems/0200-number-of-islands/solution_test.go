package numberofislands

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "example1",
			grid: [][]byte{
				[]byte("11110"),
				[]byte("11010"),
				[]byte("11000"),
				[]byte("00000"),
			},
			want: 1,
		},
		{
			name: "example2",
			grid: [][]byte{
				[]byte("11000"),
				[]byte("11000"),
				[]byte("00100"),
				[]byte("00011"),
			},
			want: 3,
		},
		{
			name: "single island",
			grid: [][]byte{
				[]byte("1"),
			},
			want: 1,
		},
		{
			name: "no land",
			grid: [][]byte{
				[]byte("000"),
				[]byte("000"),
			},
			want: 0,
		},
		{
			name: "diagonal separate",
			grid: [][]byte{
				[]byte("101"),
				[]byte("010"),
				[]byte("101"),
			},
			want: 5,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gridCopy := copyGrid(tc.grid)
			if got := numIslands(gridCopy); got != tc.want {
				t.Fatalf("numIslands() = %d, want %d", got, tc.want)
			}
		})
	}
}

func copyGrid(src [][]byte) [][]byte {
	dup := make([][]byte, len(src))
	for i := range src {
		dup[i] = append([]byte(nil), src[i]...)
	}
	return dup
}
