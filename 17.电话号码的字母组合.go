/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	ret := []string{}
	if digits == "" {
		return ret
	}

	ret = append(ret, "")
	for i := 0; i < len(digits); i++ {
		pre := ret
		ret = []string{}
		for _, num := range m[digits[i]] {
			for _, v := range pre {
				ret = append(ret, v+num)
			}
		}
	}

	return ret
}

// @lc code=end

