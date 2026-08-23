package trie

/*
Trie
- Every trie node contains 26 Trie nodes
- Whenever a new character is pattern is encountered we will create a new node and attach to current idx
- If we have traversed entire word, then we will mark last node with wordEnds as True
- If we want to search for exact match we search until wordEnds appear as True and idx reaches end
- If we want to just search the prefix we only search until idx reaches end
*/

type Trie struct {
	node     [26]*Trie
	wordEnds bool
}

func Constructor() Trie {
	return Trie{
		node:     [26]*Trie{},
		wordEnds: false,
	}
}

func recursiveInsert(trie *Trie, idx int, word string) {
	if idx == len(word) {
		trie.wordEnds = true
		return
	}
	if trie.node[word[idx]-'a'] == nil {
		newTrie := Trie{
			node:     [26]*Trie{},
			wordEnds: false,
		}
		trie.node[word[idx]-'a'] = &newTrie
	}
	recursiveInsert(trie.node[word[idx]-1], idx+1, word)
}

func (this *Trie) Insert(word string) {
	recursiveInsert(this, 0, word)
}

func recursiveExact(trie *Trie, idx int, word string) bool {
	if trie == nil {
		return false
	}

	if idx == len(word) {
		return trie.wordEnds
	}

	return recursiveExact(trie.node[word[idx]-'a'], idx+1, word)
}

func (this *Trie) SearchExactWord(word string) bool {
	return recursiveExact(this, 0, word)
}

func recursivePrefix(trie *Trie, idx int, word string) bool {

	if idx == len(word) {
		return true
	}

	if trie == nil {
		return false
	}

	return recursivePrefix(trie.node[word[idx]-'a'], idx+1, word)
}

func (this *Trie) PrefixMatch(prefix string) bool {
	return recursivePrefix(this, 0, prefix)
}
