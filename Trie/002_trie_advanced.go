package trie

/*
Trie
- words counter,for every character we see we increment counter by 1
- Exact words count, just return the wordEnds the count there. ( additional condition words ends should be true)
- Prefix words count, reach to that point and just return the count.
- Delete word, decrement all occurances of this, if count becomes 0 then make the node nil
*/

type Trie2 struct {
	node     [26]*Trie2
	wordEnds bool
	count    int
}

func Constructor2() Trie2 {
	return Trie2{
		node:     [26]*Trie2{},
		wordEnds: false,
		count:    0,
	}
}

func recursiveInsert2(trie *Trie2, idx int, word string) {
	if idx == len(word) {
		trie.wordEnds = true
		return
	}
	if trie.node[word[idx]-'a'] == nil {
		newTrie := Trie2{
			node:     [26]*Trie2{},
			wordEnds: false,
			count:    1,
		}
		trie.node[word[idx]-'a'] = &newTrie
	} else {
		trie.count += 1
	}
	recursiveInsert2(trie.node[word[idx]-1], idx+1, word)
}

func (this *Trie2) Insert2(word string) {
	recursiveInsert2(this, 0, word)
}

func recursiveCountWords(trie *Trie2, idx int, word string, result *int) int {
	if trie == nil {
		return 0
	}

	if idx == len(word) {
		return trie.count
	}

	return recursiveCountWords(trie.node[word[idx]-'a'], idx+1, word, result)
}

func (this *Trie2) CountWords(word string) int {
	result := 0
	return recursiveCountWords(this, 0, word, &result)
}

func recursiveDelete(trie *Trie2, idx int, word string) bool {
	if trie == nil {
		return false
	}

	if idx == len(word) {
		return trie.wordEnds
	}

	canDelete := recursiveDelete(trie.node[word[idx]-'a'], idx+1, word)
	if canDelete {
		trie.count -= 1
	}
	return canDelete
}

func (this *Trie2) DeleteWord(word string) {
	recursiveDelete(this, 0, word)
}
