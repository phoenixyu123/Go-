// 假如还有97天放假，问：xx个星期零x天
package main

import "fmt"

func main() {
	var sum int = 97
	fmt.Printf("还有%d个星期零%d天\n", sum/7, sum%7)
}
