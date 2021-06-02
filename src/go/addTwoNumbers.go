package src

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{Val: 0, Next: nil}
	l3 := ret
	carry := 0
	x := 0
	y := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		} else {
			y = 0
		}
		s := x + y + carry
		carry = s / 10
		l3.Next = &ListNode{s % 10, nil}
		l3 = l3.Next
	}
	if carry > 0 {
		l3.Next = &ListNode{Val: 1, Next: nil}
	}
	return ret.Next
}
