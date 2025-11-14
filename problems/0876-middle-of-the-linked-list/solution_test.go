package middleofthelinkedlist

import "testing"

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{name: "odd length", vals: []int{1, 2, 3, 4, 5}, want: []int{3, 4, 5}},
		{name: "even length", vals: []int{1, 2, 3, 4, 5, 6}, want: []int{4, 5, 6}},
		{name: "single node", vals: []int{7}, want: []int{7}},
		{name: "two nodes", vals: []int{1, 2}, want: []int{2}},
		{name: "mixed values", vals: []int{9, 8, 7, 6, 5, 4, 3}, want: []int{6, 5, 4, 3}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			head := buildList(tc.vals)
			got := listToSlice(middleNode(head))
			if !slicesEqual(got, tc.want) {
				t.Fatalf("middleNode(%v) = %v, want %v", tc.vals, got, tc.want)
			}
		})
	}
}

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var out []int
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
