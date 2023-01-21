/**
如果一个密码满足以下所有条件，我们称它是一个 强 密码：


 它有至少 8 个字符。
 至少包含 一个小写英文 字母。
 至少包含 一个大写英文 字母。
 至少包含 一个数字 。
 至少包含 一个特殊字符 。特殊字符为："!@#$%^&*()-+" 中的一个。
 它 不 包含 2 个连续相同的字符（比方说 "aab" 不符合该条件，但是 "aba" 符合该条件）。


 给你一个字符串 password ，如果它是一个 强 密码，返回 true，否则返回 false 。



 示例 1：

 输入：password = "IloveLe3tcode!"
输出：true
解释：密码满足所有的要求，所以我们返回 true 。


 示例 2：

 输入：password = "Me+You--IsMyDream"
输出：false
解释：密码不包含数字，且包含 2 个连续相同的字符。所以我们返回 false 。


 示例 3：

 输入：password = "1aB!"
输出：false
解释：密码不符合长度要求。所以我们返回 false 。



 提示：


 1 <= password.length <= 100
 password 包含字母，数字和 "!@#$%^&*()-+" 这些特殊字符。


 Related Topics 字符串 👍 15 👎 0

*/

package leetcode

func Q2299Main() {
	strongPasswordCheckerII("ziyS5FrPQhoQ5oEWRpHW2MjI7sGfcMJdcsjQnIyRbdCilvFaQN07jQtAkOklbjFrU5KcHzw4EvJ41Yz2Ykyd")
}

// leetcode submit region begin(Prohibit modification and deletion)
func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}

	checkList := [4]bool{}
	// 0 一个小写英文
	// 1 一个大写英文
	// 2 一个数字
	// 3 一个特殊字符

	for i, char := range password {
		// 不包含 2 个连续相同的字符
		if i > 0 && byte(char) == password[i-1] {
			return false
		}

		// 0 一个小写英文
		if char >= 'a' && char <= 'z' {
			checkList[0] = true
			continue
		}

		// 1 一个大写英文
		if char >= 'A' && char <= 'Z' {
			checkList[1] = true
			continue
		}

		// 2 一个数字
		if char >= '0' && char <= '9' {
			checkList[2] = true
			continue
		}

		// 3 一个特殊字符
		checkList[3] = true
	}

	return checkList[0] && checkList[1] && checkList[2] && checkList[3]
}

// leetcode submit region end(Prohibit modification and deletion)
