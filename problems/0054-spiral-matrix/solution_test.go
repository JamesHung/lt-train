package spiralmatrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name: "square3",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "rect3x4",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "singleRow",
			matrix: [][]int{
				{3, 4, 5},
			},
			want: []int{3, 4, 5},
		},
		{
			name: "singleCol",
			matrix: [][]int{
				{3},
				{4},
				{5},
			},
			want: []int{3, 4, 5},
		},
		{
			name: "skinny2x3",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: []int{1, 2, 3, 6, 5, 4},
		},
		{
			name: "skinny3x2",
			matrix: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			want: []int{1, 2, 4, 6, 5, 3},
		},
		{
			name: "negativeValues",
			matrix: [][]int{
				{-1, -2},
				{-3, -4},
			},
			want: []int{-1, -2, -4, -3},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := spiralOrder(tc.matrix); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("spiralOrder(%v) = %v, want %v", tc.matrix, got, tc.want)
			}
		})
	}
}
