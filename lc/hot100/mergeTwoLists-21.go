package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	res := new(ListNode)
	h := res
	l1 := list1
	l2 := list2
	for l1 != nil && l2 != nil {
		tmp := new(ListNode)
		if l1.Val >= l2.Val {
			tmp.Val = l2.Val
			h.Next = tmp
			l2 = l2.Next
		} else {
			tmp.Val = l1.Val
			h.Next = tmp
			l1 = l1.Next
		}
		h = h.Next
	}

	if l1 == nil && l2 == nil {
		return res
	}

	if l1 != nil {
		h.Next = l1
	}
	if l2 != nil {
		h.Next = l2
	}
	// 这样写也行
	//if l1 != nil {
	//	for l1 != nil {
	//		tmp := new(ListNode)
	//		tmp.Val = l1.Val
	//		h.Next = tmp
	//		h = h.Next
	//		l1 = l1.Next
	//	}
	//}
	//
	//if l2 != nil {
	//	for l2 != nil {
	//		tmp := new(ListNode)
	//		tmp.Val = l2.Val
	//		h.Next = tmp
	//		h = h.Next
	//		l2 = l2.Next
	//	}
	//}

	return res.Next
}
