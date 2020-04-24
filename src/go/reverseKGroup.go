package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := dummy
	end := dummy

	for end != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = reverse(start)
		start.Next = next
		pre = start

		end = pre
	}
	return dummy.Next
}

func reverse(start *ListNode) *ListNode {
	var pre *ListNode
	cur := start
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	ans := reverseKGroup(l1, 3)
	for ans != nil {
		println(ans.Val)
		ans = ans.Next
	}
}
