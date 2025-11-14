package addbinary

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{name: "example1", a: "11", b: "1", want: "100"},
		{name: "example2", a: "1010", b: "1011", want: "10101"},
		{name: "carry entire length", a: "1111", b: "1", want: "10000"},
		{name: "different lengths", a: "1", b: "111", want: "1000"},
		{name: "long same digits", a: "0", b: "0", want: "0"},
		{name: "large inputs", a: "1000000000000001", b: "1111111111111111", want: "11000000000000000"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := addBinary(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("addBinary(%q, %q) = %q, want %q", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
