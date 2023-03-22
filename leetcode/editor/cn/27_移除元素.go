/**
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。



 说明:

 为什么返回数值是整数，但输出的答案是数组呢?

 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

 你可以想象内部操作如下:


// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
int len = removeElement(nums, val);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}




 示例 1：


输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums
= [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。


 示例 2：


输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3]
解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元
素。




 提示：


 0 <= nums.length <= 100
 0 <= nums[i] <= 50
 0 <= val <= 100


 Related Topics 数组 双指针 👍 1720 👎 0

*/

package leetcode

import "fmt"

func Q27() {
	tCases := [][]int{
		{1, 2, 3, 0, 5, 6},
		{0},
		{0, 0},
		{},
		{1, 2, 3, 0, 0, 0},
		{1, 2, 3, 0, 0, 0, 7, 0, 0},
		{0, 0, 0, 2},
		{0, 2, 3, 0, 5, 6, 0},
	}
	for _, tCase := range tCases {
		var t []int = make([]int, len(tCase))
		copy(t, tCase)
		res := removeElement(t, 0)
		fmt.Printf("测试用例：%v    结果：%v    返回值：%v %v\n", tCase, t[:res], res, t)
	}

}

// leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	endP := 0 // 最终得到的 结果 的末尾指针
	for p := 0; p < len(nums); p++ {
		if nums[p] != val {
			// endP 肯定小于等于 p，所以可以直接覆盖
			nums[endP] = nums[p]
			endP++
		}
	}
	return endP // endP 的值 同时也代表了 “进行了几次复制操作”
}

// leetcode submit region end(Prohibit modification and deletion)

// 交换的方法
// 解答成功:
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2 MB,击败了22.17% 的Go用户
func removeElementSolution1(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	// endP：准备移动的元素
	// endP 不能为val
	endP := len(nums) - 1

	// 从前往后找 val，找到后用 endP 进行替换
	// 终止条件：p 越过了 endP
	for p := 0; p < len(nums) && p < endP; p++ {
		// 如果替换的值 还是 val，那么继续替换
		// 如果 endP 指到了 p ，那么说明没有可以移除的元素了
		for nums[p] == val && p <= endP {
			nums[p] = nums[endP]
			// nums[endP] = -1 // DEBUG
			endP--
		}
	}

	if endP >= 0 && nums[endP] == val {
		return endP
	} else {
		return endP + 1
	}
}
