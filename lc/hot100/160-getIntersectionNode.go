package main

import (
	"fmt"
)

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

/*
假设 A 中有 a 个节点，B 中有 b 个节点切 a<b; 相同的节点个数为 c 个；
A 走完自己再走 B，B 走完自己再走 A， 则 A 走了 a+（b-c）个，B 走了b+(a-c) 个，边走边判断，如果相等就返回。本质上是因为 AB 速率一样让他们走相同的路程相遇就会找到交点;或者都遇到 nil 没有交点

refer: https://leetcode.cn/problems/intersection-of-two-linked-lists/solutions/12624/intersection-of-two-linked-lists-shuang-zhi-zhen-l/?envType=study-plan-v2&envId=top-100-liked
*/
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	hA := headA
	hB := headB

	// 如果 A 和 B 的长度相同且没有交点，两个 nil 相等直接跳出循环。
	for hA != hB {
		if hA == nil {
			hA = headB
		} else {
			hA = hA.Next
		}
		if hB == nil {
			hB = headA
		} else {
			hB = hB.Next
		}
	}
	return hA
}

/*
本质上因为 AB 速率一样，想办法让 A，B 走相同的路程，边走边判断，如果相等则返回，如果不等则继续往下走都遇到 nil 没有交点。
假设 A 中有 a 个节点，B 中有 b 个节点切 a<b; B 比 A 多了 b-a 个节点，先让 B 先走 b-a 个节点，然后同时走，边走边判断。
*/
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	lA := 0
	lB := 0

	for head := headA; head != nil; head = head.Next {
		lA++
	}

	for head := headB; head != nil; head = head.Next {
		lB++
	}

	var hA *ListNode = headA
	var hB *ListNode = headB

	if lA > lB {
		tmp := lA - lB
		for tmp > 0 {
			hA = hA.Next
			tmp--
		}
	} else {
		tmp := lB - lA
		for tmp > 0 {
			hB = hB.Next
			tmp--
		}
	}

	for hA != hB {
		hA = hA.Next
		hB = hB.Next
	}

	return hA

}

/*
headA 中的节点放到 map 中，遍历 headB 如果有则返回对应节点。
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for head := headA; head != nil; head = head.Next {
		m[head] = struct{}{}
	}
	for head := headB; head != nil; head = head.Next {
		if _, ok := m[head]; ok {
			return head
		}
	}
	return nil
}

func main() {
	A := ListNode{Val: 4, Next: &ListNode{Val: 2, Next: nil}}
	B := ListNode{Val: 5, Next: &ListNode{Val: 1, Next: nil}}
	fmt.Println(getIntersectionNode2(&A, &B))
}
