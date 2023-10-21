package main

import "fmt"

// (base int) 是函数的参数，（func(int) int, func(int) int）是两个函数类型的返回值
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		fmt.Println(base)
		return base
	}
	return add, sub
}

func main() {
	// f1实现的是calc函数内部的add函数及add函数的外部调用（base变量）；
	// 同理，f2实现的是calc函数内部的sub函数及sub函数的外部调用（base变量）
	f1, f2 := calc(10)
	// 此行先执行f1(1),此时使用的是共用变量base 10，随后执行完add函数之后base变为11，f2(2)执行的时候base值为11，即11-2=9
	fmt.Println(f1(1), f2(2)) //11 9
	// 此行同样先执行f1(3),此时base变量为9，执行add函数9+3=12，随后执行f2(4),即12-4=8
	fmt.Println(f1(3), f2(4)) //12 8
	// 此行先执行f1(5),此时base变量值为8，8+5=13，随后执行f2(6),即13-6=7
	fmt.Println(f1(5), f2(6)) //13 7
}
