package bitmanipulation

/*
1s complement:
- Take all the bits and flip them

2s Complement:
- Apply 1s complement
- Add 1
*/

func OnesComplement(n int) int {
	return ^n
}

func TwosComplement(n int) int {
	return -n
}
