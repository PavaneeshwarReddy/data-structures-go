package bitmanipulation

/*
We need to swap 2 numbers
- We can use the sum property
- We can use bitwise xor to get the value
*/

func SwapSum(a *int, b *int) {
	c := (*a) + (*b)
	(*a) = c - (*a)
	(*b) = c - (*b)
}

func SwapXor(a *int, b *int) {
	if *a == *b {
		return
	}
	(*a) = (*a) ^ (*b)
	(*b) = (*b) ^ (*a)
	(*a) = (*a) ^ (*b)
}
