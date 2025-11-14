package mergetwosortedlists

import "fmt"

const debugMergeTwoLists = false

// ListNode defines a singly linked list node used by LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merges two sorted linked lists into one sorted list and returns its head.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if debugMergeTwoLists {
		fmt.Printf("merge start list1=%v list2=%v\n", dumpList(list1), dumpList(list2))
	}

	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		if debugMergeTwoLists {
			fmt.Printf("attach %d\n", current.Next.Val)
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	if debugMergeTwoLists {
		fmt.Printf("merge result=%v\n", dumpList(dummy.Next))
	}
	return dummy.Next
}

func dumpList(head *ListNode) []int {
	if !debugMergeTwoLists {
		return nil
	}
	out := make([]int, 0)
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}
