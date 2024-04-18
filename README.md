#### FROM Notion notes
# 1、Tips 

`Goland快捷键`

- CTRL+SHIFT+F，进行全局查找  如若无用查看是否与输入法快捷键冲突
- CTRL+R，替换文本
- CTRL+Backspace，按单词进行删除
- **SHIFT+ENTER，可以向下插入新行，即使光标在当前行的中间**
- CTRL+X，删除当前光标所在行
- **CTRL+D，复制当前光标所在行**

`栈和堆`

- 栈通常用于存储函数的局部变量、函数的调用信息（如函数调用的返回地址）、以及程序运行时所使用的临时数据
- 栈的内存分配是自动的，由编译器和运行时系统负责管理。当函数被调用时，会在栈上分配一块内存用于存储函数的局部变量和其他信息；当函数返回时，这块内存会被释放。
- 堆通常用于存储程序运行时动态创建的对象，如使用 **`new`** 或 **`make`** 关键字创建的对象。
- 堆上的内存分配由程序员自身负责管理，程序员需要手动分配和释放内存，以避免内存泄漏和其他内存管理问题。

`编译时断言`

- var _ HelloServiceInterface = (*HelloServiceClient)(nil)
- var _ HelloServiceInterface = &HelloServiceClient{}
    - 这两种写法都被称为编译时断言，它的作用是在编译时检查 **`HelloServiceClient`** 类型是否实现了 **`HelloServiceInterface`** 接口。如果 **`HelloServiceClient`** 没有实现 **`HelloServiceInterface`** 接口，编译时会产生错误，从而帮助开发人员在编译时捕获潜在的问题。

`切片`

- 将切片中的每个元素作为参数传递给函数

```go
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClienTokenAuth)))

	conn, err := grpc.NewClient("127.0.0.1:9090", opts...)
	
	// opts...是将一个切片 opts 展开成一个参数列表。在调用 grpc.NewClient 函数时，opts... 将会把 opts 中的每个元素作为参数传递给该函数。
  // 这种语法通常用于将一个切片中的元素作为函数的参数传递，以方便地将多个选项传递给函数。
```

# 2、包相关

`导包导入本地某函数`

在一个已经存在的Go项目中，如果新建了目录结构并且想要在main函数中导入新建的目录结构的函数，要确保新建的目录结构在GOPATH或者Go Modules的工作空间中（go.mod的module要与项目目录名对应上才能用 **import “项目名/目录名”** 的形式导入包）。

# 3、指针

`例子1`

```go
p := new(int)   // p, *int 类型, 指向匿名的 int 变量

// new(int) 创建了一个新的 int 变量，并返回了该变量的地址。
// 这个地址被赋值给了变量 p，因此 p 是一个指向该 int 变量的指针。
```

`例子2`

```go
func incr(p *int) int {
    *p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
    return *p
}

v := 1
incr(&v)              // side effect: v is now 2
fmt.Println(incr(&v)) // "3" (and v is 3)

// ① 函数的参数为什么是 p *int 而不是 *p int？
	在 Go 中，函数参数的类型放在变量名之后。因此，p *int 表示 p 是一个指向 int 类型变量的指针。
	这意味着 p 存储的是一个内存地址，该地址指向一个 int 类型的值。
// ② 函数为什么要用 *p++ 而不是 p++？
	在函数体内部，*p++ 结合了三个操作符：*、p 和 ++。*p 解引用指针，表示取得 p 指向的变量的值。
	++ 对变量的值进行增加操作。
	因此，*p++ 表示先获取 p 指向的变量的值，然后对该值进行自增操作。这里重要的是，p 本身并没有改变，只是增加了 p 指向的变量的值。
	
	
	如果在函数内部修改了指针 p 本身，意味着修改了指针所指向的内存地址，而不仅仅是修改了指针指向的变量的值。
	这可能会导致指针指向一个不同的内存地址，从而影响到其他指向同一内存地址的指针的值。在某些情况下，这样的修改可能会导致不可预料的结果或者内存访问错误。
// ③ 函数返回值为什么是 int 类型的，返回的是 *p？
	函数返回值为 int 类型，因为 *p 表示的是 p 指向的变量的值，类型为 int。
	因此，return *p 返回的是 p 指向的变量经过自增操作后的值。
// ④ 总结
	函数 incr 的目的是对传入的指针所指向的变量进行自增操作，并返回自增后的值。
	通过传入 &v，即 v 的地址，使得函数能够修改 v 的值，而不是仅仅传入 v 的副本。
```

# 4、并发

`make()创建chan`

- 在使用make()函数创建通道时，第二个参数表示通道的容量。通道有两种类型：有缓冲通道和无缓冲通道。
  **无缓冲通道：**当第二个参数为0时，创建的是无缓冲通道。无缓冲通道的特点是发送和接收操作是同步的，即发送操作会阻塞直到有其他goroutine接收该数据，接收操作也会阻塞直到有其他goroutine发送数据。
  **有缓冲通道：**当第二个参数大于0时，创建的是有缓冲通道。有缓冲通道允许在通道中存储一定数量的值，当通道满时发送操作会阻塞，直到有其他goroutine从通道中接收值，当通道空时接收操作会阻塞，直到有其他goroutine向通道发送值。

```go
// 例如：
// 创建一个无缓冲通道
ch := make(chan int)
// 创建一个容量为10的有缓冲通道
ch := make(chan int, 10)
```

- Goroutine 一般将其翻译为Go语言实现的协程，也就是说Golang在语言层面就实现了协程的支持

- runtime包




# 5、条件语句

`switch`

- switch不带操作对象时默认用true值代替，然后将每个case的表达式和true值进行比较

`if变量作用域`

```go
if f, err := os.Open(fname); err != nil { // compile error: unused: f
    return err
}
f.ReadByte() // compile error: undefined f
f.Close()    // compile error: undefined f
// 变量f的作用域只有在if语句内，因此后面的语句将无法引入它，这将导致编译错误。你可能会收到一个局部变量f没有声明的错误提示，具体错误信息依赖编译器的实现。

// 通常需要在if之前声明变量，这样可以确保后面的语句依然可以访问变量：
f, err := os.Open(fname)
if err != nil {
    return err
}
f.ReadByte()
f.Close()

// 你可能会考虑通过将ReadByte和Close移动到if的else块来解决这个问题：
if f, err := os.Open(fname); err != nil {
    return err
} else {
    // f and err are visible here too
    f.ReadByte()
    f.Close()
}
// 但这不是Go语言推荐的做法，Go语言的习惯是在if中处理错误然后直接返回，这样可以确保正常执行的语句不需要代码缩进。

// 下面例子中，虽然cwd在外部已经声明过，但是:=语句还是将cwd和err重新声明为新的局部变量。因为内部声明的cwd将屏蔽外部的声明，因此上面的代码并不会正确更新包级声明的cwd变量。
var cwd string

func init() {
    cwd, err := os.Getwd() // NOTE: wrong!
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory = %s", cwd)
}
```

# 6、defer与闭包函数

`闭包函数`

- 闭包函数是指在函数内部定义的函数，它可以访问并修改其外部函数的变量。闭包函数和普通函数的区别在于，闭包函数可以捕获和保持其所在函数的变量的状态，即使该变量在其所在函数已经返回后仍然可以使用。

    ```go
    /* 
    闭包函数的特征包括：
    	1、在函数内部定义了一个函数。
    	2、内部函数引用了外部函数的变量。
    	3、外部函数返回了内部函数
    */
    // 以下是一个简单的闭包函数的例子：
    func outerFunc() func() {
        message := "Hello, "
        innerFunc := func() {
            fmt.Println(message + "World!")
        }
        return innerFunc
    }
    
    func main() {
        myFunc := outerFunc()
        myFunc() // 输出：Hello, World!
    }
    ```


`defer调用其它函数闭包`

```go
/*
	当使用 defer 调用其他函数返回的闭包函数时，
	需要加上额外的括号来确保闭包函数被正确传递给 defer。
	这是因为 defer 需要一个函数调用作为参数，
	而闭包函数本身并不会被调用，只有在 defer 执行时才会调用闭包函数。
*/
func makeClosure(msg string) func() {
    return func() {
        fmt.Println(msg)
    }
}
// 要在 defer 中调用返回的闭包函数，需要像这样写：
func main() {
    defer makeClosure("Hello")() // 需要加上额外的括号
}
// 这样，makeClosure("Hello") 返回闭包函数 func() { fmt.Println("Hello") }，
// 然后 defer 将其延迟执行。
```