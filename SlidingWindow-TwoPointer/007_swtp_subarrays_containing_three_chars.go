package slidingwindowtwopointer

/*
Subarrays containing three characters
- This may not seem like sliding window and shows more characteristics as sliding window
- When we found all characters then that means all these belong some sub array now min of these three +1 is the total subarrays  possible
*/

func numberOfSubstrings(s string) int {
	mp := make(map[byte]int)
	mp['a'] = -1
	mp['b'] = -1
	mp['c'] = -1
	res := 0
	for i := range s {
		mp[s[i]] = i
		if mp['a'] != -1 && mp['b'] != -1 && mp['c'] != -1 {
			mi := min(mp['a'], min(mp['b'], mp['c']))
			res += mi + 1
		}
	}
	return res
}
