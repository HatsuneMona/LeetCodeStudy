/**
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返
回 -1。

 示例 1:

 输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4


 示例 2:

 输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1




 提示：


 你可以假设 nums 中的所有元素是不重复的。
 n 将在 [1, 10000]之间。
 nums 的每个元素都将在 [-9999, 9999]之间。


 Related Topics 数组 二分查找 👍 1025 👎 0

*/

package leetcode

import (
    "fmt"
)

func Q704Main() {
    nums := []int{-1, 0, 3, 5, 9, 12}

    for i, target := range nums {
        fmt.Printf("预期为：%v   测试结果为：%v \n", i, search(nums, target))
    }

}

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
    if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
        return -1
    }

    if len(nums) == 1 && nums[0] == target {
        return 0
    }

    // fir, mid, las := 0, len(nums)/2, len(nums)-1
    mid := len(nums) / 2

    if target < nums[mid] {
        res := search(nums[:mid], target)
        // fmt.Printf("搜索值 %v ，搜索结果%v，搜索范围 %v \n", target, res, nums[:mid])
        return res
    } else if target > nums[mid] {

        res := search(nums[mid:], target)
        // fmt.Printf("搜索值 %v ，搜索结果%v，搜索范围 %v \n", target, res, nums[mid:])
        if res == -1 {
            return -1
        }
        return mid + res
    } else { // target == nums[mid]
        // fmt.Printf("搜索值 %v ，搜索结果%v，搜索范围 %v \n", target, mid, nums)
        return mid
    }
}

// leetcode submit region end(Prohibit modification and deletion)
