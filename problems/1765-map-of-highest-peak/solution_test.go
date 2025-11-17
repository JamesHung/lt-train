package mapofhighestpeak

import (
	"reflect"
	"testing"
)

func TestHighestPeak(t *testing.T) {
	tests := []struct {
		name    string
		isWater [][]int
		want    [][]int
	}{
		{
			name:    "example1",
			isWater: [][]int{{0, 1}, {0, 0}},
			want:    [][]int{{1, 0}, {2, 1}},
		},
		{
			name:    "example2",
			isWater: [][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}},
			want:    [][]int{{1, 1, 0}, {0, 1, 1}, {1, 2, 2}},
		},
		{
			name:    "singleWater",
			isWater: [][]int{{1}},
			want:    [][]int{{0}},
		},
		{
			name:    "checkerboard",
			isWater: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
			want:    [][]int{{0, 1, 0}, {1, 2, 1}, {0, 1, 0}},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := highestPeak(tc.isWater)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("highestPeak(%v) = %v, want %v", tc.isWater, got, tc.want)
			}
		})
	}
}

