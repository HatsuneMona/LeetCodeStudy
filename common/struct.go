package common

type ListNode struct {
    Val  int
    Next *ListNode
}

// Node Definition for a Node.
type Node struct {
    Val      int
    Children []*Node
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
