package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	res := new(ListNode)
	h := res
	sum := 0

	for l1 != nil || l2 != nil {
		l1V, l2V := 0, 0
		if l1 != nil {
			l1V = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2V = l2.Val
			l2 = l2.Next
		}
		sum = l1V + l2V + flag
		sum, flag = sum%10, sum/10

		tmp := new(ListNode)
		tmp.Val = sum
		h.Next = tmp
		h = h.Next
	}

	if flag == 1 {
		tmp := new(ListNode)
		tmp.Val = 1
		h.Next = tmp
	}
	return res.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0

	h1 := l1
	h2 := l2

	res := new(ListNode)
	h := res

	for h1 != nil && h2 != nil {
		sum := h1.Val + h2.Val + flag
		var val int
		if sum >= 10 {
			val = sum - 10
			flag = 1
		} else {
			val = sum
			flag = 0
		}
		tmp := new(ListNode)
		tmp.Val = val
		h.Next = tmp

		h = h.Next
		h1 = h1.Next
		h2 = h2.Next
	}

	if h1 == nil && h2 == nil {
		if flag == 1 {
			tmp := new(ListNode)
			tmp.Val = 1
			h.Next = tmp
		}
		return res.Next
	}

	if h1 != nil {
		for h1 != nil {
			sum := h1.Val + flag

			var val int
			if sum >= 10 {
				val = sum - 10
				flag = 1
			} else {
				val = sum
				flag = 0
			}

			tmp := new(ListNode)
			tmp.Val = val
			h.Next = tmp
			h1 = h1.Next
			h = h.Next
		}

	}

	if h2 != nil {
		for h2 != nil {
			sum := h2.Val + flag
			var val int
			if sum >= 10 {
				val = sum - 10
				flag = 1
			} else {
				val = sum
				flag = 0
			}
			tmp := new(ListNode)
			tmp.Val = val
			h.Next = tmp
			h2 = h2.Next
			h = h.Next
		}
	}

	if flag == 1 {
		tmp := new(ListNode)
		tmp.Val = 1
		h.Next = tmp
	}

	return res.Next
}
func main() {
	l1 := ConstructListFromSlice([]int{2, 4, 9})
	l2 := ConstructListFromSlice([]int{5, 6, 4, 9})
	res := addTwoNumbers1(l1, l2)

	PrintList(res)
}
