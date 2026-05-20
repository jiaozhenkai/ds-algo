package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{}
	curr := root
	res := []int{}
	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		curr = curr.Right
	}
	return res
}

func main() {
	r := TreeNode{
		Val:   10,
		Right: nil,
		Left:  nil,
	}
	l := TreeNode{
		Val:   11,
		Right: nil,
		Left:  nil,
	}
	root := &TreeNode{
		Val:   0,
		Right: &r,
		Left:  &l,
	}

	//root1 := &TreeNode{
	//	Val:  3,
	//	Left: &TreeNode{Val: 9},
	//	Right: &TreeNode{
	//		Val:   20,
	//		Left:  &TreeNode{Val: 15},
	//		Right: &TreeNode{Val: 7},
	//	},
	//}
	test := inorderTraversal(root)
	fmt.Println(test)
	//fmt.Println("----")
	//tet := inorderTraversal(root1)
	//fmt.Println(tet)
}
