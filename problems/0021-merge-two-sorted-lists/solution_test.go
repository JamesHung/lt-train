package mergetwosortedlists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name   string
		list1  []int
		list2  []int
		expect []int
	}{
		{name: "example1", list1: []int{1, 2, 4}, list2: []int{1, 3, 4}, expect: []int{1, 1, 2, 3, 4, 4}},
		{name: "example2", list1: nil, list2: nil, expect: nil},
		{name: "example3", list1: nil, list2: []int{0}, expect: []int{0}},
		{name: "sameLength", list1: []int{1, 5, 7}, list2: []int{2, 4, 6}, expect: []int{1, 2, 4, 5, 6, 7}},
		{name: "oneEmpty", list1: []int{1, 2, 3}, list2: nil, expect: []int{1, 2, 3}},
		{name: "negatives", list1: []int{-10, -3, 0}, list2: []int{-4, -3, 2}, expect: []int{-10, -4, -3, -3, 0, 2}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			gotHead := mergeTwoLists(buildList(tc.list1), buildList(tc.list2))
			got := listToSlice(gotHead)
			if len(got) != len(tc.expect) {
				t.Fatalf("mergeTwoLists() = %v, want %v", got, tc.expect)
			}
			for i := range got {
				if got[i] != tc.expect[i] {
					t.Fatalf("mergeTwoLists() = %v, want %v", got, tc.expect)
				}
			}
		})
	}
}

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	out := make([]int, 0)
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}
