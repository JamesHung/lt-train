package zigzagconversion

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{
			name:    "example1",
			s:       "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			name:    "example2",
			s:       "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			name:    "singleRow",
			s:       "A",
			numRows: 1,
			want:    "A",
		},
		{
			name:    "rowsExceedLength",
			s:       "AB",
			numRows: 5,
			want:    "AB",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := convert(tc.s, tc.numRows); got != tc.want {
				t.Fatalf("convert(%q, %d) = %q, want %q", tc.s, tc.numRows, got, tc.want)
			}
		})
	}
}
