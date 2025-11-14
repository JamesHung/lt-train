package floodfill

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	tests := []struct {
		name   string
		image  [][]int
		sr, sc int
		color  int
		want   [][]int
	}{
		{
			name:  "example repaint",
			image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:    1,
			sc:    1,
			color: 2,
			want:  [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			name:  "example already colored",
			image: [][]int{{0, 0, 0}, {0, 0, 0}},
			sr:    0,
			sc:    0,
			color: 0,
			want:  [][]int{{0, 0, 0}, {0, 0, 0}},
		},
		{
			name:  "touching border",
			image: [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 0}},
			sr:    0,
			sc:    1,
			color: 3,
			want:  [][]int{{0, 3, 3}, {3, 3, 3}, {3, 3, 0}},
		},
	}

	for _, tc := range tests {
		tc := tc
		// Copy image to avoid sharing underlying slices between table entries
		imgCopy := make([][]int, len(tc.image))
		for i := range tc.image {
			imgCopy[i] = append([]int(nil), tc.image[i]...)
		}

		t.Run(tc.name, func(t *testing.T) {
			got := floodFill(imgCopy, tc.sr, tc.sc, tc.color)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("floodFill() = %v, want %v", got, tc.want)
			}
		})
	}
}
