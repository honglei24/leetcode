//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := dummy
	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		pre.Next = second
		first.Next = second.Next
		second.Next = first

		pre = first
		head = first.Next
	}
	return dummy.Next
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	ans := swapPairs(l1)
	for ans != nil {
		println(ans.Val)
		ans = ans.Next
	}
}
