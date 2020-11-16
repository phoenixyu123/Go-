package main

//223断言assert

import "fmt"

type Point struct {
	x int
	y int
}

func main() {

	var a interface{} //空接口接收任何变量
	var point = Point{
		x: 1,
		y: 2,
	}
	a = point
	var b Point
	// b = a //error,不能把a当作Point赋值给Point
	b = a.(Point) //断言，看看能不能转换,看a是不是Point类型，是的话就转为b（Point）
	fmt.Println(b)

	//类型断言案例
	// var x interface{}
	// var b2 float32 = 1.1
	// x = b2 //空接口接收任意类型
	// //x=>float32
	// // y := x.(float64)//error,不可以进行float32->float64
	// y := x.(float32)
	// fmt.Printf("y的类型：%T,y:=%v\n", y, y) //float32,1.1

	//带检测的类型断言，不panic****
	var x interface{}
	var b2 float32 = 1.1
	x = b2
	if y, ok := x.(float64); ok { // 这种写法很常见
		fmt.Println("转换成功")
		fmt.Printf("y的类型：%T,y:=%v\n", y, y)
	} else {
		fmt.Println("转换失败") //断言失败给出提示继续执行，而不是结束运行
	}
	fmt.Println("继续执行")
}
