/**
给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

 请你计算并返回达到楼梯顶部的最低花费。



 示例 1：


输入：cost = [10,15,20]
输出：15
解释：你将从下标为 1 的台阶开始。
- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
总花费为 15 。


 示例 2：


输入：cost = [1,100,1,1,1,100,1,1,100,1]
输出：6
解释：你将从下标为 0 的台阶开始。
- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
总花费为 6 。




 提示：


 2 <= cost.length <= 1000
 0 <= cost[i] <= 999


 Related Topics 数组 动态规划 👍 1041 👎 0

*/

package leetcode

import (
	"fmt"
)

func Q746Main() {
	testCase := []int{10, 15, 20}
	fmt.Println("结果输出：", minCostClimbingStairs(testCase))
}

// leetcode submit region begin(Prohibit modification and deletion)
func minCostClimbingStairs(cost []int) int {
	// 楼梯终点（楼梯从0开始计数）
	stairLen := len(cost) + 1

	// 总花费
	// sumCost := 0

	// 动态规划解题第一步：定义dp[i]
	// i ：到达第 i 层    // i 从 0 开始计数
	// dp[i] = 到达 i 层时，所需的最低花费
	dp := make([]int, stairLen)
	// 因为递推时，只需要考虑前两节楼梯的minCost，所以 本题应该有内存O(1)级做法，类似《Q70_爬楼梯》

	// 动态规划解题第二步：确定递推公式
	// dp[i] = min(dp[i-2] + cost[i-2], dp[i-1] + cost[i-1])
	// 解释：到达第 i 层时，因为只能 从i-2层上两节台阶 或 从i-1层上一节台阶
	//      所以需要从 这两种方案里，找出去i层成本最低的方案

	// 动态规划解题第三步：初始化dp数组
	// 去 下标为0 或 下标为1 的台阶的成本为 0
	// dp = []int{0, 0}     // 初始化数组时，已经是 0 了

	// 动态规划解题第四步：遍历顺序
	// 从前往后
	// 从第2节台阶开始爬楼梯，前两节免费，不用爬。爬到 3 才是终点。         0  1  2（第3节楼梯）  3（终点）
	for i := 2; i < stairLen; i++ {
		// 预计花费
		aCost, bCost := dp[i-2]+cost[i-2], dp[i-1]+cost[i-1]

		if aCost < bCost {
			dp[i] = aCost
		} else {
			dp[i] = bCost
		}
		// fmt.Printf("第 i=%v 层递推，最小花费为 %v ，dp 【%v】\n", i, dp[i], dp[:i+1])
	}

	return dp[stairLen-1]
}

// leetcode submit region end(Prohibit modification and deletion)
