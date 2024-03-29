/**
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。



 示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]


 示例 2：


输入：root = [1]
输出：[[1]]


 示例 3：


输入：root = []
输出：[]




 提示：


 树中节点数目在范围 [0, 2000] 内
 -1000 <= Node.val <= 1000


 Related Topics 树 广度优先搜索 二叉树 👍 1487 👎 0

*/

package leetcode

import (
	. "leetcode/common"
)

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levelResult := [][]int{{root.Val}}

	if root.Left != nil {
		levelResult = append(levelResult, levelOrder(root.Left)...)
	}

	if root.Right != nil {
		for level, result := range levelOrder(root.Right) {
			if level+1 < len(levelResult) {
				levelResult[level+1] = append(levelResult[level+1], result...)
			} else {
				levelResult = append(levelResult, result)
			}
		}
	}

	return levelResult

}

// leetcode submit region end(Prohibit modification and deletion)
