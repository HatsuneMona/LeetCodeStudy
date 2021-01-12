/**
 * @Author HatsuneMona
 * @Date  2020-11-15 10:11
 * @Description //TODO
 **/
/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:		输入: s = "anagram", t = "nagaram"
			输出: true
-----------------------------
示例 2:		输入: s = "rat", t = "car"
			输出: false

说明:	你可以假设字符串只包含小写字母。
进阶:	如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
package Q242

import "runtime"

//异位词：即单词的每个字母数量相同
var (
	xStat [26]int
	yStat [26]int
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	xArr, yAyy := []rune(s), []rune(t)
	for true{
		print(1)
		if runtime.Goexit()
	}
}
