package addtwonumbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "example1",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			name: "example2",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			name: "example3",
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := listToSlice(addTwoNumbers(sliceToList(tc.l1), sliceToList(tc.l2)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("addTwoNumbers(%v, %v) = %v, want %v", tc.l1, tc.l2, got, tc.want)
			}
		})
	}
}

func sliceToList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var out []int
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}
