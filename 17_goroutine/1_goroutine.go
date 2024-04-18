// 本篇写的不全，详见   go并发：https://www.liwenzhou.com/posts/Go/concurrence/

// # 基本概念
/*
串行、并发与并行
串行：我们都是先读小学，小学毕业后再读初中，读完初中再读高中。
并发：同一时间段内执行多个任务（你在用微信和两个女朋友聊天）。
并行：同一时刻执行多个任务（你和你朋友都在用微信和女朋友聊天）。

进程、线程和协程
进程（process）：程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
线程（thread）：操作系统基于进程开启的轻量级进程，是操作系统调度执行的最小单位。
协程（coroutine）：非操作系统提供而是由用户自行创建和控制的用户态‘线程’，比线程更轻量级。
*/

/*
goroutine

Goroutine 是 Go 语言支持并发的核心，在一个Go程序中同时创建成百上千个goroutine是非常普遍的，
一个goroutine会以一个很小的栈开始其生命周期，一般只需要2KB。区别于操作系统线程由系统内核进行调度，
goroutine 是由Go运行时（runtime）负责调度。例如Go运行时会智能地将 m个goroutine 合理地分配给n个操作系统线程，实现类似m:n的调度机制，不再需要Go开发者自行在代码层面维护一个线程池。

Goroutine 是 Go 程序中最基本的并发执行单元。每一个 Go 程序都至少包含一个 goroutine——main goroutine，当 Go 程序启动时它会自动创建。

在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能——goroutine，
当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个 goroutine 去执行这个函数就可以了，就是这么简单粗暴。
*/

/*
go关键字
Go语言中使用 goroutine 非常简单，只需要在函数或方法调用前加上go关键字就可以创建一个 goroutine ，从而让该函数或方法在新创建的 goroutine 中执行。

go f()  // 创建一个新的 goroutine 运行函数f

匿名函数也支持使用go关键字创建 goroutine 去执行。
go func(){
  // ...
}()

一个 goroutine 必定对应一个函数/方法，可以创建多个 goroutine 去执行相同的函数/方法。
*/

/*
- runtime 包
> runtime.Gosched()
	让出CPU时间片,重新等待安排任务(大概意思就是本来计划的好好好的周末出去烧烤,但是你妈让你去相亲,
	两种情况第一就是你相亲速度非常快,见面就黄不耽误你继续烧烤,第二种情况就是你相亲速度特别慢,
	见面就是你依我依的,耽误了烧烤,但是还馋就是耽误了烧烤你还得去烧烤

package main

import (
	"fmt"
	"runtime"
	"time"
)

func longRunningTask() {
	for i := 0; i < 5; i++ {
		fmt.Println("Long running task - working...")
		time.Sleep(1 * time.Second)
		// 让出执行权限，让其他 goroutine 有机会执行
		runtime.Gosched()
	}
}

func shortTask() {
	for i := 0; i < 3; i++ {
		fmt.Println("Short task - working...")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// 创建一个长时间运行的 goroutine
	go longRunningTask()

	// 创建一个短时间运行的 goroutine
	go shortTask()

	// 主 goroutine 什么也不做，只是让其他 goroutine 运行
	for {
		runtime.Gosched()
	}
}


> runtime.Goexit()
	退出当前协程(一边烧烤一边相亲,突然发现相亲对象太丑影响烧烤,果断让她滚蛋,然后也就没有然后了)

> runtime.GOMAXPROCS
Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是
机器上的CPU核心数。例如在一个8核心的机器上,调度器会吧Go代码同时调度到8个OS线程上
(GOMAXPROCS是m:n调度中的n)。
Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。
Go1.5版本之前,默认使用的是单核心执行。Go1.5版本之后,默认使用全部的CPU逻辑核心数。
我们可以通过将任务分配到不同的CPU逻辑核心上实现并行的效果,这里举个例子:

两个任务只有一个逻辑核心,此时是做完一个任务再做另一个任务。
将逻辑核心数设为2,此时两个任务并行执行.

Go语言中的操作系统线程和goroutine的关系:
1.一个操作系统线程对应用户态多个goroutine。
2.go程序可以同时使用多个操作系统线程。
3.goroutine和OS线程是多对多的关系,即m:n。
*/

// channel
/*
单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

虽然可以使用共享内存进行数据交换，但是共享内存在不同的 goroutine 中容易发生竞态问题。为了保证数据交换的正确性，很多并发模型中必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。

Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
*/

// chanel类型
/// chanel是Go语言中的一种特有的类型，声明通道类型变量的格式如下：
/*
	var 变量名称 chan 元素类型
	其中：
		chan：是关键字
		元素类型：是指通道中传递元素的类型
	举几个例子：
	var ch1 chan int   // 声明一个传递整型的通道
	var ch2 chan bool  // 声明一个传递布尔型的通道
	var ch3 chan []int // 声明一个传递int切片的通道

	channel零值
	未初始化的通道类型变量其默认零值是nil。

	初始化channel
	声明的通道类型变量需要使用内置的make函数初始化之后才能使用。具体格式如下：
	make(chan 元素类型, [缓冲大小])
	其中,channel的缓冲大小是可选的


	举几个例子：
	ch4 := make(chan int)
	ch5 := make(chan bool, 1)  // 声明一个缓冲区大小为1的通道

	初始化channel
	make(chan 元素类型, [缓冲大小]),    channel的缓冲大小是可选的。

	channel操作
	通道共有发送（send）、接收(receive）和关闭（close）三种操作。而发送和接收操作都使用<-符号。
	**注意：**一个通道值是可以被垃圾回收掉的。通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

	关闭后的通道有以下特点：
		对一个关闭的通道再发送值就会导致 panic。
		对一个关闭的通道进行接收会一直获取值直到通道为空。
		对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
		关闭一个已经关闭的通道会导致 panic。

	无缓冲的通道
	无缓冲的通道无缓冲的通道又称为阻塞的通道。
	无缓冲的通道只有在有人接收值的时候才能发送值。就像你住的小区没有快递柜和代收点,
	快递员给你打电话必须要把这个物品送到你的手中,简单来说就是无缓冲的通道必须有接收才能发送。
	无缓冲通道上的发送操作会阻塞,直到另一个goroutine在该通道上执行接收操作,这时值才能发送成功,两个goroutine将继续执行。
	相反,如果接收操作先执行,接收方的gooroutine将阻塞,直到另一个goroutine在该通道上发送一个值。
	使用无缓冲通道进行通信将导致发送和接收的goroutine同同步化。因此,无缓冲通道也被称为同步通道。


	有缓冲的通道
	只要通道的容量大于零，那么该通道就属于有缓冲的通道
	通道的容量表示通道中能存放元素的数量。就像你小区的快递柜只有那么个多格子,格子满了就装不下了,就阻塞了,
	等到别人取走一个快递员就能往里面放

	判断通道值是否被取完了？
		1、多返回值模式
		当向通道中发送完数据时，我们可以通过close函数来关闭通道。当一个通道被关闭后，再往该通道发送值会引发panic，从该通道取值的操作会先取完通道中的值。通道内的值被接收完后再对通道执行接收操作得到的值会一直都是对应元素类型的零值。那我们如何判断一个通道是否被关闭了呢？
		value, ok := <- ch
		其中：
		value：从通道中取出的值，如果通道被关闭则返回对应类型的零值。
		ok：通道ch关闭时返回 false，否则返回 true。

		2、for range接收值
		通常我们会选择使用for range循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环。

	单向通道
	Go语言中提供了单向通道来处理这种需要限制通道只能进行某种操作的情况
	<-chan int // 只接收通道，只能接收不能发送
	chan<- int // 只发送通道，只能发送不能接收
	其中，箭头<-和关键字chan的相对位置表明了当前通道允许的操作，这种限制将在编译阶段进行检测。另外对一个只接收通道执行close也是不允许的，因为默认通道的关闭操作应该由发送方来完成。
	在函数传参及任何赋值操作中将双向通道转换为单向通道是可以的，但反过来是不可以的

	select多路复用
		Go 语言内置了select关键字，使用它可以同时响应多个通道的操作。
		Select 的使用方式类似于之前学到的 switch 语句，它也有一系列 case 分支和一个默认的分支 default。
		每个 case 分支会对应一个通道的通信（接收或发送）过程。select 会一直等待，直到其中的某个 case 的通信操作完成时，就会执行该 case 分支对应的语句。具体格式如下：
		select {
		case <-ch1:
			//...
		case data := <-ch2:
			//...
		case ch3 <- 10:
			//...
		default:
			//默认操作
		}

		Select 语句具有以下特点。
		① 可处理一个或多个 channel 的发送/接收操作。
		② 如果多个 case 的 channel 同时满足，select 会随机选择一个执行。
		③ 对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。

		package main

		import "fmt"

		func main() {
			ch := make(chan int, 1)
			for i := 1; i <= 10; i++ {
				select {
				case x := <-ch:
					fmt.Println(x)
				case ch <- i:
				}
			}
		}
		上面的代码输出内容如下。
		1
		3
		5
		7
		9

		示例中的代码首先是创建了一个缓冲区大小为1的通道 ch，进入 for 循环后：
		第一次循环时 i = 1，select 语句中包含两个 case 分支，此时由于通道中没有值可以接收，所以x := <-ch 这个 case 分支不满足，而ch <- i这个分支可以执行，会把1发送到通道中，结束本次 for 循环；
		第二次 for 循环时，i = 2，由于通道缓冲区已满，所以ch <- i这个分支不满足，而x := <-ch这个分支可以执行，从通道接收值1并赋值给变量 x ，所以会在终端打印出 1；
		后续的 for 循环以此类推会依次打印出3、5、7、9。

		select还可用于判断管道是否存满：
		package main

		import (
			"fmt"
		)

		func main() {
			// 创建一个带缓冲区大小为 2 的整数管道
			ch := make(chan int, 2)

			// 向管道中发送两个整数
			ch <- 1
			ch <- 2

			// 使用 select 语句判断管道是否已满
			select {
			case ch <- 3:
				fmt.Println("Sent a value")
			default:
				fmt.Println("Channel is full")
			}

			// 从管道中接收并打印所有值
			for len(ch) > 0 {
				fmt.Println("Received:", <-ch)
			}

			// 使用 select 语句判断管道是否为空
			select {
			case x := <-ch:
				fmt.Println("Received value:", x)
			default:
				fmt.Println("Channel is empty")
			}
			在这个示例中，我们创建了一个带有缓冲区大小为 2 的整数管道 ch，并向其中发送了两个整数。
            然后，我们使用 select 语句尝试向管道中发送第三个整数。由于管道已满，select 语句会执行 default 分支，打印 "Channel is full"。
			接着，我们使用 for 循环从管道中接收并打印所有值。
			最后，我们再次使用 select 语句尝试从空管道中接收值，由于管道为空，select 语句会执行 default 分支，打印 "Channel is empty"。

  // 并发安全与锁 sync
	有时候在Go代码中可能会存在多个goroutine同时操作一个资源临界区),这种情况会发生竞态问题(数据竞态)。
	类比现实生活中的例子有十字路口被各个方向的的汽气车竞争;还有火车上的卫生间被车厢里的人竞争。

	// sync.WaitGroup
		在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。 sync.WaitGroup有以下几个方法：
		方法名	功能
		func (wg * WaitGroup) Add(delta int)	计数器+delta
		(wg *WaitGroup) Done()	计数器-1
		(wg *WaitGroup) Wait()	阻塞直到计数器变为0

		sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了 N 个并发任务时，就将计数器值增加N。
		每个任务完成时通过调用 Done 方法将计数器减1。通过调用 Wait 来等待并发任务执行完，当计数器值为 0 时，表示所有并发任务已经完成。

	// 互斥锁
  	互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源。
	Go 语言中使用sync包中提供的Mutex类型来实现互斥锁。

  	sync.Mutex提供了两个方法供我们使用
	func (m *Mutex) Lock()	获取互斥锁
	func (m *Mutex) Unlock()	释放互斥锁

	使用互斥锁能够保证同一时间有且只有一个 goroutine 进入临界区，其他的 goroutine 则在等待锁；
	当互斥锁释放后，等待的 goroutine 才可以获取锁进入临界区，多个 goroutine 同时等待一个锁时，唤醒的策略是随机的。

  // 读写互斥锁
  	互斥锁是完全互斥的，但是实际上有很多场景是读多写少的，
	当我们并发的去读取一个资源而不涉及资源修改的时候是没有必要加互斥锁的，
	这种场景下使用读写锁是更好的一种选择。读写锁在 Go 语言中使用sync包中的RWMutex类型。

	读写锁分为两种：读锁和写锁。
	当一个 goroutine 获取到读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
	而当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待。

	sync.RWMutex提供了以下5个方法。

	方法名	功能
	func (rw *RWMutex) Lock()		加写锁
	func (rw *RWMutex) Unlock()		释放写锁
	func (rw *RWMutex) RLock()		加读锁
	func (rw *RWMutex) RUnlock()	释放读锁
	func (rw *RWMutex) RLocker() Locker	返回一个实现Locker接口的读写锁

	需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。
*/

/*
	// 原子操作 atomic包
	代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。
	针对 基本数据类型 我们还可以使用原子操作来保证并发安全,因为原子操作是Go语言提供的方法，
	它在用户态就可以完成,因此性能比加锁操作更好。Go语言中原子操作由内置的标准库sync/atomic提供。

	atomic包提供了底层的原子级内存操作,对于同步算法的实现很有用。
	这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用,使用通道或者sync包的函数/类型实现同步更好


*/

// package main

// import "fmt"

// func main() {
// 	fmt.Println("1111")
// 	fmt.Println("2222")

// }

// package main

// import "fmt"

// type Books struct {
// 	title   string
// 	author  string
// 	subject string
// 	book_id int
// }

// func main() {
// 	var Book1 Books /* 声明 Book1 为 Books 类型 */
// 	var Book2 Books /* 声明 Book2 为 Books 类型 */

// 	var Books1_struct_pointer *Books
// 	Books1_struct_pointer = &Book1

// 	/* book 1 描述 */
// 	Book1.title = "Go 语言"
// 	Book1.author = "www.runoob.com"
// 	Book1.subject = "Go 语言教程"
// 	Book1.book_id = 6495407

// 	/* book 2 描述 */
// 	Book2.title = "Python 教程"
// 	Book2.author = "www.runoob.com"
// 	Book2.subject = "Python 语言教程"
// 	Book2.book_id = 6495700

// 	/* 打印 Book1 信息 */
// 	printBook(&Book1)

// 	/* 打印 Book2 信息 */
// 	printBook(&Book2)

// 	print("使用结构体指针来访问结构体成员：")
// 	fmt.Println(Books1_struct_pointer.author)
// }
// func printBook(book *Books) {
// 	fmt.Printf("Book title : %s\n", book.title)
// 	fmt.Printf("Book author : %s\n", book.author)
// 	fmt.Printf("Book subject : %s\n", book.subject)
// 	fmt.Printf("Book book_id : %d\n", book.book_id)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say("world")
// 	say("hello")
// }

// package main

// import "fmt"

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum // 把 sum 发送到通道 c
// }

// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}

// 	fmt.Println(s[:len(s)/2])
// 	fmt.Println(s[len(s)/2:])

// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c // 从通道 c 中接收

// 	fmt.Println(x, y, x+y)
// }

// package main

// import (
// 	"fmt"
// )

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	// 如果range结束之后不关闭通道就会发生阻塞报错
// 	// fatal error: all goroutines are asleep - deadlock!
// 	// goroutine 1 [chan receive]
// 	close(c)
// }

// func main() {
// 	c := make(chan int, 10)
// 	go fibonacci(cap(c), c)
// 	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
// 	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
// 	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
// 	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

// package main

// import "fmt"

// func f1(name string) (age int) {
// 	fmt.Println("hello", name)
// 	return
// }

// // 函数形参无需定义名字，比如此处
// // f2函数的参数是一个函数f，而定义该形参时，其参数和返回值不用指定名字，只写明类型即可
// func f2(f func(string) int, name string) {
// 	f(name)
// }

// func main() {
// 	// submitTime := int64(-1)
// 	// print(submitTime)
// 	var s string
// 	fmt.Scan(&s)
// 	fmt.Println(s)
// 	f2(f1, "zhi")
// }

package main

import "fmt"

// n个台阶，一次可以走1步，也可以走2步，有多少种走法
/*
	解法1:
	将大问题拆解为小问题，因为问题是只能走1步或2步，所以将大问题拆解成小问题的拆解条件就是n-1和n-2
	递归调用，当台阶数>=1时，不管一次走1步还是一次走2步，
	最后剩余的台阶数不是1就是2，这个就可以作为结束条件
	taijie(4)
	||
	taijie(3) + taijie(2)
	||
	taijie(2) + taijie(1) + 2
	||
	1 + 2 + 2
*/
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	return taijie(n-1) + taijie(n-2)
}

func main() {
	fmt.Println(taijie(6))
}

// 台阶问题解法2：
/*
 找到规律的前提下，你会发现n=1、2、3、4、5、6时，走法分别是1、2、3、5、8、13
 即后一个数等于前两个数的和，这之后就可以根据这个规律写出该问题的解法
 该解法叫动态规划，相比于递归调用复杂度要低，效率要高一些
*/
