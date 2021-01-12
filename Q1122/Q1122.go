/**
 * @Author HatsuneMona
 * @Date  2020-11-14 19:17
 * @Description Q1122.数组的相对排序
 **/
/*
给你两个数组，arr1 和 arr2，
arr2 中的元素各不相同
arr2 中的每个元素都出现在 arr1 中
对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。
未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。

示例：
输入：	arr1 = [2,3,1,3,2,4,6,7,9,2,19],
		arr2 = [2,1,4,3,9,6]
输出：	[2,2,2,1,4,3,3,9,6,7,19]

提示：
	arr1.length, arr2.length <= 1000
	0 <= arr1[i], arr2[i] <= 1000
	arr2 中的元素 arr2[i] 各不相同
	arr2 中的每个元素 arr2[i] 都出现在 arr1 中
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/relative-sort-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
package Q1122

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	stat := make(map[int]int, len(arr2))
	notInArr2 := make([]int, 0)
	for _, I2 := range arr2 {
		stat[I2] = 0
	}
	for _, I1 := range arr1 {
		if _, ok := stat[I1]; ok != false {
			stat[I1]++
		} else {
			notInArr2 = append(notInArr2, I1)
		}
	}
	sort.Ints(notInArr2)
	out := make([]int, 0, len(arr1))
	for _, I2 := range arr2 {
		out = append(out, func(Int, num int) []int {
			temp := make([]int, num)
			for i, _ := range temp {
				temp[i] = Int
			}
			return temp
		}(I2, stat[I2])...)
	}
	return append(out, notInArr2...)
}
