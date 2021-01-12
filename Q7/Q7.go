/**
 * @Author HatsuneMona
 * @Date  2020-11-14 10:06
 * @Description //TODO
 **/
/*给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
示例 1:          示例 2:			示例 3:
输入: 123		输入: -123		输入: 120
输出: 321		输出: -321		输出: 21
注意:
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31, 2^31 − 1]。
请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
package Q7

import "strconv"

//-2^31   = -2,147,483,648
// 2^31-1 =  2,147,483,647
var (
	min = -2147483648
	max = 2147483647
)

func reverse(x int) int {
	isMinus := x < 0
	if isMinus {
		x = -x
	}
	strArr := []rune(strconv.Itoa(x))
	for i := 0; i < len(strArr)/2; i++ {
		strArr[i], strArr[len(strArr)-i-1] = strArr[len(strArr)-i-1], strArr[i]
	}
	y, err := strconv.Atoi(string(strArr))
	if y <= min || y >= max || err != nil { //有可能会溢出
		return 0
	}
	if isMinus {
		y = -y
	}
	return y
}
