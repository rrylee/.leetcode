/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	count := 0
	countNode := head
	for count < k && countNode != nil {
		count += 1
		countNode = countNode.Next
	}
	if count == k {
		i := 0
		node := head
		var next, prev *ListNode
		for i < k && node != nil {
			next = node.Next
			node.Next = prev
			prev = node
			node = next
			i += 1
		}
		head.Next = reverseKGroup(next, k)
		return prev
	} else {
		return head
	}
}

// @lc code=end

