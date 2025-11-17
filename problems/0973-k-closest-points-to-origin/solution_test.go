package kclosestpointstoorigin

import (
	"reflect"
	"testing"
)

func TestKClosest(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		k      int
		want   [][]int
	}{
		{
			name:   "example1",
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			want:   [][]int{{-2, 2}},
		},
		{
			name:   "example2",
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			want:   [][]int{{3, 3}, {-2, 4}},
		},
		{
			name:   "kEqualsN",
			points: [][]int{{0, 1}, {2, 2}},
			k:      2,
			want:   [][]int{{0, 1}, {2, 2}},
		},
		{
			name:   "tiesHandled",
			points: [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
			k:      3,
			want:   [][]int{{1, 0}, {0, 1}, {-1, 0}},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := kClosest(clonePoints(tc.points), tc.k)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("kClosest(%v, %d) = %v, want %v", tc.points, tc.k, got, tc.want)
			}
		})
	}
}

func clonePoints(pts [][]int) [][]int {
	out := make([][]int, len(pts))
	for i := range pts {
		out[i] = append([]int(nil), pts[i]...)
	}
	return out
}
