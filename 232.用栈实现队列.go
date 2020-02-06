/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	num1 []int
	num2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		num1: []int{},
		num2: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.num1 = append(this.num1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.num2) == 0 {
		this.adjust()
	}
	x := this.num2[len(this.num2)-1]
	this.num2 = this.num2[:len(this.num2)-1]
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.num2) != 0 {
		return this.num2[len(this.num2)-1]
	} else {
		this.adjust()
	}
	x := this.num2[len(this.num2)-1]
	return x
}

func (this *MyQueue) adjust() {
	if len(this.num1) != 0 {
		for i := len(this.num1) - 1; i >= 0; i-- {
			this.num2 = append(this.num2, this.num1[i])
		}
		this.num1 = []int{}
	}
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.num1) == 0 && len(this.num2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

