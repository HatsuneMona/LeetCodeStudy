/*
387. 字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

示例：
		s = "leetcode"			s = "loveleetcode"
		返回 0					返回 2

提示：你可以假定该字符串只包含小写字母。
*/
package main

func main() {
	str := "loveleetcode"
	println("结果为：", firstUniqChar(str))
}

func firstUniqChar(s string) int {
	//当字符串的长度为0/1时，为特殊情况
	if len(s) < 2 {
		if len(s) == 0 {
			return -1
		}
		return 0
	}
	//bs := []byte(s)          //因为只有小写字母，所以直接将rune数组转成byte数组
	haveSeen := [26][2]int{} //统计字符出现次数，以及字符第一次出现的位置

	for i, b := range s {
		ch := b - 97 //获取当前字母编号

		if haveSeen[ch][0] > 1 { //如果已经见过很多次了
			continue //忽略
		} else if haveSeen[ch][0] == 0 { //如果没见过
			haveSeen[ch][0]++   //标记见过一次了
			haveSeen[ch][1] = i //标记第一次见的位置
		} else if haveSeen[ch][0] == 1 { //如果已经见过了一次
			haveSeen[ch][0]++ //标记字符作废
		}
	}
	//检查统计结果
	min := -1 //记录唯一出现过的第一个字符的位置。
	for _, stat := range haveSeen {
		if stat[0] == 1 { //如果该字符只出现过一次
			if stat[1] < min || min < 0 { //如果该记录出现的位置，小于已经记录的最早位置
				min = stat[1] //刷新整体最早的位置
			}
		}
	}
	return min

	//for i := len(bs) - 1; i >= 0; i-- { //必须从后往前数
	//	ch := int(bs[i]) - 61      //获取当前字母编号
	//	if haveSeen[ch] == false { //如果之前没见过这个字符
	//		seenCount++         //统计+1
	//		haveSeen[ch] = true //标记这个字母已经见过了
	//		if seenCount == 26 {
	//			return i
	//		} //如果26个字母都见过了，而且这是最后一个字母，则直接
	//	}
	//}
}
