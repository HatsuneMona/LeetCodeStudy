/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package Q206

import "20.leecode/leetcode"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *leetcode.ListNode) *leetcode.ListNode {
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
