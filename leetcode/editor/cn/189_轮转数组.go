// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//
//
//
// 示例 1:
//
//
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4]
//
//
// 示例 2:
//
//
// 输入：nums = [-1,-100,3,99], k = 2
// 输出：[3,99,-1,-100]
// 解释:
// 向右轮转 1 步: [99,-1,-100,3]
// 向右轮转 2 步: [3,99,-1,-100]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
// 0 <= k <= 10⁵
//
//
//
//
// 进阶：
//
//
// 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
//
//
// Related Topics 数组 数学 双指针 👍 2373 👎 0

package leetcode

import "fmt"

func Q189Main() {
	// nums, k := []int{1, 2, 3, 4, 5, 6, 7}, 3
	// nums, k := []int{1, 2, 3, 4}, 2
	nums, k := []int{1}, 2
	rotate(nums, k)
	fmt.Printf("%v", nums)
}

// leetcode submit region begin(Prohibit modification and deletion)
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, num := range nums {
		// i+k 把当前元素向后移动 k个
		// %len(nums) 如果向后移动越界了，则将元素从第一位开始往后放
		newNums[(i+k)%len(nums)] = num
	}
	copy(nums, newNums)
}

// leetcode submit region end(Prohibit modification and deletion)
