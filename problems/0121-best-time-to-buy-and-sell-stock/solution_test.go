package besttimetobuyandsellstock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "example1", prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{name: "example2", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "singleDay", prices: []int{5}, want: 0},
		{name: "plateauThenRise", prices: []int{3, 3, 3, 6}, want: 3},
		{name: "valleyLate", prices: []int{9, 8, 1, 2, 10}, want: 9},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := maxProfit(tc.prices); got != tc.want {
				t.Fatalf("maxProfit(%v) = %d, want %d", tc.prices, got, tc.want)
			}
		})
	}
}
