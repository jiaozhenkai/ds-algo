package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
递归
*/

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return recur(head, nil)
}
func recur(curr, pre *ListNode) *ListNode {
	if curr == nil {
		return pre
	}
	res := recur(curr.Next, curr)
	curr.Next = pre
	return res
}

/*
双指针
cur 指向头节点，res 刚开始指向空
每次将 cur 指向的节点的值拿出来构造一个tmp节点，tmp 节点的next指向是 res 节点，然后更新 res 为 tmp， 然后更新 cur .
*/
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var res *ListNode

	for cur != nil {
		tmp := &ListNode{}
		tmp.Val = cur.Val
		tmp.Next = res
		res = tmp
		cur = cur.Next
	}

	return res
}

/*
借助数组反转
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nodes := []int{}
	for h := head; h != nil; h = h.Next {
		nodes = append(nodes, h.Val)
	}
	h := &ListNode{}
	res := h
	for i := len(nodes) - 1; i >= 0; i-- {
		h.Val = nodes[i]
		if i != 0 {
			h.Next = &ListNode{}
		}

		h = h.Next
	}
	return res
}
