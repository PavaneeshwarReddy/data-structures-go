package bitmanipulation

/*
Ith bit set or not
- We can simply move to that particular bit by right shifting the bits and make an AND with it
*/

func SetOrNot(n int, i int) bool {
	if (n>>i)&1 == 1 {
		return true
	}
	return false
}
