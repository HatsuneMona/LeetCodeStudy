/**
 * @Author HatsuneMona
 * @Date  2021-02-02 11:47
 * @Description 判断两个数组是否相等
 **/
package tools

import (
	"fmt"
)

func IsArrEqual(arrA []int, arrB []int) {
	if len(arrA) != len(arrB) {
		fmt.Println("两个数组长度不同。")
		fmt.Printf("A的长度为%v，B的长度为%v。\n", len(arrA), len(arrB))
		return
	}
	notEqualList := make([]int, len(arrA))
	stat := 0
	for i, value := range arrA {
		if value != arrB[i] {
			notEqualList[stat] = i
			stat++
		}
	}
	if stat == 0 {
		fmt.Println("验证成功！")
		return
	} else {
		fmt.Printf("两个数组共有%v处不同，差异如下：\n", stat)
		for _, value := range notEqualList[:stat] {
			fmt.Printf("差异位置：%v，值：[A[i], B[i]] == [%v, %v] \n", value, arrA[value], arrB[value])
		}
	}
}
