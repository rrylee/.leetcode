/*
 * @lc app=leetcode.cn id=1206 lang=golang
 *
 * [1206] 设计跳表
 */

// @lc code=start
type Skiplist struct {
	maxLevel int
	head     *Node
}

type Node struct {
	value int
	next  []*Node
}

func Constructor() Skiplist {
	list := Skiplist{
		maxLevel: -1,
		head
	}
}

func NewNode(level, value int) *Node {
	return &Node{
		level: level,
		value: value,
		next:  make([]*Node, level+1),
		prev:  make([]*Node, level+1),
		count: 1,
	}
}

func (this *Skiplist) Search(target int) bool {

}

func (this *Skiplist) Add(num int) {

}

func (this *Skiplist) Erase(num int) bool {

}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
// @lc code=end

