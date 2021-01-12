/**
 * @Author HatsuneMona
 * @Date  2020-11-19 18:58
 * @Description //TODO
 **/

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
示例:
		输入: [0,1,0,3,12]
		输出: [1,3,12,0,0]
说明:
		必须在原数组上操作，不能拷贝额外的数组。
		尽量减少操作次数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/move-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
package Q283

func moveZeroes(nums []int) {
	zeroSize := 0
	if len(nums) == 0 {
		return
	}
	for p := 0; p < len(nums)-zeroSize; p++ {
		if nums[p+zeroSize] == 0 {
			zeroSize++
			for i := 0; i < zeroSize; i++ {
				nums[p+i], nums[p+zeroSize+i] = nums[p+zeroSize+i], nums[p+i]
			}
		}
	}
}
