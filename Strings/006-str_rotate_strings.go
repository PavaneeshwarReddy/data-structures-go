package strings

import "strings"

/*
Rotate Strings
	- If we consider after rotation some part stays at last and some goes at the start
	- If we club the s + s , then this should contain somehow the pattern we are searching for
	- For this new string if our goal is present as a substring then we can return true else false
*/

func RotateStrings(src string, goal string) bool {
	if len(src) != len(goal) {
		return false
	}

	newStr := src + src

	return strings.Contains(newStr, goal)
}
