package strings

/*
Minimum Add to make parenthesis valid
A parentheses string is valid if and only if:

It is the empty string,
It can be written as AB (A concatenated with B), where A and B are valid strings, or
It can be written as (A), where A is a valid string.
You are given a parentheses string s. In one move, you can insert a parenthesis at any position of the string.

For example, if s = "()))", you can insert an opening parenthesis to be "(()))" or a closing parenthesis to be "())))".
Return the minimum number of moves required to make s valid.


- This this way, if we encounter an open that means its fine
- If we encounter closed bracket and open count is not zero then that means we can decrement open but if it's zero that means we need to increment closed

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
