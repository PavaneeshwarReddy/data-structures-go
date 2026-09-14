package greedy

/*
Cookie Distrubution
- cookie[i] >= greed[j] then only we can assign a cookie to a child
- If we sort and keep 2 pointer, if we are able to satisfy we decrement both or else only children
*/

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	i := len(g) - 1
	j := len(s) - 1
	res := 0

	for i >= 0 && j >= 0 {
		if g[i] <= s[j] {
			j--
			res++
		}
		i--
	}

	return res
}
