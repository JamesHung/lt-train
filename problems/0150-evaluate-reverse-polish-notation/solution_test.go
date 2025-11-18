package evaluatorpn

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "example1",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			name:   "example2",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			name:   "example3",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
		{
			name:   "negativeDivision",
			tokens: []string{"-7", "3", "/"},
			want:   -2,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := evalRPN(tc.tokens); got != tc.want {
				t.Fatalf("evalRPN(%v) = %d, want %d", tc.tokens, got, tc.want)
			}
		})
	}
}
