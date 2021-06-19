package main

func multiSearch(big string, smalls []string) [][]int {
	var (
		trie *Trie
		t    = Constructor()
		ans  = make([][]int, len(smalls))
		res  []int
	)
	trie = &t
	for i := 0; i < len(smalls); i++ {
		trie.Insert(smalls[i], i)
	}
	for i := 0; i < len(big); i++ {
		res = trie.Search(big[i:])
		for j := 0; j < len(res); j++ {
			ans[res[j]] = append(ans[res[j]], i)
		}
	}
	return ans
}

type Trie struct {
	idx   int
	nodes map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{nodes: map[byte]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string, idx int) {
	var curNode = this
	for i := 0; i < len(word); i++ {
		if _, has := curNode.nodes[word[i]]; !has {
			var t = Constructor()
			curNode.nodes[word[i]] = &t
		}
		curNode = curNode.nodes[word[i]]
	}
	// 为了区分index为0还是不是一个匹配字符串
	curNode.idx = idx + 1
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) []int {
	var (
		curNode = this
		ans     []int
	)
	for i := 0; i < len(word); i++ {
		if _, has := curNode.nodes[word[i]]; !has {
			//fmt.Printf("word: %s, i: %d \n", word, i)
			return ans
		}
		if curNode.nodes[word[i]].idx > 0 {
			//fmt.Printf("idx: %d \n", curNode.nodes[word[i]].idx)
			ans = append(ans, curNode.nodes[word[i]].idx-1)
		}
		curNode = curNode.nodes[word[i]]
	}
	return ans
}

// https://leetcode-cn.com/problems/multi-search-lcci/
//类似于这样的场景：给你一个长句子，再给你一堆“敏感词”，然后让你找敏感词在句子里的位置（因为要把敏感词换成 ***）。

//把敏感词 smalls 的数量记为 t，把敏感词里最长的字符串长度记为 k，把长句子 big 的长度记为 b。
//具体步骤：
//1）把这堆敏感词建成一颗 Trie 树，时间复杂度是 O(tk)。
//2）遍历长句子的每一个字母，检查“以该字母作为起点”的话，是否可以在 trie 中找到结果。时间复杂度是 O(bk)
//综上，总的时间复杂度是 O(tk + bk)。
