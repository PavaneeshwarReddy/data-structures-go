package trie

/*
Distinct Substrings
- Recursively add all combinations to the result array
*/

func DistinctSubstrings(trie *Trie, currWord string, result *[]string) {
	if trie == nil {
		return
	}

	*result = append(*result, currWord)

	for idx, val := range trie.node {
		if val != nil {
			DistinctSubstrings(val, currWord+string(rune('a'+idx)), result)
		}

	}
}
