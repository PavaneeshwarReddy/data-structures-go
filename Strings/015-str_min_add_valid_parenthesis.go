package strings

/*
Minimum Add to make parenthesis valid
	- if you found ( we can increment open, if we found ) we can increment close
	- if open > 0 when we found ) then that means it's a vaid we can decrement open
*/

func MinAddParenthesis(s string) int {
	open := 0
	close := 0

	for _, val := range s {
		if val == '(' {
			open++
		} else {
			if open > 0 {
				open--
			} else {
				close++
			}
		}
	}

	return open + close
}
