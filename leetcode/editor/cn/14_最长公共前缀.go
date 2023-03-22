/**
编写一个函数来查找字符串数组中的最长公共前缀。

 如果不存在公共前缀，返回空字符串 ""。



 示例 1：


输入：strs = ["flower","flow","flight"]
输出："fl"


 示例 2：


输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。



 提示：


 1 <= strs.length <= 200
 0 <= strs[i].length <= 200
 strs[i] 仅由小写英文字母组成


 Related Topics 字典树 字符串 👍 2694 👎 0

*/

package leetcode

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) < 2 { // 1 <= strs.length
		return strs[0]
	}
	commonStr := strs[0]

	for _, str := range strs {
		for i, _ := range commonStr {
			// str 与 公共前缀串，其中的一个字符对不上，需要裁剪公共前缀串
			// 如果 str 的长度 小于 公共前缀串，也需要对公共前缀串进行裁剪
			if i == len(str) || str[i] != commonStr[i] {
				commonStr = commonStr[:i]
				break
			}
		}
		// 不存在 公共前缀串
		if len(commonStr) == 0 {
			break
		}
	}
	return commonStr
}

// leetcode submit region end(Prohibit modification and deletion)
