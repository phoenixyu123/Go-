package main

func AddUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res //true:55 else:false
}

func GetSub(n1 int, n2 int) int {
	return n1 - n2
}
