/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

// @lc code=start
type LRUCache struct {
	cap  int
	data map[int]*Node
	head *Node
	tail *Node
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cap:  capacity,
		data: make(map[int]*Node, capacity),
		head: &Node{},
		tail: &Node{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.data[key]; !ok {
		return -1
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
		curNode := this.head.next
		this.head.next = node
		node.prev = this.head
		curNode.prev = node
		node.next = curNode
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.data[key]; ok {
		node.value = value
		node.prev.next = node.next
		node.next.prev = node.prev
		curNode := this.head.next
		this.head.next = node
		node.prev = this.head
		curNode.prev = node
		node.next = curNode
		return
	}

	if len(this.data) < this.cap {
		node := &Node{
			key:   key,
			value: value,
		}
		this.data[key] = node
		curHead := this.head.next
		this.head.next = node
		node.prev = this.head
		node.next = curHead
		curHead.prev = node
		return
	} else {
		tail := this.tail.prev
		this.tail.prev.prev.next = this.tail
		this.tail.prev = this.tail.prev.prev
		delete(this.data, tail.key)
		tail.key = key
		tail.value = value
		this.data[key] = tail
		curHead := this.head.next
		this.head.next = tail
		tail.prev = this.head
		tail.next = curHead
		curHead.prev = tail
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
