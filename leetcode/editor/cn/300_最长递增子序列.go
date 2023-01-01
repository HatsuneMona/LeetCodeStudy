/**
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列
。

 示例 1：


输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。


 示例 2：


输入：nums = [0,1,0,3,2,3]
输出：4


 示例 3：


输入：nums = [7,7,7,7,7,7,7]
输出：1




 提示：


 1 <= nums.length <= 2500
 -10⁴ <= nums[i] <= 10⁴




 进阶：


 你能将算法的时间复杂度降低到 O(n log(n)) 吗?


 Related Topics 数组 二分查找 动态规划 👍 2930 👎 0

*/

package leetcode

import (
	"fmt"
)

func Q300Main() {
	fmt.Printf("结果为：%v\n", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	// 第一步：
	// dp[i] == 以num[i]结尾的，最长递增子序列的长度
	dp := make([]int, len(nums))

	// 第二步：
	// 递推公式 dp[i] = max(dp[j]+1, dp[i])    j 属于 [0, i)

	// 第三步：
	// 初始化：在任意位置，都有[]int{nums[i]（这个数本身）} 是最长递增子序列
	dp[0] = 1

	// 第四步：
	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				if dp[i] > result {
					result = dp[i]
				}
			}
		}
	}
	// fmt.Printf("dp array: %v\n", dp)
	return result
}

// leetcode submit region end(Prohibit modification and deletion)
