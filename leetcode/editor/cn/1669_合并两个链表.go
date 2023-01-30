/**
给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。

 请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。

 下图中蓝色边和节点展示了操作后的结果：

 请你返回结果链表的头指针。



 示例 1：




输入：list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
输出：[0,1,2,1000000,1000001,1000002,5]
解释：我们删除 list1 中下标为 3 和 4 的两个节点，并将 list2 接在该位置。上图中蓝色的边和节点为答案链表。


 示例 2：


输入：list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,10000
03,1000004]
输出：[0,1,1000000,1000001,1000002,1000003,1000004,6]
解释：上图中蓝色的边和节点为答案链表。




 提示：


 3 <= list1.length <= 10⁴
 1 <= a <= b < list1.length - 1
 1 <= list2.length <= 10⁴


 Related Topics 链表 👍 73 👎 0

*/

package leetcode

import (
	"fmt"
	. "leetcode/common"
	"leetcode/tools"
)

type input struct {
	list1 *ListNode
	a     int
	b     int
	list2 *ListNode
}

func Q1669Main() {
	testCase := []struct {
		in  input
		out *ListNode
	}{
		{
			in: input{
				list1: tools.ListNodeCreateByArr([]int{0, 1, 2, 3, 4, 5}),
				a:     3,
				b:     4,
				list2: tools.ListNodeCreateByArr([]int{1000000, 1000001, 1000002}),
			},
			out: tools.ListNodeCreateByArr([]int{0, 1, 2, 1000000, 1000001, 1000002, 5}),
		},
		{
			in: input{
				list1: tools.ListNodeCreateByArr([]int{0, 1, 2, 3, 4, 5, 6}),
				a:     2,
				b:     5,
				list2: tools.ListNodeCreateByArr([]int{1000000, 1000001, 1000002, 1000003, 1000004}),
			},
			out: tools.ListNodeCreateByArr([]int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}),
		},
		//{
		//	in: input{
		//		list1: tools.ListNodeCreateByArr([]int{0, 1, 2}),
		//		a:     2,
		//		b:     5,
		//		list2: tools.ListNodeCreateByArr([]int{1000000, 1000001}),
		//	},
		//	out: tools.ListNodeCreateByArr([]int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}),
		//},
	}

	for _, c := range testCase {
		if !tools.ListNodeCheckEqual(c.out, mergeInBetween(c.in.list1, c.in.a, c.in.b, c.in.list2)) {
			fmt.Printf("error! %v", c)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	var head, end *ListNode
	pNode := list1
	for i := 0; i <= b; i++ {
		if i == a-1 {
			head = pNode
		}
		pNode = pNode.Next
	}
	end = pNode

	pNode = list2
	for pNode.Next != nil {
		pNode = pNode.Next
	}
	head.Next = list2
	pNode.Next = end

	return list1
}

//leetcode submit region end(Prohibit modification and deletion)
