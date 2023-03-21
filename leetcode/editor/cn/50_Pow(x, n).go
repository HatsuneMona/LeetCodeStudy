/**
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xⁿ ）。



 示例 1：


输入：x = 2.00000, n = 10
输出：1024.00000


 示例 2：


输入：x = 2.10000, n = 3
输出：9.26100


 示例 3：


输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25




 提示：


 -100.0 < x < 100.0
 -231 <= n <= 231-1
 n 是一个整数
 -104 <= xⁿ <= 104


 Related Topics 递归 数学 👍 1139 👎 0

*/

package leetcode

import (
	"fmt"
	"math"
)

func Q50() {
	// fmt.Println(myPow(2, 10), "预期：", math.Pow(2, 10))
	fmt.Println(myPow(0.00001, 2147483647), "预期：", math.Pow(0.00001, 2147483647))
}

// leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return pow(x, n)
	} else {
		return 1 / pow(x, -n)
	}
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	t := pow(x, n/2)
	if n%2 == 1 {
		// 奇数
		return t * t * x
	} else {
		// 偶数
		return t * t
	}
}

// leetcode submit region end(Prohibit modification and deletion)
