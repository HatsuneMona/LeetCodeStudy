/**
给定两个字符串 s 和 t ，判断它们是否是同构的。

 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。



 示例 1:


输入：s = "egg", t = "add"
输出：true


 示例 2：


输入：s = "foo", t = "bar"
输出：false

 示例 3：


输入：s = "paper", t = "title"
输出：true



 提示：





 1 <= s.length <= 5 * 10⁴
 t.length == s.length
 s 和 t 由任意有效的 ASCII 字符组成


 Related Topics 哈希表 字符串 👍 535 👎 0

*/

package leetcode

import (
	"fmt"
)

func Q205Main() {
	s, t := "paper", "title"
	fmt.Printf("测试结果：%v\n", isIsomorphic(s, t))
}

// leetcode submit region begin(Prohibit modification and deletion)
func isIsomorphic(s string, t string) bool {
	cMap := make([]byte, 128)
	exist := make([]bool, 128)

	// b := 'a'
	// ttttt := make(map[byte]byte, len(s))
	//
	// fmt.Printf("slice 内存分配 大小：%v\n", (unsafe.Sizeof(b)*50+unsafe.Sizeof(cMap))*2)
	// fmt.Printf("map   内存分配 大小：%v\n", unsafe.Sizeof(ttttt)+unsafe.Sizeof(b)*2*8*128)

	for i := 0; i < len(s); i++ {
		sChar := []byte(s)[i]
		tChar := []byte(t)[i]

		if cMap[sChar] == 0 && !exist[tChar] {
			// 如果 sChar 没有映射关系
			// sChar -> ?

			// 并且 tChar 也没有被其他 sChar 所映射过
			// ? -> tChar
			// （满足  不同字符不能映射到同一个字符）
			// 此时可以添加映射关系
			cMap[sChar] = tChar
			exist[tChar] = true

		} else if cMap[sChar] != tChar {
			// 如果 sChar 已经有映射关系了
			// 但是已经存在的 映射关系c 与 事实上的tChar 冲突
			return false
		}

	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)
