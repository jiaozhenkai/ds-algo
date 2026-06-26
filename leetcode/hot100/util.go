package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConstructListFromSlice(s []int) *ListNode {
	res := new(ListNode)
	h := res
	for _, v := range s {
		tmp := new(ListNode)
		tmp.Val = v
		h.Next = tmp
		h = h.Next
	}
	return res.Next
}

func PrintList(node *ListNode) {
	for h := node; h != nil; h = h.Next {
		fmt.Printf("%d ", h.Val)
	}
	fmt.Println()
}
