package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
有问题
*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	tmp := dummy
	first := head
	second := head.Next

	for tmp.Next != nil && tmp.Next.Next != nil {
		tmp.Next = second
		first.Next = second.Next
		second.Next = first
		second = first.Next.Next
		first = first.Next

	}
	return dummy.Next

}

func main() {
	head := ConstructListFromSlice([]int{1, 2, 3, 4})
	res := swapPairs(head)
	PrintList(res)
}
