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

	ret2 := ret(200)

	fmt.Println(ret2)
}
