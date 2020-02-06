/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	curNode := &ListNode{}
	dummyNode := curNode

	for l1 != nil || l2 != nil {
		if l1 == nil {
			curNode.Next = l2
			curNode = curNode.Next
			l2 = l2.Next
			continue
		}
		if l2 == nil {
			curNode.Next = l1
			curNode = curNode.Next
			l1 = l1.Next
			continue
		}
		if l1.Val < l2.Val {
			curNode.Next = l1
			curNode = curNode.Next
			l1 = l1.Next
			continue
		} else {
			curNode.Next = l2
			curNode = curNode.Next
			l2 = l2.Next
			continue
		}
	}
	return dummyNode.Next
}

// @lc code=end

