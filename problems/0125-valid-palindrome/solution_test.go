package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "example1", s: "A man, a plan, a canal: Panama", want: true},
		{name: "example2", s: "race a car", want: false},
		{name: "example3", s: " ", want: true},
		{name: "mixedCase", s: "No 'x' in Nixon", want: true},
		{name: "singleChar", s: "Z", want: true},
		{name: "onlyPunctuation", s: "!!!", want: true},
		{name: "edgeNonPalindrome", s: "ab", want: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := isPalindrome(tc.s); got != tc.want {
				t.Fatalf("isPalindrome(%q) = %t, want %t", tc.s, got, tc.want)
			}
		})
	}
}
