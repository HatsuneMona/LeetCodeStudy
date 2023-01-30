package common

type ListNode struct {
	Val  int
	Next *ListNode
}

// ToArr ListNode 转换为 []int 数组
func (l *ListNode) ToArr() (arr []int) {
	pNode := l
	start := l // 如果遇到 循环链表 终止
	for pNode != nil && pNode.Next != start {
		arr = append(arr, pNode.Val)
		pNode = pNode.Next
	}
	return
}
