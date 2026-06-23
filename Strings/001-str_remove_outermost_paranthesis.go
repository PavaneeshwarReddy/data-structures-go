package strings

/*
Remove Outermost Paranthesis
	- If  '(' is found then increment count, if count > 0 then only include
	- If  ')' is found then decrement count, if count > 0 then only include
*/

func RemoveOutermostPranthesis(s string) string {
	result := ""
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if count > 0 {
				result += string(s[i])
			}
			count += 1
		} else {
			count -= 1
			if count > 0 {
				result += string(s[i])
			}
		}
	}

	return result
}
