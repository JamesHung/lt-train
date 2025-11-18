package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "example1",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "example2",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "zeroAmount",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "multipleCoins",
			coins:  []int{2, 5, 10, 1},
			amount: 27,
			want:   4,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := coinChange(tc.coins, tc.amount); got != tc.want {
				t.Fatalf("coinChange(%v, %d) = %d, want %d", tc.coins, tc.amount, got, tc.want)
			}
		})
	}
}
