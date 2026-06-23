package strings

/*
Largest Odd Number
	- if from the last if it's an odd then entire string from there is odd
*/

func LargestOddNumber(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i]-'0')%2 == 1 {
			return s[:i+1]
		}
	}

	return ""
}
