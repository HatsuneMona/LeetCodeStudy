/*
628. 三个数的最大乘积
给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

示例 1:					示例 2:
	输入: [1,2,3]			输入: [1,2,3,4]
	输出: 6					输出: 24

注意:
	给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
	输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。
*/
package main

import "fmt"

/*
MyNote：
全正数情况：
	找到三个最大的正数。

全负数情况：
	此时需要找3个绝对值最小的负数，出现0需要忽略掉。

正负数混合情况：
	1.三个最大正数相乘
	2.两个绝对值最大的负数相乘，再乘一个最大正数
	需要找到至少两个绝对值最大的负数，a,b。将其转换为正数，替换掉已知最大的三个正数。
	特殊：如果只有四个数：两个正数，两个负数{3,5,-1,-2}
		必须将两个负数包含（要挑选两个最大的负数），然后再找一个最大的正数。

负数和0混合情况：
	至少有3个0，那么可以直接输出0

特殊情况：
	nums[]的长度小于4（等于3）
	纯0数组
*/

func main() {
	//fmt.Println(maximumProduct([]int{1,2,3}))
	//fmt.Println(maximumProduct([]int{1,2,3,4}))
	fmt.Println(maximumProduct([]int{1, 2, -1, -2}))
}

func maximumProduct(nums []int) int {
	if len(nums) < 4 { //处理特殊情况
		return nums[0] * nums[1] * nums[2]
	}
	numplus := []int{0 /*远离1000*/, 0, 0 /*靠近1000*/} //存储最大的正数
	//TODO 用不着三个绝对值最大的负数，只需要两个即可
	numsubmax := []int{0 /*远离-1000*/, 0, 0 /*靠近-1000*/}     //存储绝对值最大的负数
	numsubmin := []int{-1001 /*远离0*/, -1001, -1001 /*靠近0*/} //存储绝对值最小的负数
	zeroCount:=0
	for _, n := range nums {
		if n == 0 { //对数组中所有的0进行忽略
			zeroCount++
			continue
		}
		if n > 0 { //数组中的正数
			if n > numplus[0] { //大于已知最大的三个数中最小的一个数
				numplus[0] = n
				for j := 1; j < len(numplus); j++ { //冒泡法排序
					if numplus[j] < numplus[j-1] {
						numplus[j], numplus[j-1] = numplus[j-1], numplus[j]
					}
				}
				if numplus[0] == 1000 { //如果最大的三个整数都是1000了（题目中指明了每一个元素的范围）
					return 1000 * 1000 * 1000 //可以直接返回最大值
				}
			}
		}
		if n < 0 { //数组中的负数
			if n > numsubmin[0] { //比离0最远的数，更靠近0
				numsubmin[0] = n
				for j := 1; j < len(numsubmin); j++ { //冒个泡
					if numsubmin[j] < numsubmin[j-1] {
						numsubmin[j], numsubmin[j-1] = numsubmin[j-1], numsubmin[j]
					}
				}
				//这里不能直接返回最大值，不能确定后续元素中是否有正数
			}
			if n < numsubmax[0] { //比远离-1000的数，更靠近-1000
				numsubmax[0] = n
				//TODO 用不着三个绝对值最大的负数，只需要两个即可
				for j := 1; j < len(numsubmax); j++ { //冒个泡
					if numsubmax[j] > numsubmax[j-1] {
						numsubmax[j], numsubmax[j-1] = numsubmax[j-1], numsubmax[j]
					}
				}
				//不能直接返回最大值，三个 -1000 不能构成最大值（但能构成最小值）
			}
		}
	} //for 循环所有元素 结束
	//	numplus := []int{/*远离1000*/, 0, /*靠近1000*/}               //存储最大的正数
	//	numsubmax := []int{/*远离-1000*/, -1001, /*靠近-1000*/} 		//存储绝对值最大的负数
	//	numsubmin := []int{/*远离0*/, 0, /*靠近0*/}                   //存储绝对值最小的负数

	if numplus[2] == 0 { //如果最靠近1000的数是0，那么可以认为数组中没有正数
		//检查是否有足够（至少三个）0
		if zeroCount>2{
			return 0
		}
		//此时，挑三个最小的负数，相乘即为答案
		return numsubmin[0] * numsubmin[1] * numsubmin[2]
	}

	if numsubmin[2] == 0 && numsubmin[0] == 0 { //如果存储绝对值最小的负数中，没有元素，那么可以认为数组中没有负数
		//此时，挑三个最大的正数，相乘即为答案
		return numplus[0] * numplus[1] * numplus[2]
	}

	//既有正数，又有负数的情况
	//1.三个最大正数相乘
	a := numplus[0] * numplus[1] * numplus[2]
	//2.两个绝对值最大的负数相乘，再乘一个最大正数
	//TODO 用不着三个绝对值最大的负数，只需要两个即可
	b := numplus[2] * numsubmax[2] * numsubmax[1]
	if a > b {
		return a
	}
	return b
}
