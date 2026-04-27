package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
fast 每次走 2 步
slow 每次走 1 步
假设环的长度为 C，当 slow 进入环时，fast 在 slow 前面的距离假设为 d; 那么 fast 要追上 slow 的距离为 C-d
fast-slow = 1，也就是每次比 slow 多走一步，这样从 C-d, C-d-1, C-d-2 ...... 1,0， 最终会追上；

数学上来说，相遇条件：fast 和 slow 在环上同一位置：
fast的位置 ≡ slow的位置 (mod C) 等价于 (f₀ + 2t) mod C = (s₀ + t) mod C， f0是slow 进入环时 fast 在环上的当前位置， s0 slow 进入环时在环上的起始位置;所以：

(f₀ + 2t) - (s₀ + t) ≡ 0 (mod C)
f₀ + 2t - s₀ - t ≡ 0 (mod C)
f₀ - s₀ + t ≡ 0 (mod C)
t ≡ -(f₀ - s₀) (mod C)
t ≡ (s₀ - f₀) (mod C) （含义：t 除以 C 的余数等于 (s₀ - f₀) mod C）

t 的最小正整数解
t = ((s₀ - f₀) mod C + C) mod C （加一个 C 再取模是为了处理 s₀ - f₀ 为负数的情况。）
t 解了出来。
*/
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false
}

/*
如果有环，则必然重复遍历到某一个已经遍历过的节点，所以把节点放到 map 中保存。
*/
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{})

	for h := head; h != nil; h = h.Next {
		if _, ok := m[h]; ok {
			return true
		} else {
			m[h] = struct{}{}
		}
	}
	return false
}
