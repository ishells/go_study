package main

import "fmt"

func main() {

	fmt.Println("hello world")
	// 1、if条件语句

	// Go语言规定与if匹配的左括号{必须与if和表达式放在同一行，{放在其他位置会触发编译错误。
	// 同理，与else匹配的{也必须与else写在同一行，else也必须与上一个if或else if右边的大括号在同一行。

	/*
		if 表达式1{
			分支1
		}else if 表达式2{
			分支2
		}else{
			分支3
		}
	*/
	/*
		在 if 表达式之前可以添加一个执行语句，再根据变量值进行判断``
		if score := 65; score >= 90 {
			fmt.Println("A")
		} else if score > 75 {
			fmt.Println("B")
		} else {
			fmt.Println("C")
		}
	*/

	// 2、go语言中只有for循环，没有其它的循环语句

	/*
		for 初始语句;条件表达;结束语句{
			循环体语句
		}
	*/

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 2.1 for循环的初始语句可以被忽略，但是初始语句后的分号必须要写
	/*
		i := 0
		for ; i < 10; i++ {
			fmt.Println(i)
		}
	*/

	// 2.2 for循环的初始语句和结束语句可以都省略
	/*
		i := 0
		for i < 10 {
			fmt.Println(i)
			i++
		}
	*/

	// 2.3 for range循环（键值循环）
	s := "helloworld世界"
	// i 表示索引， v 表示值
	fmt.Println("从字符串中拿出索引和字符：")
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	// 从字符串中拿出具体的字符
	fmt.Println("从字符串中拿出具体的字符: ")
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	// 2.4 continue、break跳出循环
	/*
		break 跳出循环
		continue	开始下一次循环

	*/

	// 2.5 使用goto语句简化代码
	/*
		goto语句通过标签进行代码间的无条件跳转。
		goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
		Go语言中使用goto语句能简化一些代码的实现过程。
		例如双层嵌套的for循环要退出时
	*/
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")

	// 3、switch canume
	// 当已给定条件的case都不满足的时候，默认会执行 default，
	// Go语言规定每个switch只能有一个default分支

	for num := 0; num < 5; num++ {
		switch num {
		case 1:
			fmt.Println("num is: 1")
		case 2:
			fmt.Println("num is: 2")
		case 3:
			fmt.Println("num is: 3")
		case 4:
			fmt.Println("num is: 4")
		default:
			fmt.Println("无效的输入！")
		}
	}

	// 一个分支可以有多个值，多个case值中间使用英文逗号进行分割
	switch num1 := 7; num1 {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(num1)
	}

	// 分支还能使用表达式，这时候switch后面不需要在跟判断变量
	age := 30
	switch {
	case age < 25:
		fmt.Println("年龄小于25")
	case age > 25 && age < 35:
		fmt.Println("年龄大于25小于35")
	case age > 60:
		fmt.Println("年龄大于60")
	default:
		fmt.Println("年轻真好")
	}

	// fallthrough 语法可以主动执行该case的下一个case
	// 下面例子当匹配到a的case时，顺带把b case也执行了！
	stringsss1 := "a"
	switch {
	case stringsss1 == "a":
		fmt.Println("a")
		fallthrough
	case stringsss1 == "b":
		fmt.Println("b")
	case stringsss1 == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

}
