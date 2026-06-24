package strings

import (
	"strconv"
	"unicode"
)

/*
String to int (atoi)
	- unicode : IsDigit, IsNumber, IsSpace , ....
*/

func StringToInt(s string) int {
	result := ""
	negative := false
	for _, val := range s {
		if val != ' ' {
			if val == '-' || val == '+' {
				if result != "" {
					break
				}
				result += string(val)
			} else if unicode.IsDigit(val) {
				result += string(val)
			} else {
				break
			}
		} else {
			if result != "" {
				break
			}
		}
	}

	if negative {
		result = "-" + result
	}

	if result == "" {
		return 0
	} else {
		val, _ := strconv.ParseInt(result, 10, 32)
		return int(val)
	}
}
