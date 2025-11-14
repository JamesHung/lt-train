package linkedlistcycle

import "testing"

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name  string
		build func() *ListNode
		want  bool
	}{
		{
			name: "example cycle",
			build: func() *ListNode {
				n1 := &ListNode{Val: 3}
				n2 := &ListNode{Val: 2}
				n3 := &ListNode{Val: 0}
				n4 := &ListNode{Val: -4}
				n1.Next = n2
				n2.Next = n3
				n3.Next = n4
				n4.Next = n2
				return n1
			},
			want: true,
		},
		{
			name: "cycle length two",
			build: func() *ListNode {
				n1 := &ListNode{Val: 1}
				n2 := &ListNode{Val: 2}
				n1.Next = n2
				n2.Next = n1
				return n1
			},
			want: true,
		},
		{
			name:  "single node no cycle",
			build: func() *ListNode { return &ListNode{Val: 1} },
			want:  false,
		},
		{
			name:  "empty",
			build: func() *ListNode { return nil },
			want:  false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := hasCycle(tc.build())
			if got != tc.want {
				t.Fatalf("hasCycle() = %v, want %v", got, tc.want)
			}
		})
	}
}
