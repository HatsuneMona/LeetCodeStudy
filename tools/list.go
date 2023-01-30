package tools

import (
	"fmt"
	. "leetcode/common"
)

// ListNodeCreateByArr 通过arr数组生成 ListNode
func ListNodeCreateByArr(input []int) (output *ListNode) {
	if len(input) < 1 {
		return nil
	}

	pNode := &ListNode{
		Val:  input[0],
		Next: nil,
	}
	output = pNode

	if len(input) > 1 {
		for _, val := range input[1:] {
			pNode.Next = &ListNode{
				Val:  val,
				Next: nil,
			}
			pNode = pNode.Next
		}
	}

	return
}

// ListNodeCheckEqual 检查 ListNode 是否相等
func ListNodeCheckEqual(l1, l2 *ListNode) bool {
	if fmt.Sprintf("%v", l1.ToArr()) == fmt.Sprintf("%v", l1.ToArr()) {
		return true
	} else {
		return false
	}
}
