/*
Q206_反转一个单链表.反转一个单链表。

示例:
	输入: 1->2->3->4->5->NULL
	输出: 5->4->3->2->1->NULL

进阶:
	你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

Related Topics 链表
👍 1483 👎 0*/
package Q206_反转一个单链表

import . "20.leecode/leetcodeType"

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	revList := *head
	pRevList := &revList
	revList.Next = nil
	for node := head.Next; node != nil; node = node.Next {
		tempNode := *node
		tempNode.Next = pRevList
		pRevList = &tempNode
	}
	return pRevList
}
