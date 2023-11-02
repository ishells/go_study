// example 1
// package main

// import "sync"

// // 声明全局等待组变量
// var wg sync.WaitGroup

// func hello() {
// 	println("hello")
// 	wg.Done() // 告知当前goroutine完成
// }

// func main() {
// 	wg.Add(1) // 登记1个goroutine
// 	go hello()
// 	println("你好")
// 	wg.Wait() // 阻塞等待登记的goroutine完成
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup

// func hello(i int) {
// 	defer wg.Done() //goroutine结束就登记+1
// 	fmt.Println("hello", i)
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1) // 启动一个goroutine就登记+1
// 		go hello(i)
// 	}
// 	wg.Wait() // 等待所有登记的goroutine都结束
// }

// 多次执行上面的代码会发现每次终端上打印数字的顺序都不一致。这是因为10个 goroutine 是并发执行的，而 goroutine 的调度是随机的。

package main

import "fmt"

func recv(ch chan int) {
	// for {
	// 	v, ok := <-ch
	// 	if !ok {
	// 		fmt.Println("通道已经关闭")
	// 		break
	// 	} else {
	// 		fmt.Printf("v: %#v, ok: %#v \n", v, ok)
	// 	}
	// }
	// 可以改写为 for range接收值
	for v := range ch {
		fmt.Printf("v: %#v\n", v)
	}
}

func main() {
	channel := make(chan int, 2)
	channel <- 1
	channel <- 2
	close(channel)
	recv(channel)
}
