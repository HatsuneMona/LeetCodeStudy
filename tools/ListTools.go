/**
 * @Author HatsuneMona
 * @Date  2021-01-30 09:20
 * @Description 方便创建一些数据结构而产生的工具
 **/
package tools

import (
	. "20.leecode/leetcodeType"
)

/**
 * @作者      HatsuneMona
 * @描述      该工具可以根据数组，生成对应的链表数据结构。
 * @编写日期   2021/1/30 09:20
 * @参数      []int
 * @返回值     *ListNode
**/
func ArrToList(src []int) *ListNode {
	if len(src) == 0 {
		return nil
	}
	dst := new(ListNode)
	dst.Val = src[0]
	p := dst
	for i := 1; i < len(src); i++ {
		t := new(ListNode)
		t.Val = src[i]
		p.Next = t
		p = p.Next
	}
	return dst
}

/**
 * @作者      HatsuneMona
 * @描述      该工具，可以根据链表数据结构，将数据整理为数组
 * @编写日期   2021/1/30 09:32
 * @参数      *ListNode
 * @返回值     []int
**/
func ListToArr(src *ListNode) []int {
	length := 0
	for node := src; node != nil; node = node.Next {
		length++
	}
	dst := make([]int, length)
	node := src
	for i := 0; i < length; i++ {
		dst[i] = node.Val
		node = node.Next
	}
	return dst
}
