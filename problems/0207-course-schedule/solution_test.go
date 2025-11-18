package courseschedule

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name        string
		numCourses  int
		prereqs     [][]int
		wantCanPass bool
	}{
		{
			name:        "example1",
			numCourses:  2,
			prereqs:     [][]int{{1, 0}},
			wantCanPass: true,
		},
		{
			name:        "example2",
			numCourses:  2,
			prereqs:     [][]int{{1, 0}, {0, 1}},
			wantCanPass: false,
		},
		{
			name:        "disconnected",
			numCourses:  4,
			prereqs:     [][]int{{3, 2}, {1, 0}},
			wantCanPass: true,
		},
		{
			name:        "selfLoop",
			numCourses:  1,
			prereqs:     [][]int{{0, 0}},
			wantCanPass: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := canFinish(tc.numCourses, tc.prereqs)
			if got != tc.wantCanPass {
				t.Fatalf("canFinish(%d, %v) = %v, want %v", tc.numCourses, tc.prereqs, got, tc.wantCanPass)
			}
		})
	}
}
