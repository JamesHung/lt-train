package matrix01

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want [][]int
	}{
		{
			name: "example1",
			mat:  [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			name: "example2",
			mat:  [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
		{
			name: "singleRow",
			mat:  [][]int{{1, 1, 0, 1}},
			want: [][]int{{2, 1, 0, 1}},
		},
		{
			name: "allZeros",
			mat:  [][]int{{0, 0}, {0, 0}},
			want: [][]int{{0, 0}, {0, 0}},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := updateMatrix(tc.mat)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("updateMatrix(%v) = %v, want %v", tc.mat, got, tc.want)
			}
		})
	}
}
