/**
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。







 示例 1：


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]


 示例 2：


输入：head = [1,2]
输出：[2,1]


 示例 3：


输入：head = []
输出：[]




 提示：


 链表中节点的数目范围是 [0, 5000]
 -5000 <= Node.val <= 5000




 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

 Related Topics 递归 链表 👍 2807 👎 0

*/

package leetcode

import (
	"fmt"
	. "leetcode/common"
)

func Q206Main() {
	root := &ListNode{
		Val:  0,
		Next: nil,
	}
	testCase := root

	for i := 1; i <= 2; i++ {
		testCase.Next = &ListNode{
			Val:  i,
			Next: nil,
		}
		testCase = testCase.Next
	}

	result := reverseList(root.Next)

	for result != nil {
		fmt.Printf(" %v ", result.Val)
		result = result.Next
	}
	fmt.Println()
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 链表有[0, 1]个元素
	if head == nil || head.Next == nil {
		return head
	}

	a, b := head, head.Next
	var c *ListNode

	// 防止 a <-> b 的问题
	a.Next = nil

	for b != nil {
		// c: 防断链
		c = b.Next

		// a -> b -> c 改为 a <- b   c -> ...
		b.Next = a

		// 操作窗口 后移一位
		a, b = b, c
	}

	return a
}

// leetcode submit region end(Prohibit modification and deletion)
