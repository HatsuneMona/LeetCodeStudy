/**
给定一个包含大写字母和小写字母的字符串
 s ，返回 通过这些字母构造成的 最长的回文串 。

 在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。



 示例 1:


输入:s = "abccccdd"
输出:7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。


 示例 2:


输入:s = "a"
输入:1




 提示:


 1 <= s.length <= 2000
 s 只由小写 和/或 大写英文字母组成


 Related Topics 贪心 哈希表 字符串 👍 475 👎 0

*/

package leetcode

import (
    "fmt"
)

func Q409Main() {
    num := int(^uint(0) >> 2 << 1)
    fmt.Printf("num = %v     %08b\n", num, num)
    
    str := "abccccdd"
    fmt.Printf("输出结果：%v\n", longestPalindrome(str))
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) int {
    
    if len(s) < 2 {
        return len(s)
    }
    
    // 哈希表解法
    charStat := make([]int, 64)
    
    for _, char := range []byte(s) {
        charStat[char-65]++
    }
    
    flag := false
    length := 0
    for _, num := range charStat {
        if num != 0 {
            length += int(^uint(0)>>2-1) & num
            // fmt.Println("num = ", num, "     int(^uint(0)>>2<<1) & num = ", int(^uint(0)>>2<<1)&num)
            flag = flag || num%2 == 1
        }
    }
    
    if flag {
        length++
    }
    
    return length
    
}

// leetcode submit region end(Prohibit modification and deletion)
