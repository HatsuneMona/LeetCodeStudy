/**
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。



 示例 1:


输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。


 示例 2:


输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。




 提示:


 1 <= s.length, p.length <= 3 * 10⁴
 s 和 p 仅包含小写字母


 Related Topics 哈希表 字符串 滑动窗口 👍 1038 👎 0

*/

package leetcode

import (
	"fmt"
)

func Q438Main() {
	s, p := "cbaebabacd", "abc"
	fmt.Println("输出结果：", findAnagrams(s, p))
}

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	pMap := make([]int, 26)

	for _, b := range []byte(p) {
		pMap[b-'a']++
	}

	result := []int{}

	ps, pe := 0, len(p)-1
	ss := []byte(s)
	for ; pe < len(ss); pe = ps + len(p) - 1 {

		// pe
		if pMap[ss[pe]-'a'] == 0 {
			ps = pe + 1
			continue // 隐式：pe++
		}

		// if pe-ps == len(p)-1 {
		checkMap := make([]int, 26)

		check := true
		for pCheck := ps; pCheck <= pe; pCheck++ {
			checkMap[ss[pCheck]-'a']++
			if checkMap[ss[pCheck]-'a'] > pMap[ss[pCheck]-'a'] {
				check = false
				break
			}
		}

		if check {
			result = append(result, ps)
		}

		ps++
		// }
	}

	return result
}

// leetcode submit region end(Prohibit modification and deletion)
