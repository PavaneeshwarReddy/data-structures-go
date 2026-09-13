package slidingwindowtwopointer

/*
Fruits in a Basket
- The pattern stays like, there is a fixed contraint that is 2 baskets and should contain only unique fruits
- Well this is the pattern similar to longest substring with non repeating characters but now with some basket size
*/

func checkValidScenario(baskets map[int]int) bool {
	return len(baskets) <= 2
}

func totalFruit(fruits []int) int {
	baskets := make(map[int]int)
	start, end, result := 0, 0, 0

	for end < len(fruits) {
		baskets[fruits[end]] += 1

		for !checkValidScenario(baskets) {
			baskets[fruits[start]] -= 1
			if baskets[fruits[start]] == 0 {
				delete(baskets, fruits[start])
			}
			start++
		}

		result = max(end-start+1, result)
		end++
	}

	return result

}
