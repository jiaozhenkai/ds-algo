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
	pre := dummy // pre 含义：指向待交换的两个节点的前一个，比如 abcd, 交换 cd 时指向 a

	for pre.Next != nil && pre.Next.Next != nil {
		first := pre.Next
		second := pre.Next.Next

		//tmp := second.Next
		pre.Next = second
		first.Next = second.Next
		second.Next = first

		pre = pre.Next.Next
		// pre = first
	}
	return dummy.Next

}

func main() {
	head := ConstructListFromSlice([]int{1, 2, 3, 4})
	res := swapPairs(head)
	PrintList(res)
}
