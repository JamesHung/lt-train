package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{name: "example", height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{name: "two bars", height: []int{1, 1}, want: 1},
		{name: "plateau", height: []int{4, 4, 4, 4}, want: 12},
		{name: "increasing", height: []int{1, 2, 3, 4, 5}, want: 6},
		{name: "decreasing", height: []int{5, 4, 3, 2, 1}, want: 6},
		{name: "zeros", height: []int{0, 0, 0, 0}, want: 0},
		{name: "single tall inside", height: []int{2, 3, 10, 5, 7, 8, 9}, want: 36},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := maxArea(tc.height); got != tc.want {
				t.Fatalf("maxArea(%v) = %d, want %d", tc.height, got, tc.want)
			}
		})
	}
}
