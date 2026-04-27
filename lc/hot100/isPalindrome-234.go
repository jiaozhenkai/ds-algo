package main

import "fmt"

/*
可以先反转后半部分链表，再遍历。
找中间的元素有两种办法，1. 统计个数找到 2. 快慢指针。

*/

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	midNode := midNode1(head)
	second := revList(midNode)
	for s := second; s != nil; s = s.Next {
		if head.Val != s.Val {
			return false
		}
		head = head.Next
	}

	return true
}

func revList(head *ListNode) *ListNode {
	var res *ListNode

	for h := head; h != nil; h = h.Next {
		tmp := new(ListNode)
		tmp.Val = h.Val
		tmp.Next = res
		res = tmp
	}
	return res
}
func midNode1(head *ListNode) *ListNode {
	total := 0
	for h := head; h != nil; h = h.Next {
		total++
	}
	start := 0
	//if total%2 == 0 {
	//	start = total / 2
	//} else {
	//	start = total/2 + 1
	//}
	start = total / 2
	var second *ListNode = head
	for i := 0; i < start; i++ {
		second = second.Next
	}
	return second
}

func midNode2(head *ListNode) *ListNode {
	first := head
	second := head

	for second != nil && second.Next != nil { // 注意这里的条件；奇数个节点直接取到中间位置，偶数个节点需要考虑偏左还是偏右，这里是偏右的；偏左的话应该是 second.Next != nil && second.Next.Next != nil
		first = first.Next
		second = second.Next.Next
	}
	return first
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}
	fmt.Println(isPalindrome(head))
}
