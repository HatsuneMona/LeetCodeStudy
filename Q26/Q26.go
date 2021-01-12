/*
26. 删除排序数组中的重复项
	给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
	不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
示例 1:
		给定数组 nums = [1,1,2],
		函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
		你不需要考虑数组中超出新长度后面的元素。
示例 2:
		给定 nums = [0,0,1,1,1,2,2,3,3,4],
		函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
		你不需要考虑数组中超出新长度后面的元素。
说明:
		为什么返回数值是整数，但输出的答案是数组呢?
		请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
		你可以想象内部操作如下:
			// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
			int len = removeDuplicates(nums);
			// 在函数里修改输入数组对于调用者是可见的。
			// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
			for (int i = 0; i < len; i++) {
				print(nums[i]);
			}
*/
package main

func main() {
	var test = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	l := removeDuplicates(test)
	for i := 0; i < l; i++ {
		print(test[i], " , ")

	}
}

func removeDuplicates(nums []int) int {
	count := 0 //用于统计数量
	//对于 1.空的数组；2.只有一个元素的数组，直接返回该数组的长度
	if len(nums) < 2 {
		return len(nums)
	}

	for _, n := range nums {
		//注意：题目中已经说明，这是一个排序的数组
		if n != nums[count] {	//count 	指向末尾的位置
			count++				//count++ 	指向需要修改的位置。
			nums[count] = n
		}

		/*
			//这个是未排序的做法
			flog := true //是否是仅出现一次的元素
			for j := 0; j < count; j++ {
				if nums[j] == n {
					flog = false
					break
				}
			}
			if flog {
				nums[count] = n
				count++
			}
		*/
	}
	count++
	return count
}
