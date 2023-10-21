package main

import "fmt"

type Base struct {
	name string
}

func (b *Base) m1() int {
	return 666
}

type Son struct {
	// 如果结构体之间存在匿名嵌套关系，则 子结构体 可以继承 父结构体中的方法
	// 匿名的方式，如果改成 base Base ,则无法直接继承父结构体中的方法，不过可以间接调用父结构体中的方法
	// Son结构体.base属性.m1()也是可以完成调用父结构体中的方法的
	//Base Base
	Base
	age int
}

func (s *Son) m2() int {
	return 888
}

func main() {
	son := Son{age: 18, Base: Base{name: "son"}}
	//fmt.Println(son.Base.m1())
	fmt.Println(son.m1())
	fmt.Println(son.m2())

}
