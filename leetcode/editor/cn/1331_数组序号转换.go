/**
给你一个整数数组 arr ，请你将数组中的每个元素替换为它们排序后的序号。

 序号代表了一个元素有多大。序号编号的规则如下：


 序号从 1 开始编号。
 一个元素越大，那么序号越大。如果两个元素相等，那么它们的序号相同。
 每个数字的序号都应该尽可能地小。




 示例 1：

 输入：arr = [40,10,20,30]
输出：[4,1,2,3]
解释：40 是最大的元素。 10 是最小的元素。 20 是第二小的数字。 30 是第三小的数字。

 示例 2：

 输入：arr = [100,100,100]
输出：[1,1,1]
解释：所有元素有相同的序号。


 示例 3：

 输入：arr = [37,12,28,9,100,56,80,5,12]
输出：[5,3,4,2,8,6,7,1,3]




 提示：


 0 <= arr.length <= 10⁵
 -10⁹ <= arr[i] <= 10⁹


 Related Topics 数组 哈希表 排序 👍 81 👎 0

*/

package leetcode

import (
	"fmt"
	"sort"
)

// > 2022/07/28 14:20:27
// 解答成功:
//	 执行耗时:64 ms,击败了79.17% 的Go用户
//	 内存消耗:11.8 MB,击败了60.42% 的Go用户
func Q1331Main() {
	test := []int{40, 10, 20, 30}
	fmt.Printf("测试结果：%v", arrayRankTransform(test))
}

//leetcode submit region begin(Prohibit modification and deletion)
func arrayRankTransform(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}

	sortMap := map[int]int{}
	sortArr := make([]int, 0, len(arr))

	for _, i := range arr {
		if _, ok := sortMap[i]; !ok {
			sortMap[i] = 0
			sortArr = append(sortArr, i)
		}
	}

	sort.Ints(sortArr)
	for index, i := range sortArr {
		sortMap[i] = index + 1 // 题目要求从 1 开始，而非从 0 开始
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = sortMap[arr[i]]
	}
	return arr
}

//leetcode submit region end(Prohibit modification and deletion)
