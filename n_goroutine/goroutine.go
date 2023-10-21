// go并发：https://www.liwenzhou.com/posts/Go/concurrence/

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
