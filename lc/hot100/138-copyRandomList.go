package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
拼接+拆分链表
refer: https://leetcode.cn/problems/copy-list-with-random-pointer/solutions/2361362/138-fu-zhi-dai-sui-ji-zhi-zhen-de-lian-b-6jeo/?envType=study-plan-v2&envId=top-100-liked

有问题，后面再重新写吧
*/
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 拼接
	for h := head; h != nil; h = h.Next {
		tmp := h.Next
		node := new(Node)
		node.Val = h.Val
		h.Next = node
		node.Next = tmp
	}

	// 构造 random 指针
	for h := head; h != nil; h = h.Next.Next {
		node := h.Next //  copy 出来的节点
		if h.Random != nil {
			node.Random = h.Random.Next
		}
	}
	res := head.Next
	// 拆分
	for h := head; h != nil; h = h.Next {
		node := h.Next //  copy 出来的节点
		h.Next = node.Next
		node.Next = h.Next.Next
	}

	return res
}

/*
第一次遍历，搞一个新 node，key 为旧 node，value 为新 node
第二次遍历原链表，不断用 map[h] 获取新 node, 此时新 node 的 next 就是 map[h.next]
random 就是 map[h.random]
*/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodeMap := make(map[*Node]*Node) //
	for h := head; h != nil; h = h.Next {
		tmp := new(Node)
		tmp.Val = h.Val
		nodeMap[h] = tmp

	}

	for h := head; h != nil; h = h.Next {
		v, _ := nodeMap[h]
		v.Next = nodeMap[h.Next]
		v.Random = nodeMap[h.Random]
	}

	return nodeMap[head]
}
