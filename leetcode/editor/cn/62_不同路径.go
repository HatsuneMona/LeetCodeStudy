/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

 问总共有多少条不同的路径？



 示例 1：


输入：m = 3, n = 7
输出：28

 示例 2：


输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下


 示例 3：


输入：m = 7, n = 3
输出：28


 示例 4：


输入：m = 3, n = 3
输出：6



 提示：


 1 <= m, n <= 100
 题目数据保证答案小于等于 2 * 10⁹


 Related Topics 数学 动态规划 组合数学 👍 1594 👎 0

*/

package leetcode

import (
	"fmt"
)

func Q62Main() {
	m, n := 3, 7
	fmt.Println("输出结果：", uniquePaths(m, n))
}

// leetcode submit region begin(Prohibit modification and deletion)
// m 行     n 列
func uniquePaths(m int, n int) int {
	// 特殊情况
	if m == 1 || n == 1 {
		return 1
	}

	// 动态规划解题第一步：定义dp[i]
	// x, y ：坐标
	// dp[x][y] = 到达(x,y)，共有几种走法
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 动态规划解题第二步：确定递推公式
	// dp[x][y] = (isset(dp[x-1][y]) ? dp[x-1][y] : 0) + (isset(dp[x][y-1]) ? dp[x][y-1] : 0)
	// 解释：若要去往(x,y)，必须从(x-1,y)或(x,y-1)前往(x,y)，且只有一种方法。

	// 动态规划解题第三步：初始化dp数组
	// dp = | 1 ? ? ? |  or  | 1 1 1 1 |
	//      | ? ? ? ? |      | 1 ? ? ? |
	//      | ? ? ? ? |      | 1 ? ? ? |
	// 只有一种方法，去这些位置

	// 动态规划解题第四步：遍历顺序
	// 从左往右，从上到下
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if x == 0 || y == 0 {
				dp[x][y] = 1
				// fmt.Printf("1\t")
			} else {
				dp[x][y] = dp[x-1][y] + dp[x][y-1]
				// fmt.Print(dp[x-1][y]+dp[x][y-1], "\t")
			}

		}
		// fmt.Println("")
	}

	return dp[m-1][n-1]
}

// leetcode submit region end(Prohibit modification and deletion)
