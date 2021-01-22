/*
88. 合并两个有序数组

给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。

示例 1：
	输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	输出：[1,2,2,3,5,6]

示例 2：
	输入：nums1 = [1], m = 1, nums2 = [], n = 0
	输出：[1]

提示：
	nums1.length == m + n
	nums2.length == n
	0 <= m, n <= 200
	1 <= m + n <= 200
	-109 <= nums1[i], nums2[i] <= 109
*/

package mmn

import "fmt"

func mmn() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	fmt.Printf("调用函数前a的地址：%p\n", a)
	merge(a, 3, b, len(b))
	fmt.Println(a)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Printf("运算前nums1的地址：%p\n", nums1)
	//特殊情况1：arrayA或arrayB其中至少有一个数组为空。
	if m == 0 { //第一个数组为空，第二个数组有东西。或第一个第二个数组都是空的
		for i, n := range nums2 {
			nums1[i] = n
		}
		return
	}
	if n == 0 { //第二个数组是空的情况下
		return
	}
	m--
	n--
	for i := m + n - 1; i >= 0; i-- { //从后往前遍历
		if n == -1 { //B循环完了，但是A中还有元素
			//本身就是升序的，所以不用管了
			return
		} else if m == -1 { //A循环完了，但是B中还有元素
			nums1[i] = nums2[n]
			n--
			continue
		} else if nums1[m] < nums2[n] {
			nums1[i] = nums2[n]
			n--
		} else if nums1[m] >= nums2[n] {
			nums1[i] = nums1[m]
			m--
		}
	}
	fmt.Printf("运算后nums1的地址：%p\n", nums1)
}
