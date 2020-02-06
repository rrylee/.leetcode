/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node1 := head
	node2 := head

	for i := 0; i < n; i++ {
		node1 = node1.Next
	}
	if node1 == nil {
		return head.Next
	}
	for node1.Next != nil {
		node1 = node1.Next
		node2 = node2.Next
	}
	node2.Next = node2.Next.Next
	return head
}

// @lc code=end

