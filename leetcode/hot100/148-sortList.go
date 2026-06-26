package main

import "sort"

/*
归并排序
找中点，断链，归并
*/
func sortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 找中点
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 右半部分链表 并且断开
	tmp := slow.Next
	slow.Next = nil

	left := sortList1(head)
	right := sortList1(tmp)

	res := new(ListNode)
	cur := res

	//归并
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}

	if left == nil {
		cur.Next = right
	}

	if right == nil {
		cur.Next = left
	}

	return res.Next
}

/*
放到数组中排序再搞一个新的
*/
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	list := []int{}
	for h := head; h != nil; h = h.Next {
		list = append(list, h.Val)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	res := new(ListNode)
	cur := res
	for _, v := range list {
		tmp := new(ListNode)
		tmp.Val = v
		cur.Next = tmp
		cur = cur.Next
	}

	return res.Next
}
