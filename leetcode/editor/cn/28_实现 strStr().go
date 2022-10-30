/**
实现 strStr() 函数。

 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不
存在，则返回 -1 。

 说明：

 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。



 示例 1：


输入：haystack = "hello", needle = "ll"
输出：2


 示例 2：


输入：haystack = "aaaaa", needle = "bba"
输出：-1




 提示：


 1 <= haystack.length, needle.length <= 10⁴
 haystack 和 needle 仅由小写英文字符组成


 Related Topics 双指针 字符串 字符串匹配 👍 1512 👎 0

*/

package leetcode

import "fmt"

func TestStrStr(haystack string, needle string) int {
	return strStr(haystack, needle)
}

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {

	// 题目已知 “haystack 和 needle 仅由小写英文字符组成”
	// 否则需要用 utf8.RuneCountInString()
	if len(needle) == 0 {
		return 0
	} else if len(haystack) < len(needle) {
		return -1
	}

	// 暴力解法
	// > 2022/07/27 08:49:12
	// 解答成功:
	// 	    执行耗时:0 ms,击败了100.00% 的Go用户
	//  	内存消耗:1.8 MB,击败了90.99% 的Go用户
	{

		//// i 主串
		//// j 子串
		//i, j := 0, 0
		//for {
		//
		//	// 主串、子串指针所指的字符相同
		//	if haystack[i] == needle[j] {
		//		j++ // 子串指针向后移一位
		//		i++ // 主串指针向后移一位
		//
		//	} else {
		//		// 匹配失败
		//		i = i - j + 1 // 主串指针，归到子串第一位匹配的字符 的下一个字符上
		//		j = 0         // 子串指针归到0位
		//
		//	}
		//
		//	// EXIT   找到后退出
		//	// 当 j == len(needle) 时，找到匹配值，输出。
		//	if j == len(needle) {
		//		return i - j
		//	}
		//
		//	// EXIT   未找到退出
		//	// 主串指针已遍历完主串
		//	// TODO 为了性能，可以多加一个条件：长度不够。
		//	if i >= len(haystack) {
		//		return -1
		//	}
		//}
		//
		//return -1
	}

	// KMP 算法
	// > 2022/07/27 18:33:28
	// 解答成功:
	// 执行耗时:4 ms,击败了100.00% 的Go用户
	// 内存消耗:2 MB,击败了7.15% 的Go用户
	{
		next := make([]int, len(needle))
		getNext(next, needle)

		// i 主串
		// j 子串
		i, j := 0, 0
		for ; i < len(haystack); i++ {

			fmt.Printf("haystack[%v] = '%c' ; needle[%v] = '%c' ;\n", i, haystack[i], j, needle[j])

			// 主串、子串指针所指的字符不同，回溯 j 指针
			for j > 0 && haystack[i] != needle[j] {
				j = next[j-1] // 查表回溯 j 指针
			}

			if haystack[i] == needle[j] {
				j++
			}

			// EXIT  找到后退出
			if j == len(needle) {
				return i - j + 1
			}
		}
		// EXIT 没找到退出
		return -1
	}
}

// getNext
// 求 KMP 算法所需的 next 数组
func getNext(next []int, needle string) {

	// needle   a a b a a f
	// next     0 1 0 1 2 0

	// 前缀定义  a、aa、aab、aaba、aabaa
	// 后缀定义

	// 初始化
	// j 指向 前缀末尾位置；j 还代表着 i 之前（包括 i ）之前的子串的最长相等前后缀长度。
	// i 指向 后缀末尾位置
	j := 0
	next[0] = 0 // 0 的位置需要回退到 0
	for i := 1; i < len(needle); i++ {

		// 处理前后缀不相同的情况
		// 指针情况     j i
		// needle    a a b a a f
		// next      0 1 x . . .
		for j > 0 && needle[i] != needle[j] {
			// 此时 j 应该向前回退
			// j 回退的位置，看它前一位 next 数组的值，就是 j 要回退的下标
			if j > 0 {
				j = next[j-1]
			}
		}

		// 处理前后缀相同的情况
		// 指针情况   j i
		// needle    a a b a a f
		// next      0 x . . . .
		if needle[i] == needle[j] {
			j++
		}

		// 更新Next数组
		next[i] = j
	}

	fmt.Printf("getNext 数组结果： %v \n", next)

}

//leetcode submit region end(Prohibit modification and deletion)
