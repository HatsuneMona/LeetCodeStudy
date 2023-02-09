/**
你需要设计一个包含验证码的验证系统。每一次验证中，用户会收到一个新的验证码，这个验证码在 currentTime 时刻之后 timeToLive 秒过期。如果验
证码被更新了，那么它会在 currentTime （可能与之前的 currentTime 不同）时刻延长 timeToLive 秒。

 请你实现 AuthenticationManager 类：


 AuthenticationManager(int timeToLive) 构造 AuthenticationManager 并设置 timeToLive 参数。
 generate(string tokenId, int currentTime) 给定 tokenId ，在当前时间 currentTime 生成一个新的验证码。
 renew(string tokenId, int currentTime) 将给定 tokenId 且 未过期 的验证码在 currentTime 时刻更新。
										如果给定 tokenId 对应的验证码不存在或已过期，请你忽略该操作，不会有任何更新操作发生。
 countUnexpiredTokens(int currentTime) 请返回在给定 currentTime 时刻，未过期 的验证码数目。


 如果一个验证码在时刻 t 过期，且另一个操作恰好在时刻 t 发生（renew 或者 countUnexpiredTokens 操作），过期事件 优先于 其他操作。



 示例 1：


输入：
["AuthenticationManager", "renew", "generate", "countUnexpiredTokens", "generate", "renew", "renew", "countUnexpiredTokens"]
[[5], ["aaa", 1], ["aaa", 2], [6], ["bbb", 7], ["aaa", 8], ["bbb", 10], [15]]
输出：
[null, null, null, 1, null, null, null, 0]


解释：
AuthenticationManager authenticationManager = new AuthenticationManager(5); // 构造 AuthenticationManager ，设置 timeToLive = 5 秒。
authenticationManager.renew("aaa", 1); // 时刻 1 时，没有验证码的 tokenId 为 "aaa" ，没有验证码被更新。
authenticationManager.generate("aaa", 2); // 时刻 2 时，生成一个 tokenId 为 "aaa" 的新验证码。
authenticationManager.countUnexpiredTokens(6); // 时刻 6 时，只有 tokenId 为 "aaa" 的验证码未过期，所以返回 1 。
authenticationManager.generate("bbb", 7); // 时刻 7 时，生成一个 tokenId 为 "bbb" 的新验证码。
authenticationManager.renew("aaa", 8); // tokenId 为 "aaa" 的验证码在时刻 7 过期，且 8 >= 7，所以时刻 8 的renew 操作被忽略，没有验证码被更新。
authenticationManager.renew("bbb", 10); // tokenId 为 "bbb" 的验证码在时刻 10 没有过期，所以renew 操作会执行，该 token 将在时刻 15 过期。
authenticationManager.countUnexpiredTokens(15); // tokenId 为 "bbb" 的验证码在时刻 15 过期，tokenId 为 "aaa" 的验证码在时刻 7 过期，所有验证码均已过期，所以返回 0 。




 提示：


 1 <= timeToLive <= 10⁸
 1 <= currentTime <= 10⁸
 1 <= tokenId.length <= 5
 tokenId 只包含小写英文字母。
 所有 generate 函数的调用都会包含独一无二的 tokenId 值。
 所有函数调用中，currentTime 的值 严格递增 。
 所有函数的调用次数总共不超过 2000 次。


 Related Topics 设计 哈希表 👍 71 👎 0

*/

package leetcode

import "fmt"

func Q1797Main() {
	//am := Constructor(28)
	//am.CountUnexpiredTokens(2)
	//am.Renew("xokiw", 6)
	//am.Generate("ofn", 7)
	//am.Renew("dses", 15)
	//am.CountUnexpiredTokens(17)
	//am.Renew("ofzu", 19)
	//am.Generate("dses", 20)
	//am.CountUnexpiredTokens(23)
	//am.CountUnexpiredTokens(27)
	//am.CountUnexpiredTokens(30)

	am := Constructor(13)
	am.Renew("ajvy", 1)
	am.CountUnexpiredTokens(3) // 0
	am.CountUnexpiredTokens(4) // 0
	am.Generate("fuzxq", 5)
	am.Generate("izmry", 7)
	am.Renew("puv", 12)
	am.Generate("ybiqb", 13)
	am.Generate("gm", 14)
	am.CountUnexpiredTokens(15) // 4
	am.CountUnexpiredTokens(18) // 3
	am.CountUnexpiredTokens(19) // 3
	am.Renew("ybiqb", 21)
	am.CountUnexpiredTokens(23) // 2
	am.CountUnexpiredTokens(25) // 2
	am.CountUnexpiredTokens(26) // 2
	//am.Head.PrintLinkedList()
	am.Generate("aqdm", 28) // 2
	//am.Head.PrintLinkedList()
	am.CountUnexpiredTokens(29)
	am.Renew("puv", 30)
}

func (l *LinkedList) PrintLinkedList() {
	i := 0
	for q := l.Prev; q != l; q = q.Prev {
		i++
		fmt.Printf("node_%v ：%+v\n", i, q)
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
type AuthenticationManager struct {
	Head           *LinkedList
	GlobalLiveTime int
}

type LinkedList struct {
	TokenId      string
	CurrentTime  int
	EnabledToken bool
	Prev         *LinkedList
	Next         *LinkedList
}

func Constructor(timeToLive int) AuthenticationManager {
	manager := AuthenticationManager{}
	manager.GlobalLiveTime = timeToLive

	begin := new(LinkedList)
	manager.Head = begin
	for i := 0; i < timeToLive && i <= 2000; i++ {
		node := new(LinkedList)
		node.Prev = manager.Head // 往回指
		manager.Head.Next = node // 往前指
		manager.Head = node
	}
	manager.Head.Next = begin
	begin.Prev = manager.Head

	return manager
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	// 所有 generate 函数的调用都会包含独一无二的 tokenId 值。
	// 所以不用检查之前有没有这个值

	this.Head.CurrentTime = currentTime
	this.Head.TokenId = tokenId
	this.Head.EnabledToken = true
	this.Head = this.Head.Next
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {

	for q := this.Head.Prev; q != this.Head; q = q.Prev {
		if q.EnabledToken == true && q.TokenId == tokenId && currentTime < q.CurrentTime+this.GlobalLiveTime {
			// 找到了 token 并且 时间可以被续期
			// 解约之前的 token
			q.EnabledToken = false
			// 续期
			this.Head.EnabledToken = true
			this.Head.TokenId = tokenId
			this.Head.CurrentTime = currentTime
			this.Head = this.Head.Next
			return
		}
		if currentTime >= q.CurrentTime+this.GlobalLiveTime || q.CurrentTime == 0 {
			break
		}
	}
	// 没有找到 token 或者不可以被续期
	this.Head.EnabledToken = false
	this.Head.CurrentTime = currentTime
	this.Head = this.Head.Next
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	count := 0
	for q := this.Head.Prev; q != this.Head; q = q.Prev {
		if q.EnabledToken == true && currentTime < q.CurrentTime+this.GlobalLiveTime {
			count++
			//fmt.Printf("countNode_%v：%+v\n", count, q)
		}
		if currentTime >= q.CurrentTime+this.GlobalLiveTime || q.CurrentTime == 0 {
			break
		}
	}

	this.Head.EnabledToken = false
	this.Head.CurrentTime = currentTime
	this.Head = this.Head.Next
	//fmt.Println("测试结果：", count)
	return count
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
//leetcode submit region end(Prohibit modification and deletion)

// 解答成功:
//	执行耗时:80 ms,击败了46.88% 的Go用户
//	内存消耗:7.4 MB,击败了40.63% 的Go用户

//type AuthenticationManager struct {
//	TokenList      map[string]int // tokenId表（ [token] = endTime  ）
//	GlobalLiveTime int
//}
//
//func Constructor(timeToLive int) AuthenticationManager {
//	return AuthenticationManager{
//		TokenList:      make(map[string]int),
//		GlobalLiveTime: timeToLive,
//	}
//}
//
//func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
//	// 所有 generate 函数的调用都会包含独一无二的 tokenId 值
//
//	this.TokenList[tokenId] = currentTime + this.GlobalLiveTime
//}
//
//func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
//	if endTime, exist := this.TokenList[tokenId]; exist && endTime > currentTime {
//		this.TokenList[tokenId] = currentTime + this.GlobalLiveTime
//	}
//}
//
//func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
//	count := 0
//	for _, endTime := range this.TokenList {
//		if endTime > currentTime {
//			count++
//		}
//	}
//	return count
//}
