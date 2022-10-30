/**
给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。

 返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。



 示例 1：


输入：s = "a1b2"
输出：["a1b2", "a1B2", "A1b2", "A1B2"]


 示例 2:


输入: s = "3z4"
输出: ["3z4","3Z4"]




 提示:


 1 <= s.length <= 12
 s 由小写英文字母、大写英文字母和数字组成


 Related Topics 位运算 字符串 回溯 👍 468 👎 0

*/

package leetcode

import (
	"fmt"
)

func Q784Main() {
	testStr := "abDC"
	result := letterCasePermutation(testStr)

	for _, s := range result {
		fmt.Println("      ", s)
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func letterCasePermutation(s string) []string {
	// 记录一共有多少个字母，以及字母的位置在哪里。
	letterPlaceList := make([]int, 0, len(s))
	// root 是一个字符串，只包含 数字与小写字母
	root := make([]byte, len(s))
	// result 是结果
	result := make([]string, 0, 1<<len(letterPlaceList))

	// 统计字母数量，并初始化root字符串
	for i, b := range []byte(s) {
		if 64 < b && b < 91 {
			root[i] = b + 32
			letterPlaceList = append(letterPlaceList, i)
		} else if 96 < b && b < 123 {
			root[i] = b
			letterPlaceList = append(letterPlaceList, i)
		} else {
			root[i] = b
		}
	}

	// 特殊情况
	if len(letterPlaceList) == 0 {
		return []string{s}
	}

	result = append(result, string(root))

	for len(letterPlaceList) != 0 {
		pl := letterPlaceList[0]

		newStrings := make([]string, len(result))
		for i := 0; i < len(result); i++ {
			newString := []byte(result[i])
			newString[pl] -= 32 // 转小写
			newStrings[i] = string(newString)
		}

		result = append(result, newStrings...)

		letterPlaceList = letterPlaceList[1:]
	}

	return result
}

// leetcode submit region end(Prohibit modification and deletion)
