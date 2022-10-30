/**
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个
节点也可以是它自己的祖先）。”

 例如，给定如下二叉搜索树: root = [6,2,8,0,4,7,9,null,null,3,5]





 示例 1:

 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。


 示例 2:

 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。



 说明:


 所有节点的值都是唯一的。
 p、q 为不同节点且均存在于给定的二叉搜索树中。


 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 956 👎 0

*/

package leetcode

import (
	. "../../../common"
	"../../../tools"
	"fmt"
)

func Q235Main() {
	testTree := []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}

	treeNodeList := tools.CreateTreeNodeListByIntValList(testTree)

	tree := tools.CreateTree(treeNodeList)

	result := lowestCommonAncestor(tree, treeNodeList[1], treeNodeList[4])

	fmt.Printf("result common Node is %v \n", result.Val)
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if root.Val > p.Val && root.Val > q.Val {
		// p  q  都在当前节点的左侧
		if left := lowestCommonAncestor(root.Left, p, q); left != nil {
			return left
		}
	}

	if root.Val < p.Val && root.Val < q.Val {
		// p  q  都在当前节点的右侧
		if right := lowestCommonAncestor(root.Right, p, q); right != nil {
			return right
		}
	}

	return root
}

// leetcode submit region end(Prohibit modification and deletion)
