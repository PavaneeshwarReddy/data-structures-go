package strings

import "math"

/*
Sum of Beauty of all Subtrings
	- Just traverse all substrings
	- Keep increment map for every substring starting from current character
	- Calculate min and max frequency differences
	- Sum all those results
*/

func diffMinMax(mp map[byte]int) int {
	min_val := math.MaxInt
	max_val := math.MinInt

	for _, value := range mp {
		min_val = min(min_val, value)
		max_val = max(max_val, value)
	}

	if min_val == math.MaxInt {
		min_val = 0
	}
	if max_val == math.MinInt {
		max_val = 0
	}

	return max_val - min_val
}

func SumOfBeautySubtrings(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		mp := make(map[byte]int)
		for j := i; j < len(s); j++ {
			mp[s[j]]++
			result += diffMinMax(mp)
		}
	}
	return result
}
