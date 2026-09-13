package slidingwindowtwopointer

/*
Longest Repeating Character Replacement
- Always think about condition to be checked while moving the sliding window
- In this case, if length of current subarray - max Freq is the left over that needs to be flipped so that should be less then k flips
*/

func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)
	start, end, res, maxFreq := 0, 0, 0, 0

	for end < len(s) {
		freqMap[s[end]]++
		maxFreq = max(maxFreq, freqMap[s[end]])

		for (end-start+1)-maxFreq > k {
			freqMap[s[start]]--
			if freqMap[s[start]] == 0 {
				delete(freqMap, s[start])
			}
			maxFreq = 0
			for _, val := range freqMap {
				maxFreq = max(maxFreq, val)
			}
			start++
		}

		res = max(res, end-start+1)
		end += 1

	}

	return res
}
