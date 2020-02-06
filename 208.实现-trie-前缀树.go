package main

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	Children [26]*Trie
	End      bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		v := cur.Children[index]
		if v == nil {
			newTrie := Constructor()
			cur.Children[index] = &newTrie
		}
		cur = cur.Children[index]
	}
	cur.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		v := cur.Children[index]
		if v == nil {
			return false
		}
		cur = v
	}
	if cur.End {
		return true
	}
	for _, trie := range cur.Children {
		if trie != nil {
			return false
		}
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		v := cur.Children[index]
		if v == nil {
			return false
		}
		cur = v
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
