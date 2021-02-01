/*
888. 公平的糖果棒交换
爱丽丝和鲍勃有不同大小的糖果棒：A[i] 是爱丽丝拥有的第 i 根糖果棒的大小，B[j] 是鲍勃拥有的第 j 根糖果棒的大小。
因为他们是朋友，所以他们想交换一根糖果棒，这样交换后，他们都有相同的糖果总量。
（一个人拥有的糖果总量是他们拥有的糖果棒大小的总和。）
返回一个整数数组 ans，其中 ans[0] 是爱丽丝必须交换的糖果棒的大小，ans[1] 是 Bob 必须交换的糖果棒的大小。
如果有多个答案，你可以返回其中任何一个。保证答案存在。

示例 1：							示例 2：
	输入：A = [1,1], B = [2,2]		输入：A = [1,2], B = [2,3]
	输出：[1,2]						输出：[1,2]

示例 3：							示例 4：
	输入：A = [2], B = [1,3]			输入：A = [1,2,5], B = [2,4]
	输出：[2,3]						输出：[5,4]

提示：
	1 <= A.length <= 10000
	1 <= B.length <= 10000
	1 <= A[i] <= 100000
	1 <= B[i] <= 100000
	保证爱丽丝与鲍勃的糖果总量不同。
	答案肯定存在。
*/

package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	A := []int{79, 97, 89, 31, 66}
	B := []int{66, 23, 76, 91, 60}
	fmt.Println(fairCandySwap(A, B))
}

func fairCandySwap(A []int, B []int) []int {
	aTotal := 0
	bTotal := 0
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for _, value := range A {
			aTotal += value
		}
		wg.Done()
	}()
	go func() {
		for _, value := range B {
			bTotal += value
		}
		wg.Done()
	}()
	wg.Wait()
	need := (aTotal - bTotal) / 2 //两个输出，只需要相差need个即可
	//sort.Ints(A)//A数组，不涉及到二分查找，不需要排序
	sort.Ints(B)
	for _, aValue := range A {
		bValue := aValue - need
		if bValue < B[0] || bValue > B[len(B)-1] {
			continue
		}
		if BinarySearch(B, bValue) == true {
			return []int{aValue, bValue}
		}
	}
	return []int{-1, -1}
}

func BinarySearch(arr []int, target int) bool {
	if len(arr) < 3 {
		for _, value := range arr {
			if value == target {
				return true
			}
		}
		return false
	}
	a := 0
	b := len(arr)
	if arr[(a+b)/2] < target {
		return BinarySearch(arr[(a+b)/2:], target)
	} else {
		return BinarySearch(arr[:(a+b)/2+1], target)
	}
}
