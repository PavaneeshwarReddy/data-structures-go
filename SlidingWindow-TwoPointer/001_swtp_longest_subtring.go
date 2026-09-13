package slidingwindowtwopointer

/*
Longest Substring without Repeating Characters
- Start with the same positions start and end and then move
- Maintain a count map that can maintain counts of each character
- Range always give you runes and iterating over idices will give you byte


NOTE:
- Sliding windows only works during predictable values
- I mean only positive values, if end increases value increases and start increase value decreases
*/

func checkDuplicate(arr map[byte]int) bool {
	for _, val := range arr {
		if val > 1 {
			return true
		}
	}

	return false
}

func lengthOfLongestSubstring(s string) int {
	countMap := make(map[byte]int)
	maxLength := 0
	start := 0
	end := 0

	for start <= end && end < len(s) {
		if val, ok := countMap[s[end]]; ok {
			countMap[s[end]] = val + 1
		} else {
			countMap[s[end]] = 1
		}
		end++
		repeating := checkDuplicate(countMap)
		if !repeating {
			maxLength = max(maxLength, end-start)
		} else {
			if val, ok := countMap[s[start]]; ok {
				countMap[s[start]] = val - 1
			}
			start++
		}
	}

	return maxLength
}
