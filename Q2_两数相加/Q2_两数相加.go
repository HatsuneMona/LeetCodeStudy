/*
2. 两数相加
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例 1：									示例 2：
	输入：l1 = [2,4,3], l2 = [5,6,4]			输入：l1 = [0], l2 = [0]
	输出：[7,0,8]							输出：[0]
	解释：342 + 465 = 807.

示例 3：
	输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	输出：[8,9,9,9,0,0,0,1]

提示：
	每个链表中的节点数在范围 [1, 100] 内
	0 <= Node.val <= 9
	题目数据保证列表表示的数字不含前导零
*/

package main

import . "20.leecode/leetcode"

func main() {

}

/*
type ListNode struct {
    Val  int
    Next *ListNode
}*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//假设l1是最长的
	node2 := l2
	node1 := l1
	for ; node1 != nil; node1 = node1.Next {
		if node2 != nil {
			node1.Val += node2.Val
			node2 = node2.Next
		}
		if node1.Val > 9 { //检查进位
			node1.Val -= 10
			if node1.Next != nil { //检查是否有空间进位
				node1.Next.Val++
			} else { //没有空间进位的情况下，创造空间
				node1.Next = &ListNode{Val: 1, Next: nil}
			}
		}
		if node1.Next == nil {
			node1.Next = node2
			break
		}
	}
	return l1
}
