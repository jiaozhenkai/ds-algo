package main

import (
	"fmt"
	"slices"
)

type BinaryTreeNode struct {
	rChild *BinaryTreeNode
	lChild *BinaryTreeNode
	value  int
}

// 递归中序
func inOrderRer(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	inOrderRer(root.lChild)
	fmt.Printf("%d  ", root.value)
	inOrderRer(root.rChild)
}

// 层序
func levelOrder(root *BinaryTreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}

	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node.value)
		if node.lChild != nil {
			queue = append(queue, node.lChild)
		}
		if node.rChild != nil {
			queue = append(queue, node.rChild)
		}
	}

	return result
}

// 层序
func levelOrder1(root *BinaryTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}

	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		tmp := []int{}
		nextLevel := []*BinaryTreeNode{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			tmp = append(tmp, node.value)
			if node.lChild != nil {
				nextLevel = append(nextLevel, node.lChild)
			}
			if node.rChild != nil {
				nextLevel = append(nextLevel, node.rChild)
			}
		}
		result = append(result, tmp)
		queue = nextLevel
	}

	return result
}

// 迭代前序
func preOrderIter(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	stack := []*BinaryTreeNode{root}

	for len(stack) > 0 {
		// [stack bottom, .... stack top]
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", node.value)
		if node.rChild != nil {
			stack = append(stack, node.rChild)
		}
		if node.lChild != nil {
			stack = append(stack, node.lChild)
		}
	}
}

// 迭代中序
/*
先把 root 节点的所有左子节点全部入队。
注意 for 停止条件
*/
func inOrderIter(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	// [stack bottom, .... stack top]
	var stack []*BinaryTreeNode
	// current root node
	curr := root
	for curr != nil || len(stack) > 0 {
		// left child to stack all.
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.lChild
		}

		curr = stack[len(stack)-1]
		fmt.Printf("%d  ", curr.value)
		stack = stack[:len(stack)-1]
		curr = curr.rChild
	}
}

// 迭代后序
/*
双栈思想（反向遍历）：
root 先入栈+出栈，放入 result 中。
leftChild rightClild 依次入栈，pop 时，rightChild 先出，leftChild 后出。
则整个顺序为：root, rightChild, leftChild; 这个顺序刚好和想要的顺序相反，reverse 一下。
*/
func postOrderIter(root *BinaryTreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*BinaryTreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.value)
		if node.lChild != nil {
			stack = append(stack, node.lChild)
		}
		if node.rChild != nil {
			stack = append(stack, node.rChild)
		}
	}
	slices.Reverse(result)
	return result
}

func postOrderIter1(root *BinaryTreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*BinaryTreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// result = append([]int{node.value}, result...)
		result = slices.Insert(result, 0, node.value)
		if node.lChild != nil {
			stack = append(stack, node.lChild)
		}
		if node.rChild != nil {
			stack = append(stack, node.rChild)
		}
	}

	return result
}

func main() {
	r := BinaryTreeNode{
		value:  10,
		rChild: nil,
		lChild: nil,
	}
	l := BinaryTreeNode{
		value:  11,
		rChild: nil,
		lChild: nil,
	}
	root := &BinaryTreeNode{
		value:  0,
		rChild: &r,
		lChild: &l,
	}

	root1 := &BinaryTreeNode{
		value:  3,
		lChild: &BinaryTreeNode{value: 9},
		rChild: &BinaryTreeNode{
			value:  20,
			lChild: &BinaryTreeNode{value: 15},
			rChild: &BinaryTreeNode{value: 7},
		},
	}
	test := postOrderIter1(root)
	fmt.Println(test)
	fmt.Println("----")
	tet := postOrderIter1(root1)
	fmt.Println(tet)

}
