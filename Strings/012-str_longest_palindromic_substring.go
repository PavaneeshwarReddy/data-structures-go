package strings

/*
Longest Palindromic Substring
	- Traverse through all the subtrings
	- Check whether that substring is a valid palindrom or not
	- If it's a plaindrom then calculate length , if greater than max store as result
*/

func checkPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func LongestPalindromicSubstring(s string) string {
	maxLen := 0
	result := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if checkPalindrome(s[i : j+1]) {
				if maxLen < (j + 1 - i) {
					maxLen = j + 1 - i
					result = s[i : j+1]
				}
			}
		}
	}
	return result
}
