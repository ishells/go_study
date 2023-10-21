package main

import "fmt"

// Go语言中(目前版本1.17.1)是没有异常机制的，但是使用 panic/recover 模式来处理错误
// panic 可以在任何地方引发，但是 recover 只在 defer 调用的函数中有效。

// func funcA() {
// 	fmt.Println("func A")
// }

// func funcB() {
// 	panic("panic in B")
// }

// func funcC() {
// 	fmt.Println("func C")
// }
// func main() {
// 	funcA()
// 	funcB()
// 	funcC()
// }

// 以上程序输出：
/*
func A
panic: panic in B

goroutine 1 [running]:
main.funcB(...)
        .../code/go_study/10_panicRecover/panicRecover.go:13
main.main()
        .../code/go_study/10_panicRecover/panicRecover.go:21 +0xa5
exit status 2
*/

// 程序运行期间 funcB 中引发了 panic 导致程序崩溃异常退出了。
// 这个时候，就可以通过 recover 将程序恢复回来，继续向后执行

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	// defer定义了一个匿名立即执行函数
	defer func() {
		err := recover()
		//如果程序出现了panic错误,可以通过recover进行恢复
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}

/*
	注意：
	1、recover()必须搭配defer使用。
	2、defer一定要在可能引发panic的语句之前定义。
*/
