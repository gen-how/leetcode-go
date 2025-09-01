package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{}
	tail := head
	for (l1 != nil) || (l2 != nil) || (carry > 0) {
		d1 := 0
		if l1 != nil {
			d1 = l1.Val
			l1 = l1.Next
		}
		d2 := 0
		if l2 != nil {
			d2 = l2.Val
			l2 = l2.Next
		}
		sum := d1 + d2 + carry
		carry = sum / 10
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
	}
	return head.Next
}
