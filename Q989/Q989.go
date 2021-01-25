/*
989. 数组形式的整数加法
	对于非负整数 X 而言，X 的数组形式是每位数字按从左到右的顺序形成的数组。
	例如，如果 X = 1231，那么其数组形式为 [1,2,3,1]。
	给定非负整数 X 的数组形式 A，返回整数 X+K 的数组形式。

示例 1：								示例 2：
	输入：A = [1,2,0,0], K = 34			输入：A = [2,7,4], K = 181
	输出：[1,2,3,4]						输出：[4,5,5]
	解释：1200 + 34 = 1234				解释：274 + 181 = 455
示例 3：								示例 4：
	输入：A = [2,1,5], K = 806			输入：A = [9,9,9,9,9,9,9,9,9,9], K = 1
	输出：[1,0,2,1]						输出：[1,0,0,0,0,0,0,0,0,0,0]
	解释：215 + 806 = 1021				解释：9999999999 + 1 = 10000000000

提示：
	1 <= A.length <= 10000
	0 <= A[i] <= 9
	0 <= K <= 10000
	如果 A.length > 1，那么 A[0] != 0
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(addToArrayForm([]int{9, 9, 9}, 1))
	fmt.Println(addToArrayForm([]int{2, 1, 5}, 806))
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
	fmt.Println(addToArrayForm([]int{2, 7, 4}, 181))
	fmt.Println(addToArrayForm([]int{1, 2}, 3456))
}

/*
My Note:
特殊情况：
	A.length == 1 并且 A[0] = 0（数A的值是0）
		直接返回 K 数组打散后的结果
	K == 0
		直接返回A
	(K + A).length > A.length
		在处理 A[0] 的进位时，会发生越界。
		此时应先扩容A数组，再处理进位。
*/
func addToArrayForm(A []int, K int) []int {
	//K的数值是0
	if K == 0 {
		return A
	}
	arrTemp := [5]int{}
	Klen := 0
	for i := 0; i < 5; i++ {
		//arr[0]低位
		//arr[4]高位
		arrTemp[i] = K / func() int {
			pow := 1
			for x := 0; x < i; x++ {
				pow *= 10
			}
			return pow
		}() % 10
		if arrTemp[i] != 0 {
			Klen = i + 1
		}
	}
	Karr := arrTemp[0:Klen]
	//A的数值是0
	if len(A) == 1 && A[0] == 0 {
		//return Karr //这里错了，Karr是倒序！
		//翻转Karr
		for i := 0; i < Klen/2; i++ {
			Karr[i], Karr[Klen-i-1] = Karr[Klen-i-1], Karr[i]
		}
		return Karr
	}
	//A比K短，需要延长A
	if Klen > len(A) {
		A = append(make([]int, Klen-len(A)), A...)
	}
	//处理相加部分
	for i := 0; i < Klen; i++ {
		A[len(A)-1-i] += Karr[i]
	}
	//处理进位，不包括第一位（A[0]，最高位）
	for i := len(A) - 1; i > 0; i-- {
		if A[i] > 9 {
			A[i] -= 10
			A[i-1]++
		} else if i < len(A)-Klen { //A[i]是个位数，且这里已经不是K能相加的区域了
			//因为加法的进位不会“跨越”，即在个位上相加，如果十位没有进位，则绝对不会在百位上发生进位
			return A
		}
	}
	//单独把A[0]拉出来，处理进位
	if A[0] > 9 {
		A[0] -= 10
		A = append([]int{1}, A...)
	}
	return A
}
