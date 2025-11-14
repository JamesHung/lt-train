package reverselinkedlist

import "testing"

func TestReverseIterative(t *testing.T) {
	testCases := []struct {
		name string
		vals []int
		want []int
	}{
		{name: "empty", vals: nil, want: nil},
		{name: "single", vals: []int{1}, want: []int{1}},
		{name: "example", vals: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{name: "two elements", vals: []int{1, 2}, want: []int{2, 1}},
		{name: "with negatives", vals: []int{-1, 0, 1}, want: []int{1, 0, -1}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			head := buildList(tc.vals)
			got := listToSlice(reverse(head))
			if !slicesEqual(got, tc.want) {
				t.Fatalf("reverse(%v) = %v, want %v", tc.vals, got, tc.want)
			}
		})
	}
}

func TestReverseRecursive(t *testing.T) {
	testCases := []struct {
		name string
		vals []int
		want []int
	}{
		{name: "empty", vals: nil, want: nil},
		{name: "single", vals: []int{7}, want: []int{7}},
		{name: "long list", vals: []int{1, 2, 3, 4}, want: []int{4, 3, 2, 1}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			head := buildList(tc.vals)
			got := listToSlice(reverseListRecursive(head))
			if !slicesEqual(got, tc.want) {
				t.Fatalf("reverseRecursive(%v) = %v, want %v", tc.vals, got, tc.want)
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
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	if len(result) == 0 {
		return nil
	}
	return result
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
