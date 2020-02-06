/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := []byte{}
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; !ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == v {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end

