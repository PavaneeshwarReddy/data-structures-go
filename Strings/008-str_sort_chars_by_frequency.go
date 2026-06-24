package strings

import "sort"

/*
Sort Chars by frequency
	- Take a map record the frequency of individual character
	- Sort them and add to result in that frequency
*/

type Info struct {
	char rune
	freq int
}

func SortStringByCharFrequency(s string) string {
	mp := make(map[rune]int)
	for _, val := range s {
		mp[val]++
	}

	count := []Info{}
	for key, val := range mp {
		count = append(count, Info{char: key, freq: val})
	}
	sort.Slice(count, func(i int, j int) bool {
		return count[i].freq > count[j].freq
	})

	result := make([]byte, 0, len(s))
	for _, val := range count {
		for i := 0; i < val.freq; i++ {
			result = append(result, byte(val.char))
		}
	}

	return string(result)

}
