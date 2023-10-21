package main

import "fmt"

// 闭包指的是一个函数和与其他相关的引用环境组合而成的实体。
// （闭包是一个函数，这个函数包含了它外部作用域的一个变量）

/*
	闭包底层原理：
	1、函数可以作为return返回值
	2、函数内部找变量的顺序，先在自己内部找，找不到在往外部找
*/

// 简单来说，闭包 = 函数 + 外部引用环境。
// 例子：

func adder(x int) func(int) int {
	// var x = 100
	return func(y int) int {
		println(y)
		println(x)
		x += y
		println(x)
		return x
	}
}

func main() {
	// 变量 ret 是一个函数（用以接收adder函数的返回值，即匿名函数）
	// 并且它引用了其外部作用域中的x变量，此时ret就是一个闭包，在ret的生命周期内，变量x也一直有效
	ret := adder(100)
	// 相当于
	/*
		ret = func (y int) int{
			println(y)
			println(x)
			x = 100 + y
			print(x)
			return x
		}
	*/
	// ret2 调用ret时就是为形参y值传入一个实参200就是y值
	ret2 := ret(200)

	fmt.Println(ret2)
}

/*
	目前理解的闭包操作即定义一个函数的返回值为函数，来达到调用外部环境参数的目的：
		1、首先定义一个函数A，函数A本身需要参数，并通过函数返回值定义一个匿名函数接受外部参数
		2、然后定义变量向函数A传不同实参值，并接收函数A的返回值，即预先定义的匿名函数（实现函数A）
		3、通过向第二步定义的不同变量（对应不同参数的同一个匿名函数）传入不同参数达到不同的目的（调用外部环境参数）

	示例如下（自己写的例子）：
	//package main
	//
	//import "fmt"
	//
	//func stringSum(nameA string) func(nameB string) string {
	//	return func(nameB string) string {
	//		return nameA + nameB
	//	}
	//}
	//
	//func main() {
	//	nameA := stringSum("张")
	//	nameB := stringSum("李")
	//
	//	fmt.Println(nameA("三"))
	//	fmt.Println(nameB("四"))
	//
	//}
*/
