package trie

/*
Longest valid word with all prefixes present in array
- For every string we need to check whether each character addition has wordEnds true
*/

func recursiveAllPrefixEnds(trie *Trie, idx int, word string) bool {
	if trie == nil || len(word) == idx {
		return true
	}

	return trie.wordEnds && recursiveAllPrefixEnds(trie.node[word[idx]-'a'], idx+1, word)
}

func LongestValidWord(words []string) string {
	maxLength := 0
	result := ""

	trie := Trie{
		node:     [26]*Trie{},
		wordEnds: false,
	}

	for _, val := range words {
		trie.Insert(val)
	}

	for _, val := range words {
		if recursiveAllPrefixEnds(&trie, 0, val) {
			if len(val) > maxLength {
				result = val
				maxLength = len(val)
			}
		}
	}
	return result
}
