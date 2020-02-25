package solution1

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var temp ListNode
	for l1.Next != nil {
		for l2.Next != nil {
			if l1.Val > l2.Val {
				temp.Val = l2.Val
				l2 = l2.Next
				temp = temp.Next
				//temp.Next = mergeTwoLists(l1, l2)
			} else {
				temp.Val = l1.Val
				l1 = l1.Next
				temp = temp.Next
				//temp.Next = mergeTwoLists(l2, l1)
			}
		}
	}
	return *ListNode(temp)
}
