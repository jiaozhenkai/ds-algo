package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
141 题的变形

设：
a = 头节点到入环点的距离（a 可以认为指向了入环点）
b = 入环点到相遇点的距离（b 可以认为指向了相遇点）
C = 环长

相遇时：
slow 走了：a + b
fast 走了：a + b + nC（以 b 点为 fast 的起点为准，多绕了 n 圈）
fast 走的步数 = 2 × slow 走的步数

```
a + b + nC = 2(a + b)
nC = a + b
a = nC - b // 某一圈从相遇点开始到入环点的距离，但是nC -b 不容易理解，看下一步；
a = (n-1)C + (C - b) //从头节点出发走 a 步，等于从相遇点出发走 (n-1) 整圈再走 C - b 步，两者都会到达入环点。
```
*/
func detectCycle1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			slow = head // 随意让一个指针指向头节点。
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

func detectCycle(head *ListNode) *ListNode {
	m := map[*ListNode]struct{}{}

	for h := head; h != nil; h = h.Next {
		if _, ok := m[h]; ok {
			return h
		} else {
			m[h] = struct{}{}
		}
	}
	return nil
}
