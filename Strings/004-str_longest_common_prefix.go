package strings

/*
Longest Common Prefix
	- Find the min length of all strings
	- Traverse atmost minLength and verify in all strings, if not matched return result
*/

func LongestCommonPrefix(s []string) string {
	minLength := 10000
	for _, v := range s {
		minLength = min(minLength, len(v))
	}

	result := ""
	for i := 0; i < minLength; i++ {
		common := ""
		for j := 0; j < len(s); j++ {
			if common == "" {
				common = string(s[i][j])
			} else {
				if common != string(s[i][j]) {
					return result
				}
			}
		}
	}

	return ""
}
