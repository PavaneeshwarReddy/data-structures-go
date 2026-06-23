package strings

import "strings"

/*
Reverse Words
	- Split the string into string array
	- Iterate in reverse and append to the result
*/

func ReverseWords(s string) string {
	splitStr := strings.Split(s, " ")
	result := ""

	for i := len(splitStr) - 1; i >= 0; i-- {
		if result != "" {
			result += " "
		}
		result += string(splitStr[i])
	}
	return result
}
