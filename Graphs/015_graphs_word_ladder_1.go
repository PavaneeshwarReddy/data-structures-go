package graphs

/*
Word ladder
- First we need to consider, a set of words
- We should take each word and manipulate all it's character one by one and check whether that is exisiting in the word list
- If exisits just delete and proceed.
- Then if the level matches then we can simply, return the length

*/

type Word struct {
	Word  []byte
	Level int
}

func convertCheck(word []byte, words map[string]bool) [][]byte {
	results := [][]byte{}

	for i := range word {
		tempWord := append([]byte(nil), word...)
		for j := range 26 {
			tempWord[i] = byte(j) + 'a'
			if _, ok := words[string(tempWord)]; ok {
				newTempWord := append([]byte(nil), tempWord...)
				results = append(results, newTempWord)
				delete(words, string(tempWord))
			}
		}
	}

	return results
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := make(map[string]bool)

	for _, word := range wordList {
		if word != beginWord {
			words[word] = true
		}
	}

	stack := []Word{{Word: []byte(beginWord), Level: 1}}

	for len(stack) > 0 {
		newWords := []Word{}
		for _, word := range stack {
			if string(word.Word) == endWord {
				return word.Level
			}
			changeResults := convertCheck(word.Word, words)
			for _, res := range changeResults {
				newWords = append(newWords, Word{Word: res, Level: word.Level + 1})
			}
		}

		stack = newWords
	}

	return 0

}
