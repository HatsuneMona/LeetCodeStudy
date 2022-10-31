/**
神奇字符串 s 仅由 '1' 和 '2' 组成，并需要遵守下面的规则：


 神奇字符串 s 的神奇之处在于，串联字符串中 '1' 和 '2' 的连续出现次数可以生成该字符串。


 s 的前几个元素是 s = "1221121221221121122……" 。如果将 s 中连续的若干 1 和 2 进行分组，可以得到 "1 22 11 2
1 22 1 22 11 2 11 22 ......" 。每组中 1 或者 2 的出现次数分别是   1 2 1 2 2 1 2 2 .....
." 。上面的出现次数正是 s 自身。

 给你一个整数 n ，返回在神奇字符串 s 的前 n 个数字中 1 的数目。



 示例 1：


输入：n = 6
输出：3
解释：神奇字符串 s 的前 6 个元素是 “122112”，它包含三个 1，因此返回 3 。


 示例 2：


输入：n = 1
输出：1




 提示：


 1 <= n <= 10⁵


 Related Topics 双指针 字符串 👍 118 👎 0

*/

package leetcode

import (
	"fmt"
)

func Q481Main() {
	fmt.Println("输出结果：", magicalString(9))
	fmt.Sprintf("")
}

// leetcode submit region begin(Prohibit modification and deletion)
func magicalString(n int) int {
	realString := make([]int8, n)
	// realString[0] = 2 // 指定第 -1 个字符 2

	numOf1 := 0

	// 根据 p 指针的值，在 q 位置生成新字符串
	p, q := 0, 0
	for q < n {
		// b ：q这个位置应该填什么  1 或是 2
		b := int8(1)
		if q-1 >= 0 && realString[q-1] == 1 {
			b = 2
		}

		realString[q] = b
		// fmt.Println("生成的字符串：", realString[:q+1])
		if b == 1 {
			numOf1++
		}
		if realString[p] == 2 && q+1 < n {
			q++
			realString[q] = b
			// fmt.Println("生成的字符串：", realString[:q+1])
			if b == 1 {
				numOf1++
			}
		}
		// 计数
		p++
		q++
	}

	return numOf1
}

// leetcode submit region end(Prohibit modification and deletion)
