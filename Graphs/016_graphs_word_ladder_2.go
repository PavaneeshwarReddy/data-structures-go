package graphs

/*
WordAdv ladder - 2
- Same as word ladder-1 but here instead of just count we need to return the sequence
- While storing sequence, we should remove the words that are generated and used at that level
- This ignore storing of multiple sequences again and again.
- If we found the result then we should stop further and just return the results.
*/
type WordAdv struct {
	WordAdv  []byte
	Sequence [][]byte
	Level    int
}

func convertCheck2(word []byte, words map[string]bool) [][]byte {
	results := [][]byte{}

	for i := range word {
		tempWord := append([]byte(nil), word...)

		for j := range 26 {
			newChar := byte(j) + 'a'

			if tempWord[i] == newChar {
				continue
			}

			tempWord[i] = newChar

			if _, ok := words[string(tempWord)]; ok {
				newTempWord := append([]byte(nil), tempWord...)
				results = append(results, newTempWord)
			}
		}
	}

	return results
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	words := make(map[string]bool)

	for _, word := range wordList {
		if word != beginWord {
			words[word] = true
		}
	}

	// If endWord doesn't exist, there is no possible transformation.
	if !words[endWord] {
		return [][]string{}
	}

	stack := []WordAdv{
		{
			WordAdv:  []byte(beginWord),
			Sequence: [][]byte{[]byte(beginWord)},
			Level:    1,
		},
	}

	globalResults := [][]string{}

	for len(stack) > 0 {

		newWords := []WordAdv{}

		visitedThisLevel := make(map[string]bool)
		foundEndWord := false

		for _, word := range stack {

			if string(word.WordAdv) == endWord {

				result := []string{}

				for _, seq := range word.Sequence {
					result = append(result, string(seq))
				}

				globalResults = append(globalResults, result)

				foundEndWord = true
				continue
			}

			changeResults := convertCheck2(word.WordAdv, words)

			for _, res := range changeResults {

				resString := string(res)
				newSequence := append(
					append([][]byte(nil), word.Sequence...),
					res,
				)

				newWord := WordAdv{
					WordAdv:  res,
					Sequence: newSequence,
					Level:    word.Level + 1,
				}

				newWords = append(newWords, newWord)

				visitedThisLevel[resString] = true
			}
		}

		if foundEndWord {
			break
		}

		for word := range visitedThisLevel {
			delete(words, word)
		}

		stack = newWords
	}

	return globalResults
}
