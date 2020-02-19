package main

import (
	"fmt"
)

func main() {
	list := Constructor()
	list.AddAtHead(84)
	list.AddAtTail(2)
	list.AddAtTail(39)
	fmt.Println(list.Get(3))
	fmt.Println(list.Get(1))
	list.AddAtTail(42)
	list.AddAtIndex(1, 80)
	list.AddAtHead(14)
	list.AddAtHead(1)
	list.AddAtTail(53)
	list.AddAtTail(98)
	list.AddAtTail(19)
	list.AddAtTail(12)
	fmt.Println(list.Get(2))
}

// 1 14 84 80  2 39 42

/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 */

// @lc code=start
type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  int
	next *Node
	prev *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	list := MyLinkedList{
		head: &Node{},
		tail: &Node{},
		size: 0,
	}
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	if index+1 < this.size-index {
		cur := this.head
		for i := 0; i <= index; i++ {
			cur = cur.next
		}
		return cur.val
	} else {
		cur := this.tail
		for i := 0; i < this.size-index; i++ {
			cur = cur.prev
		}
		return cur.val
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	prev, next := this.head, this.head.next
	this.size += 1
	node := &Node{
		val:  val,
		next: next,
		prev: prev,
	}
	prev.next = node
	next.prev = node
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	prev, next := this.tail.prev, this.tail
	this.size += 1
	node := &Node{
		val:  val,
		next: next,
		prev: prev,
	}
	prev.next = node
	next.prev = node
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}
	if index > this.size {
		return
	}

	var prev, next *Node
	if index+1 < this.size-index {
		prev, next = this.head, this.head.next
		for i := 0; i < index; i++ {
			prev = next
			next = next.next
		}
	} else {
		prev, next = this.tail.prev, this.tail
		for i := 0; i < this.size-index; i++ {
			next = prev
			prev = prev.prev
		}
	}

	node := &Node{
		val:  val,
		next: next,
		prev: prev,
	}
	prev.next = node
	next.prev = node
	this.size += 1
	return
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	node := this.head.next
	for i := 0; i < index; i++ {
		node = node.next
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	this.size -= 1
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end
