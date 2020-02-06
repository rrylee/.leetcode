/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 */

// @lc code=start
type MapSum struct {
	sum      int
	Children [26]*MapSum
	End      bool
	Val      int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this
	pos := []byte{}
	for i := 0; i < len(key); i++ {
		pos = append(pos, key[i]-'a')
		v := cur.Children[key[i]-'a']
		if v == nil {
			cur.Children[key[i]-'a'] = &MapSum{}
		} else {
			//v.Num += val
		}
		cur = cur.Children[key[i]-'a']
	}

	if cur.End {
		oldVal := cur.Val
		cur2 := this
		for _, p := range pos {
			cur2.Children[p].sum += val - oldVal
			cur2 = cur2.Children[p]
		}
	} else {
		cur2 := this
		for _, p := range pos {
			cur2.Children[p].sum += val
			cur2 = cur2.Children[p]
		}

		cur.End = true
	}
	cur.Val = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this
	for i := 0; i < len(prefix); i++ {
		v := cur.Children[prefix[i]-'a']
		if v == nil {
			return 0
		}
		cur = v
	}
	return cur.sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end

