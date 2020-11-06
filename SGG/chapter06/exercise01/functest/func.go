package functest

//F is a func
func F(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*F(n-1) + 1
	}
}
