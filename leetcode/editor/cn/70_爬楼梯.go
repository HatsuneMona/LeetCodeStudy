/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？



 示例 1：


输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶

 示例 2：


输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶




 提示：


 1 <= n <= 45


 Related Topics 记忆化搜索 数学 动态规划 👍 2718 👎 0

*/

package leetcode

func Q70Main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {

	// 1、2层的情况
	// 一层台阶时，有1种上楼方法；二层台阶时，有2种上楼方法
	if n < 3 {
		return n
	}

	// abc 分别对应 n-2, n-1, n
	// 按动态规划的做题思路，abc分别对应 dp[n-2], dp[n-1], dp[n]
	// dp数组的初始化值为：dp[1] = 1 （一层台阶时，有1种上楼方法）
	// dp数组的初始化值为：dp[2] = 2 （二层台阶时，有2种上楼方法）
	a, b, c := 1, 2, 0

	// 遍历计算方向：
	// 从1.2楼开始，往后续楼层计算
	// 1.2楼为已知，若计算第3楼，则需要计算1次；计算第4楼，则需要计算2次。
	for n > 2 {
		// 递推公式
		// dp[n] = dp[n-2] + dp[n-1]
		c = a + b
		a, b = b, c

		n--
	}

	return c

}

// leetcode submit region end(Prohibit modification and deletion)
