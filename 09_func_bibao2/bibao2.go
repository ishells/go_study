package main

import "fmt"

func f1(f func()) {
	fmt.Println("This is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("This is f2")
	fmt.Println(x + y)
}

// 要求：
// 将f2作为f1的参数传入
// (使用闭包)

// 由上边的f1、f2函数可以看出，f1的参数需要是一个无参数、无返回值的函数
// 而f2函数本身有两个参数 x、y,所以如果想要将f2函数作为f1函数的参数，那就要用到闭包的形式

func f3(f func(int, int), x, y int) func() {
	// 匿名函数使用外部的变量x、y
	tmp := func() {
		// 调用参数
		f(x, y)
	}
	return tmp
}

func main() {
	// ret变量接收的是f3函数的返回值，也就是f3函数内的匿名函数tmp
	// 把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	ret := f3(f2, 100, 200)
	// f1(ret)，传入参数是ret变量，ret变量也就是f3函数的返回值tmp
	// 在f3内部的匿名函数tmp中，调用的正是传入的第一个参数，又因为f3的第一个参数正是f2函数
	// 所以这里也就是，f1调用f3，f3调用f2，最终相当于f1调用了参数类型不一致的f2
	// 因为f1和f3参数一致，所以f1()可以调用f3()，因为f3定义的参数类型与f2类型一致，所以f2可以作为参数传入f3
	f1(ret)
}
