/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p1 := l1
	p2 := l2
	curr := dummyHead
	carry := 0

	for p1 != nil || p2 != nil {
		var x, y, sum int
		if p1 != nil {
			x = p1.Val
		}
		if p2 != nil {
			y = p2.Val
		}
		sum = x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}

// @lc code=end

