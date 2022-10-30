/**
你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本
都是错的。

 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。

 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版
本。你应该尽量减少对调用 API 的次数。

 示例 1：


输入：n = 5, bad = 4
输出：4
解释：
调用 isBadVersion(3) -> false
调用 isBadVersion(5) -> true
调用 isBadVersion(4) -> true
所以，4 是第一个错误的版本。


 示例 2：


输入：n = 1, bad = 1
输出：1




 提示：


 1 <= bad <= n <= 2³¹ - 1


 Related Topics 二分查找 交互 👍 814 👎 0

*/

package leetcode

import (
    "fmt"
)

var q278TestCase []bool

func Q278Main() {
    n, bad := 2126753390, 1702766719
    q278TestCase = make([]bool, n+1)

    for i := bad; i < n+1; i++ {
        q278TestCase[i] = true
    }

    fmt.Printf("结果：%v\n", firstBadVersion(n))

}

func isBadVersion(version int) bool {
    if version == 0 {
        panic("version == 0 err.")
    }
    return q278TestCase[version]
}

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    if isBadVersion(1) {
        return 1
    }

    min, mid, max := 1, n/2, n

    // times := 0
    for {
        // fmt.Printf("times: %v ,interregional[%v,%v] ,new mid: %v\n", times, min, max, mid)
        // times++
        // time.Sleep(time.Second * 1)
        if isBadVersion(mid) {
            if !isBadVersion(mid - 1) {
                return mid
            }
            max = mid
            mid = (max + min) / 2
        } else {
            if isBadVersion(mid + 1) {
                return mid + 1
            }
            min = mid
            mid = (max + min) / 2
        }

    }

}

// leetcode submit region end(Prohibit modification and deletion)
