package main

import (
	"SGG/chapter06/exercise01/functest"
	"fmt"
)

func main() {

	var n int = 0
	fmt.Scanln(&n)
	fmt.Println(functest.F(n))

}
