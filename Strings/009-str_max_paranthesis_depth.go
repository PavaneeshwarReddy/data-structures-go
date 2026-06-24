package strings

/*
Max Paranthesis Depth
	- Just increment count if you encounter ( else decrement if )
	- Store max length
*/

func MaxParanthesisDepth(s string) int {
	maxLen := 0
	currLen := 0
	for _, v := range s {
		if v == '(' {
			currLen++
		}
		if v == ')' {
			currLen++
		}
		maxLen = max(maxLen, currLen)
	}

	return maxLen
}
