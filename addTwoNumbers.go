package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry := 0
	curr := dummy
	// while there are still values in l1 or l2
	for l1 != nil || l2 != nil {
		sum := carry

		// if l1 is not null, add value to sum
		// and iterate to the next value
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// if l2 is not null, add the value to sum,
		// and iterate to the next value
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// sum / 10 will tell us what carry is
		// (should only ever be 1 or 0)
		carry = sum / 10
		// add the value of the current sum to the return list
		curr.Next = &ListNode{Val: sum % 10}
		// iterate the current return node
		curr = curr.Next
	}

	// if we have iterated both lists and there is still a carry,
	// we need to add this to the list as well
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummy
}
