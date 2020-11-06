package main

import "fmt"

//display golang ptr type
func main() {
	//basic data type distribution in memory
	var i int = 10
	// how to get i 's address? -->&i
	fmt.Println("i 's address =", &i)
	// use i 's adr search the data in it
	fmt.Println(*&i)

	//init a ptr
	//var ptr *int = &i-->ptr is a point variable
	//-->ptr 's type is (*int)
	//-->ptr 's data is (&i)==ptr store i 's adr
	//==>ptr also has a adr !=i 's adr
	var ptr *int = &i
	fmt.Printf("ptr = %v\n", ptr)
	fmt.Printf("ptr 's adr = %v\n", &ptr)
	//how to get ptr-->i 's data
	//step1:find ptr-->adr==i 's adr
	//step2:use '*' find adr 's data==i->data
	fmt.Printf("ptr-->data = %v\n", *ptr) //use *
	// ----------------------------------------
	// -----------------test-------------------
	//1)get a num(int) 's adr and print'm
	//2)put num 's adr in ptr(pointer) and change
	//num 's data by using ptr
	var num int = 114514
	fmt.Printf("num 's adr = %v\n", &num)
	var ptr1 *int = &num
	//var ptr2 *float32 = &num-------< U should use same type
	//between ptr and num
	*ptr1 = 2 * *ptr1
	//[!]:dont use '*ptr=2*ptr',cuz 2*ptr==2*(i 's adr)
	fmt.Printf("num 's data = %v\n", num)

	// THINKING
	var a int = 100
	var b int = 200
	var ptr2 *int = &a
	*ptr2 = 300
	ptr2 = &b
	*ptr2 = 400
	fmt.Printf("a=%v  b=%v\n", a, b)

}
