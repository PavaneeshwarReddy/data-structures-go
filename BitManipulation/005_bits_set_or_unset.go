package bitmanipulation

/*
Set or unset the last bit
- If we do xor of 2 numbers which is 1 ^ num then we always fip the last bit
*/

func SetOrUnset(num int) int {
	return num ^ 1
}
