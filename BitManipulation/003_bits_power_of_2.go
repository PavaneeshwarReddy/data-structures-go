package bitmanipulation

/*
Power of two
- If only one bit is set for entire bits then it will power of 2

- If you thing all power of 2 has this configuration that 1 bit will be set and other bits will be unset
4 - 100
3 - 011
So we do an & operation then all bits will be zero
*/

func PowerOf2Check(n int) bool {

	if n > 0 && (n&(n-1)) == 0 {
		return true
	}

	return false

}
