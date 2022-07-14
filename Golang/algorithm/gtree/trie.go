package gtree

type WordFilter struct {
	prefixTrie *trie
	suffixTrie *trie
}

// 只包含小写英文字母
type trie struct {
	idxs     map[int]struct{} // 保存words的下标
	children []*trie
}

func Constructor(words []string) WordFilter {
	prefix, suffix := &trie{}, &trie{}
	for idx, word := range words {
		n := len(word)
		curPrefix, curSuffix := prefix, suffix
		for i := 0; i < n; i++ {
			preCh, sufCh := word[i], word[n-i-1]

			if len(curPrefix.children) == 0 {
				curPrefix.children = make([]*trie, 26)
			}
			curPrefix.children[preCh-'a'] = &trie{}
			curPrefix = curPrefix.children[preCh-'a']
			if curPrefix.idxs == nil {
				curPrefix.idxs = map[int]struct{}{}
			}
			curPrefix.idxs[idx] = struct{}{}

			if len(curSuffix.children) == 0 {
				curSuffix.children = make([]*trie, 26)
			}
			curSuffix.children[sufCh-'a'] = &trie{}
			curSuffix = curSuffix.children[sufCh-'a']
			if curSuffix.idxs == nil {
				curSuffix.idxs = map[int]struct{}{}
			}
			curSuffix.idxs[idx] = struct{}{}
		}
	}
	return WordFilter{
		prefixTrie: prefix,
		suffixTrie: suffix,
	}
}

func (this *WordFilter) F(pref string, suff string) int {
	curPrefix := this.prefixTrie
	for _, ch := range pref {
		curPrefix = curPrefix.children[ch-'a']
		if curPrefix == nil {
			return -1
		}
	}

	curSuffix := this.suffixTrie
	for i, n := 0, len(suff); i < n; i++ {
		ch := suff[n-i-1]
		curSuffix = curSuffix.children[ch-'a']
		if curSuffix == nil {
			return -1
		}
	}

	res := -1
	for key := range curPrefix.idxs {
		if _, ok := curSuffix.idxs[key]; ok {
			if key > res {
				res = key
			}
		}
	}
	return res
}

type WordFilter2 struct {
	pre *TrieNode
	suf *TrieNode
}

type TrieNode struct {
	ind []int
	arr [26]*TrieNode
}

func Constructor2(words []string) WordFilter2 {
	pre := &TrieNode{ind: []int{}, arr: [26]*TrieNode{}}
	suf := &TrieNode{ind: []int{}, arr: [26]*TrieNode{}}
	for i, word := range words {
		build := pre
		for _, b := range word {
			num := b - 'a'
			if build.arr[num] == nil {
				build.arr[num] = &TrieNode{ind: []int{}, arr: [26]*TrieNode{}}
			}
			build = build.arr[num]
			build.ind = append(build.ind, i)
		}
	}
	for j, word := range words {
		build := suf
		l := len(word)
		for i := l - 1; i >= 0; i-- {
			num := rune(word[i]) - 'a'
			if build.arr[num] == nil {
				build.arr[num] = &TrieNode{ind: []int{}, arr: [26]*TrieNode{}}
			}
			build = build.arr[num]
			build.ind = append(build.ind, j)
		}
	}
	return WordFilter2{pre: pre, suf: suf}
}

func (this *WordFilter2) F(pref string, suff string) int {
	pre := this.pre
	suf := this.suf
	var preArr []int
	var sufArr []int
	for _, b := range pref {
		num := b - 'a'
		pre = pre.arr[num]
		if pre == nil {
			return -1
		}
	}
	preArr = pre.ind
	for i := len(suff) - 1; i >= 0; i-- {
		num := rune(suff[i]) - 'a'
		suf = suf.arr[num]
		if suf == nil {
			return -1
		}
	}
	sufArr = suf.ind
	for i, j := len(preArr)-1, len(sufArr)-1; i >= 0 && j >= 0; {
		if preArr[i] == sufArr[j] {
			return preArr[i]
		}
		if preArr[i] > sufArr[j] {
			i--
		} else {
			j--
		}
	}
	return -1
}
