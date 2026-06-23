package strings

/*
Anagrams
	- If count every occurance of character matches with the other then both an anagrams
	- We need to take count of occurances of each character in first string and compare with second
*/

func Anagrams(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sCount := [26]int{}
	for _, val := range s {
		sCount[val-'a'] += 1
	}

	for _, val := range t {
		if sCount[val-'a'] == 0 {
			return false
		}
		sCount[val-'a'] -= 1
	}

	return true
}
