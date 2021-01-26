/*
1128. 等价多米诺骨牌对的数量
给你一个由一些多米诺骨牌组成的列表 dominoes。
如果其中某一张多米诺骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。
形式上，dominoes[i] = [a, b] 和 dominoes[j] = [c, d] 等价的前提是 a==c 且 b==d，或是 a==d 且 b==c。
在 0 <= i < j < dominoes.length 的前提下，找出满足 dominoes[i] 和 dominoes[j] 等价的骨牌对 (i, j) 的数量。

示例：
	输入：dominoes = [[1,2],[2,1],[3,4],[5,6]]
	输出：1

提示：
	1 <= dominoes.length <= 40000
	1 <= dominoes[i][j] <= 9
*/
package main

import "fmt"

func main() {
	dominoes := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
	fmt.Println(numEquivDominoPairs(dominoes))
}

/*
My Note:
	特殊情况：
		只有一个牌
	勘误：
		不能用99乘法表解决这个问题：
		2 * 6 == 3 * 4 == 12
	难点：
		已有三个骨牌相同 a b c，此时又得到了一个骨牌 d。
		在三个骨牌的时候，等价的骨牌对数量为：C^2_3 = 3
		在四个骨牌的时候，等价的骨牌对数量为：C^2_4 = 6 == C^2_3 + 3
	建议修改的地方
		//TODO orderStat中有1/2的内存空间被浪费（sm>la的区域）
*/
func numEquivDominoPairs(dominoes [][]int) int {
	//特殊情况：只有一张牌
	if len(dominoes) == 1 {
		return 0
	}
	result := 0 //存储结果
	//TODO orderStat中有1/2的内存空间被浪费（sm>la的区域）
	orderStat := [9][9]int{} //将多米诺骨牌“摆整齐”上面小，下面大
	//遍历骨牌
	for _, dominoe := range dominoes {
		sm := dominoe[0] - 1
		la := dominoe[1] - 1
		if sm > la {
			sm, la = la, sm
		}
		result += orderStat[sm][la]
		orderStat[sm][la]++
	}
	return result
}
