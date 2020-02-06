package main

import "fmt"

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	node = swapPairs(node)

	fmt.Println(node)
	fmt.Println(node.Next)
	fmt.Println(node.Next.Next)
	fmt.Println(node.Next.Next.Next)
}

// @lc code=start
//递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

// 迭代
func swapPairs1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	newHead := head.Next
	for head != nil && head.Next != nil {
		cur := head
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		head = next

		if head != nil {
			if head.Next != nil {
				cur.Next = head.Next
			} else {
				cur.Next = head
			}
		}
	}
	return newHead
}

// @lc code=end
