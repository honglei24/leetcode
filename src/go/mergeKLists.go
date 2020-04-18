//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	length := len(lists)
	interval := 1
	for interval < length {
		for i := 0; i < length-interval; {
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
			i = i + interval*2
		}
		interval *= 2
	}
	return lists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{
		Val:  0,
		Next: nil,
	}
	p := ans
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			p.Next = l2
			l2 = l2.Next
		} else {
			p.Next = l1
			l1 = l1.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	} else if l2 == nil {
		p.Next = l1
	}
	return ans.Next
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}
	arr := []*ListNode{l1, l2, l3}
	ans := mergeKLists(arr)
	for ans.Next != nil {
		println(ans.Val)
		ans = ans.Next
	}
}
