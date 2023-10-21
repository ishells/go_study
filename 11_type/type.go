package main

import "fmt"

// 自定义类型 和 类型别名
// 自定义类型在代码编译完成之后类型保留，类型别名只在代码编写过程中有效，编译后就没有了

// type后面跟的是类型
type myInt int     // 自定义类型
type yourInt = int // 类型别名

func main() {
	var my myInt
	my = 100
	fmt.Println(my)
	fmt.Printf("%T\n", my)

	var your yourInt
	your = 100
	fmt.Println(your)
	fmt.Printf("%T\n", your)

}
