package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
双指针：无需获取长度。

second 指针先跳 n 次，这样 second 和 first 指针之间隔了 n-1 个节点。这样在 second 跳到链表末尾 nil 时，first 刚好是倒数第 n 个节点；为了好删除，可以让 first 指向 dummy 节点。
*/
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	first := dummy
	second := head

	for i := 0; i < n; i++ {
		second = second.Next
	}

	for second != nil {
		first = first.Next
		second = second.Next
	}

	first.Next = first.Next.Next

	return dummy.Next
}

/*
删除一个链表的节点，比如  a b c 删除 b 要保存 b 的前继节点，让前继节点的 next 指针指向 b 的 next;

可以先遍历取到长度，变成正序删除，稍微注意下正序应该在哪里停止就行。
注意 dummy 节点（也就是空的 head 节点）的应用。
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	total := 0

	for h1 := head; h1 != nil; h1 = h1.Next {
		total++
	}
	n = total - n

	for i := 0; i < n; i++ {
		pre = pre.Next
	}

	pre.Next = pre.Next.Next

	return dummy.Next
}

func main() {
	l1 := ConstructListFromSlice([]int{1, 2, 3, 4, 5})
	res1 := removeNthFromEnd(l1, 2)
	PrintList(res1)

	l2 := ConstructListFromSlice([]int{1})
	res2 := removeNthFromEnd(l2, 1)
	PrintList(res2)
}
