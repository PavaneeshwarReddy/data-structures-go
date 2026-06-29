package strings

/*
Count Substrings with K distinct character
	- Traverse all substrings
	- Calculate map frequency count
	- Check the length of the map
*/

func CountSubstringKDistinctChars(s string, k int) int {
	result := 0

	for i := 0; i < len(s); i++ {
		mp := make(map[byte]int)
		for j := i; j < len(s); j++ {
			mp[s[j]]++
			if len(mp) == k {
				result += 1
			}
		}
	}
	return result
}
