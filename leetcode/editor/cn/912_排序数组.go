/**
给你一个整数数组 nums，请你将该数组升序排列。






 示例 1：


输入：nums = [5,2,3,1]
输出：[1,2,3,5]


 示例 2：


输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]




 提示：


 1 <= nums.length <= 5 * 10⁴
 -5 * 10⁴ <= nums[i] <= 5 * 10⁴


 Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 628 👎 0

*/

package leetcode

import (
	"fmt"
	"math/rand"
	"time"
)

func Q912Main() {
	unSortArr := []int{5, 2, 3, 1}
	//unSortArr := []int{1, -3}
	//unSortArr := []int{-3, 1}
	//unSortArr := make([]int, 0, 50000)
	str := fmt.Sprintf("原始数组：%v \n", unSortArr)
	//for i := 1; i < 50001; i++ {
	//	unSortArr = append(unSortArr, i)
	//}

	sT := time.Now()
	result := sortArray(unSortArr)
	cost := time.Since(sT)
	fmt.Printf("%s排序后数组：%v\n耗时：%v", str, result, cost)
}

//leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	fmt.Printf("input Arr: %v\n", nums)

	if len(nums) < 2 { // 递归终止条件（入参中只有一个或零个元素）
		return nums
	}

	// 初始化指针
	// 任意选一个 基准值
	rand.Seed(int64(time.Now().Nanosecond()))
	kp := rand.Intn(len(nums))
	key := nums[kp]
	nums[0], nums[kp] = nums[kp], nums[0]

	// i 在数组开头，如果发现有比基准值小的元素，该元素将被放置在位置 i 上。
	// j 在数组结尾，为待与基准值比较的元素。
	i, j := 0, len(nums)-1

	// 开始替换位置
	for ; j > i; j-- {
		// 比基准值小的 元素 ，准备放到基准值左面
		for nums[j] < key { // 这里不能是 if ，因为既需要判断 j 位置，也需要把交换过来的 i 的值判断了
			nums[i] = nums[j]
			i++
			nums[j] = nums[i]
		}
	}

	// j == i 时，for 循环结束
	// 把 基准值 补回到数组中
	nums[i] = key
	fmt.Printf("exchange end. output: %v\n", nums)

	// 二分递归执行剩下的部分
	result := make([]int, i, len(nums))
	// 左半部分
	fmt.Printf("ecursion left.")
	result = sortArray(nums[:i]) // golang的切片是左闭右开的
	// 处理基准值，基准值不参与到下次递归中。
	result = append(result, key)
	// 右半部分
	fmt.Printf("ecursion right.")
	result = append(result, sortArray(nums[i+1:])...)

	fmt.Printf("ecursion end. output: %v\n", result)

	return result

}

//leetcode submit region end(Prohibit modification and deletion)
