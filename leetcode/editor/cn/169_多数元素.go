//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,2,3]
//输出：3
//
// 示例 2：
//
//
//输入：nums = [2,2,1,1,1,2,2]
//输出：2
//
//
//
//提示：
//
//
// n == nums.length
// 1 <= n <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
//
// Related Topics 数组 哈希表 分治 计数 排序 👍 2486 👎 0

package leetcode

func Q169Main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	// 投票做法，因为众数 的数量大于 1/2 * n 所以
	c := 0
	cnt := 0
	for _, v := range nums {
		if c == 0 || cnt < 0 {
			c = v
			cnt = 0
		} else if c == v {
			cnt++
		} else {
			cnt--
		}
	}
	return c
}

//leetcode submit region end(Prohibit modification and deletion)
/*
func majorityElement(nums []int) int {
	// 普通做法
	stat := make(map[int]int)
	for _, num := range nums {
		if stat[num] == 0 {
			stat[num] = 0
		}
		stat[num]++
	}
	maxKey, maxCount := 0, 0
	for key, count := range stat {
		if count > maxCount {
			maxKey = key
			maxCount = count
		}
	}
	return maxKey
}
*/
