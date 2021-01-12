package main

import (
	"fmt"
	"strconv"
)

func main() {
	test := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(test))
}

func summaryRanges(nums []int) []string {
	var s []string //存储结果
	// 排除特殊输入情况，如果输入的数组为空
	if len(nums) < 2 {
		//如果输入的数组只有一个数字
		if len(nums) == 1 {
			s = append(s, strconv.Itoa(nums[0]))
		}
		return s
	}

	start := nums[0]
	length := 0
	for i := 1; i < len(nums); i++ {
		if start+length+1 == nums[i] {
			length++
		} else {
			//输出
			s = append(s, out(start, length))
			//清零，准备下次循环
			start = nums[i]
			length = 0
		}
	}
	s = append(s, out(start, length))
	return s
}

func out(s, l int) string {
	if l == 0 {
		return strconv.Itoa(s)
	} else {
		return fmt.Sprint(s, "->", s+l)
	}
}
