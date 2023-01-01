/**
给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。

 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子
序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。



 示例 1：


输入：nums = [1,3,5,4,7]
输出：3
解释：最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。


 示例 2：


输入：nums = [2,2,2,2,2]
输出：1
解释：最长连续递增序列是 [2], 长度为1。




 提示：


 1 <= nums.length <= 10⁴
 -10⁹ <= nums[i] <= 10⁹


 Related Topics 数组 👍 345 👎 0

*/

package leetcode

import (
	"fmt"
)

func Q674Main() {
	// fmt.Printf("结果为：%v\n", findLengthOfLCIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Printf("结果为：%v\n", findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func findLengthOfLCIS(nums []int) int {
	// 第一步：
	// dp[i] == 以num[i]结尾的，最长递增子序列的长度
	dp := make([]int, len(nums))

	// 第二步：
	// 递推公式 dp[i] = max(dp[i-1]+1, dp[i])

	// 第三步：
	// 初始化：在任意位置，都有[]int{nums[i]（这个数本身）} 是最长递增子序列
	dp[0] = 1

	// 第四步：
	result := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1 // 初始化

		if nums[i] > nums[i-1] && dp[i] < dp[i-1]+1 {
			dp[i] = dp[i-1] + 1
			if dp[i] > result {
				result = dp[i]
			}
		}
	}
	// fmt.Printf("dp array: %v\n", dp)
	return result
}

// leetcode submit region end(Prohibit modification and deletion)
