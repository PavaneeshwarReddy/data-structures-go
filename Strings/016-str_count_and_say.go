package strings

import (
	"strconv"
	"strings"
)

/*
Count and Say
	- Loop from 1, keep on counting
	- Keep a variable it it doesn't equals current add to string with count
*/

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	result := []int{}

	for i := 0; i < n-1; i++ {
		tempResult := result
		currVal := tempResult[0]
		count := 1

		for j := 1; j < len(result); j++ {
			if currVal != result[j] {
				currVal = result[j]
				count = 1
				tempResult = append(tempResult, count, currVal)
			} else {
				count += 1
			}
		}
		tempResult = append(tempResult, count, currVal)
	}

	var sb strings.Builder
	for _, val := range result {
		sb.WriteString(strconv.Itoa(val))
	}

	return sb.String()
}
