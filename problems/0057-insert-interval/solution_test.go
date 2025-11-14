package insertinterval

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{name: "example1", intervals: [][]int{{1, 3}, {6, 9}}, newInterval: []int{2, 5}, want: [][]int{{1, 5}, {6, 9}}},
		{name: "example2", intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, newInterval: []int{4, 8}, want: [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{name: "emptyIntervals", intervals: nil, newInterval: []int{4, 8}, want: [][]int{{4, 8}}},
		{name: "insertAtEnd", intervals: [][]int{{1, 2}}, newInterval: []int{3, 4}, want: [][]int{{1, 2}, {3, 4}}},
		{name: "insertAtBeginning", intervals: [][]int{{5, 7}}, newInterval: []int{1, 3}, want: [][]int{{1, 3}, {5, 7}}},
		{name: "containedInside", intervals: [][]int{{1, 5}}, newInterval: []int{2, 3}, want: [][]int{{1, 5}}},
		{name: "touchingRight", intervals: [][]int{{1, 2}, {5, 7}}, newInterval: []int{2, 5}, want: [][]int{{1, 7}}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := insert(tc.intervals, tc.newInterval)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("insert(%v, %v) = %v, want %v", tc.intervals, tc.newInterval, got, tc.want)
			}
		})
	}
}
