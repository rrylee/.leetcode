package main

import "fmt"

func main() {
	m := Constructor()
	m.Push(-2)
	m.Push(0)
	m.Push(-3)
	fmt.Println(m.GetMin())
	m.Pop()
	fmt.Println(m.Top())
	fmt.Println(m)
	fmt.Println(m.GetMin())
}

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	data   []int
	helper []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:   []int{},
		helper: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.helper) == 0 || this.helper[len(this.helper)-1] >= x {
		this.helper = append(this.helper, x)
	}
}

func (this *MinStack) Pop() {
	x := this.Top()
	this.data = this.data[:len(this.data)-1]
	if x == this.GetMin() {
		this.helper = this.helper[:len(this.helper)-1]
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.helper[len(this.helper)-1]
}

/*
type MinStack struct {
	data []int
}

func Constructor() MinStack {
	return MinStack{
		data: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.data) == 1 || this.data[len(this.data)-2] > x {
		this.data = append(this.data, x)
	} else {
		this.data = append(this.data, this.data[len(this.data)-2])
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-2]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-2]
}

func (this *MinStack) GetMin() int {
	return this.data[len(this.data)-1]
}
*/

/*
type MinStack struct {
	data   []int
	helper []int
}

func Constructor() MinStack {
	return MinStack{
		data:   []int{},
		helper: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.helper) == 0 || this.helper[len(this.helper)-1] > x {
		this.helper = append(this.helper, x)
	} else {
		this.helper = append(this.helper, this.helper[len(this.helper)-1])
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.helper = this.helper[:len(this.helper)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.helper[len(this.helper)-1]
}
*/

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

// -2 0 -3
// -2 -2 -3
