/**
 * @Author HatsuneMona
 * @Date  2020-11-12 15:08
 * @Description LeetCode 第225题
使用队列实现栈的下列操作：
push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:
你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-stack-using-queues
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 **/

package Q225

import "container/list"

type MyStack struct {
	size        int
	mainQueue   *list.List
	secondQueue *list.List
}

/** Initialize your data structure here. */
//构造函数，在这里初始化你的数据结构。
func Constructor() MyStack {
	return MyStack{
		size:        0,
		mainQueue:   list.New(),
		secondQueue: list.New(),
	}
}

/** Push element x onto queue. */
func (this *MyStack) Push(x int) {
	this.mainQueue.PushBack(x)
	this.size++
}

/** Removes the element on top of the queue and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		return -99999999
	}
	for i := 0; i < this.size-1; i++ {
		e := this.mainQueue.Front()
		this.secondQueue.PushBack(e.Value)
		this.mainQueue.Remove(e)
	}
	e := this.mainQueue.Front()
	this.swap()
	this.secondQueue.Init()
	this.size--
	return e.Value.(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		return -99999999
	}
	for i := 0; i < this.size-1; i++ {
		e := this.mainQueue.Front()
		this.secondQueue.PushBack(e.Value)
		this.mainQueue.Remove(e)
	}
	e := this.mainQueue.Front()
	this.secondQueue.PushBack(e.Value)
	this.swap()
	this.secondQueue.Init()
	return e.Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyStack) Empty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func (this *MyStack) swap() {
	t := this.secondQueue
	this.secondQueue = this.mainQueue
	this.mainQueue = t
}
