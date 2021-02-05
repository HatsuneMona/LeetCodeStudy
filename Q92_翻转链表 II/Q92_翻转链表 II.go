/*
92. 反转链表 II
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

示例:
	输入: 1->2->3->4->5->NULL, m = 2, n = 4
	输出: 1->4->3->2->5->NULL
说明:
	1 ≤ m ≤ n ≤ 链表长度。
*/
package main

import . "20.leecode/leetcodeType"
import "20.leecode/tools"

func main() {
	test := tools.ArrToList([]int{3, 5})
	need := tools.ArrToList([]int{5, 3})
	tools.IsListEqual(reverseBetween(test, 1, 2), need)
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//特殊情况处理
	if head == nil { //输入的链表为空
		return head
	}
	if n-m < 1 { //m == n时无需翻转任何内容
		return head
	}

	prevOfmNode := head //记录开头处，该指针指向下标为m的元素的前一个。如果开头是 head（m == 1），那么该值为 nil
	var mNode *ListNode
	//第一步：前往m-1处
	//特殊情况：从头就需要开始翻转。
	//注意：若出现这种情况，翻转过后，新 head应该为 n，而不是现有head。
	//		由于现在不知道 n的位置，故修改 head的操作 将在第三步进行
	if m == 1 { //特殊时需要采取的操作
		prevOfmNode = nil
		mNode = head
	} else { //常规
		for i := 1; i < m-1; i++ {
			prevOfmNode = prevOfmNode.Next
		}
		mNode = prevOfmNode.Next //记录m位的指针，因为翻转过后，st.Next指向nd，同时，nd.Next指向st
	}

	//第二步：开始翻转
	st := mNode   //（开始翻转前），st是 m元素
	nd := st.Next //（开始翻转前），nd是 m+1元素
	rd := nd.Next //（开始翻转前），rd是 m+2元素

	for i := 0; i < n-m-1; i++ {
		//n-m-1解释：
		//当 n==len(list)时，不能继续移动指针了！
		//此时，rd(n+1) == nil，`rd = rd.Next` 会报错。
		//并且，执行完最后一次翻转元素工作后，现有的st、nd、rd指针已可以保证 第三步后处理 工作的进行。无需再次移动这三个指针。
		nd.Next = st //nd指向st，第二位指向第一位										m+1.Next = m
		st = nd      //移动指针，旧第二位  是  新第一位									m = m+1（已知）
		nd = rd      //移动指针，旧第三位  是  新第二位									m+1 = m+2（已知）
		rd = rd.Next //移动指针，旧第三位  是  旧第三位的Next（第三位的Next不允许被修改）	m+2 = m+2.Next
	}
	//由于上述循环，少翻转了一对元素，在这里补上。
	nd.Next = st

	//（翻转过后），st是 n-1元素，后续已无用武之地
	//（翻转过后），nd是 n元素，没有任何一个元素的 Next 指向该元素（n-1.Next指向了n+2元素）
	//（翻转过后），rd是 n+1元素，没有任何一个元素的 Next 指向该元素（n.Next指向了n-1元素）

	//第三步：后处理，需要处理三个指针，在此处不要忘了处理 m == 1 的特殊情况：

	if prevOfmNode == nil { //如果翻转是从头开始的，即 m == 1时，prevOfmNode为 nil。
		//指针三：如果翻转是从头开始的，此时的 head应该是 n元素
		head = nd
	} else {
		//指针一：翻转过后，没有任何一个元素的 Next 指向 n元素。
		//需要将 m-1.Next 修改为 n元素。即：将prevOfmNode.Next指向 nd元素。
		prevOfmNode.Next = nd
	}
	//指针二：翻转过后，没有任何一个元素的 Next指向 n+1元素
	//需要将 m.Next 修改为 n+1元素。即：将 mNode.Next指向rd元素。
	mNode.Next = rd

	//结束
	return head
}
