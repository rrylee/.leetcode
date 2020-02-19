package main

func main() {
	m := Constructor()
	m.Remove(14)
}

/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] 设计哈希映射
 */

// @lc code=start
type MyHashMap struct {
	data []*node
	len  int
}

type node struct {
	key   int
	value int
	next  *node
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		data: make([]*node, 10001),
		len:  10001,
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	index := key % this.len
	data := this.data[index]
	if data == nil {
		node := &node{
			key:   key,
			value: value,
		}
		this.data[index] = node
	} else {
		for data != nil {
			if data.key == key {
				data.value = value
				return
			}
			data = data.next
		}
		node := &node{
			key:   key,
			value: value,
			next:  this.data[index],
		}
		this.data[index] = node
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	index := key % this.len
	data := this.data[index]
	for data != nil {
		if data.key == key {
			return data.value
		}
		data = data.next
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	index := key % this.len
	data := this.data[index]
	if data == nil {
		return
	}

	if data.key == key {
		this.data[index] = data.next
		return
	}

	for data.next != nil {
		if data.next.key == key {
			data.next = data.next.next
			return
		}
		data = data.next
	}
	return
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end
