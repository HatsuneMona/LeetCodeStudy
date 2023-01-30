package tools

import (
	. "leetcode/common"
)

// CreateTree 根据 LeetCode 输入参数的格式，生成一个 二叉树。
// 例：root = [6,2,8,0,4,7,9,null,null,3,5]
// 可以生成：
//
//	            6
//	        /       \
//	    2               8
//	  /   \           /   \
//	0       4       7       9
//	      /   \
//	    3       5
//
// 其中 返回值 root treeNode = 6
func CreateTree(nodeList []*TreeNode) *TreeNode {

	if len(nodeList) == 0 {
		return nil
	}

	root := nodeList[0]
	nodeList = nodeList[1:]

	fatherStack := []*TreeNode{root}

	for len(fatherStack) != 0 {
		newFatherStack := make([]*TreeNode, 0)
		for _, node := range fatherStack {
			if len(nodeList) != 0 {
				node.Left = nodeList[0]
				newFatherStack = append(newFatherStack, node.Left)
				nodeList = nodeList[1:]
			}
			if len(nodeList) != 0 {
				node.Right = nodeList[0]
				newFatherStack = append(newFatherStack, node.Right)
				nodeList = nodeList[1:]
			}
		}
		fatherStack = newFatherStack
	}

	return root
}

// CreateTreeNodeListByIntValList 将 值可以为nil 的 int 数组，转换为 *TreeNode 节点数组
// 例：testTree := []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}
//
//	输出为：treeNodeList := []*TreeNode{&TreeNode{Val: 6}, ... , nil, nil, &TreeNode{Val: 3}, ...}
func CreateTreeNodeListByIntValList(intValList []any) []*TreeNode {

	result := make([]*TreeNode, len(intValList))

	for i, intVal := range intValList {
		if intVal == nil {
			continue
		}
		if val, ok := intVal.(int); ok {
			result[i] = &TreeNode{Val: val}
		}
	}

	return result

}
