/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists1(lists []*ListNode) *ListNode {
	length := len(lists)

	if length == 0 {
		return nil
	}

	for length > 1 {
		for i := 0; i < length/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[length-i-1])
		}
		length = (length + 1) / 2
	}

	return lists[0]
}

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

