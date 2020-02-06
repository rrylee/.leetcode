/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] 设计哈希集合
 */

// @lc code=start
type MyHashSet struct {
	data []*Node
	len  int
}

type Node struct {
	key  int
	next *Node
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		data: make([]*Node, 10001),
		len:  10001,
	}
}

func (this *MyHashSet) Add(key int) {
	data := this.data[key%this.len]
	if data == nil {
		this.data[key%this.len] = &Node{
			key: key,
		}
	} else {
		if data.key == key {
			return
		}

		for data.next != nil {
			data = data.next
			if data.key == key {
				return
			}
		}

		data.next = &Node{
			key: key,
		}
	}
}

func (this *MyHashSet) Remove(key int) {
	data := this.data[key%this.len]
	if data == nil {
		return
	}
	if data.key == key {
		this.data[key%this.len] = data.next
		return
	}
	for data.next != nil {
		if data.next.key != key {
			data = data.next
		} else {
			data.next = data.next.next
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	data := this.data[key%this.len]
	if data == nil {
		return false
	}
	for data.next != nil {
		if data.key == key {
			return true
		} else {
			data = data.next
		}
	}
	return data.key == key
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

