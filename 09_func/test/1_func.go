package main

import "fmt"

// 1、函数的定义
/*
	func 函数名(参数)(返回值){
		函数体
	}

	go语言是强引用类型，任何时候声明变量都是指定数据类型
	这里的参数、返回值也都必须指定数据类型

	函数名：由字母、数字、下划线组成。但函数名的开头不能是数字。在同一个包内，函数名称也不能重名（包的概念详见后文）。
	参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
	返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
	函数体：实现指定功能的代码块

*/

// 2、函数的参数
// Go语言中函数传递的都是值类型，修改之后并不会影响初始变量的值
// 函数传参相当于复制粘贴操作，在函数内部修改的相当于原始变量的副本值，即不会改变原始值

func sum(x int, y int) (ret int) {
	return x + y
}

// 函数的参数和返回值都是可选的，例如可以实现一个既没有参数也没有返回值的函数
func sayHello() {
	fmt.Println("Hello World")
}

// 参数类型简写
// 函数的参数中如果相邻变量的类型相同，则可以省略数据类型
// intSum函数有两个参数，这两个参数的类型均为int，因此可以省略x的类型，因为y后面有类型说明，x参数也是该类型
func intSum(x, y int) int {
	return x + y
}

// 可变参数
/*

	可变参数格式为 "...type",

		如果使用 "...int"，"...string",可变参数的数据类型是被限制死的，比如可变参数只能是int类型，string类型
		如果想要使用任意类型的可变参数，可以指定类型为 "interface{}"，格式固定为"...interface{}"

	可变参数是指函数的参数变量不固定。Go语言中的可变参数通过在参数名后加  ...  来标识
	注意：可变参数通常要作为函数的最后一个参数


*/
func intSum2(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 固定参数搭配可变参数使用时，可变参数要放到固定参数的后面，示例代码如下：
// 本质上，函数的可变参数是通过切片来实现的
func intSum3(x int, y ...int) int {
	fmt.Print("sum3函数分别输出可变参数值:")
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	fmt.Print("sum3函数输出可变参数和:")
	return sum
}

// 3、函数返回值
// Go语言中函数支持多返回值，函数如果有多个返回值时，定义返回参数必须用()将所有返回值包裹起来
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

/*
函数可以有0或多个返回值，返回值需要指定数据类型，返回值通过return关键字来指定。

1、return可以有参数，也可以没有参数，这些返回值可以有名称，也可以没有名称。go中的函数可以有多个返回值。
2、return关键字中指定了参数时，返回值可以不用名称。如果return省略参数，则返回值部分必须带名称
3、当返回值有名称时，必须使用括号包围，逗号分隔，即使只有一个返回值
4、即使返回值命名了，return中也可以强制指定其它返回值的名称，也就是说return的优先级更高
5、命名的返回值是预先声明好的，在函数内部可以直接使用，无需再次声明。命名返回值的名称不能和函数参数名称相同，否则报错提示变量重复定义
6、return中可以有表达式，但不能出现赋值表达式，这和其它语言可能有所不同。例如return a+b是正确的，但return c=a+b是错误的。
————————————————
原文链接：https://blog.csdn.net/guolianggsta/article/details/124279421

*/

// 不带命名返回值
// 如果函数的返回值是无名的（不带命名返回值），则go语言会在执行return的时候会执行一个类似创建一个临时变量作为保存return值的动作。

// 返回值命名：
// 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过 return 关键字返回
// 例如：
func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 返回值补充，当函数返回值类型为slice时，nil可以看作是一个有效的slice，没必要显示返回一个长度为0的切片
func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	} else {
		fmt.Println("helloworld")
	}
	return nil
}

func printName(name string) {
	fmt.Println("hello", name)
}

// 4、defer语句
/*
	defer语句会将其后所跟随的语句进行延迟处理。
	defer所在的函数即将return时，将之前延迟处理的语句按照defer定义的逆序来执行（也就是说 defer定义的语句 先定义后执行）

	return返回值的运行机制：return并非原子操作，共分为(1) 返回值赋值；(2) RET指令两步操作。而defer语句执行在赋值之后，RET之前。
*/

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

// defer语句执行的时机
/*
	Go语言中return语句（不是原子操作）底层实现：
		return X：    返回值=X  ->  RET指令
		第一步：返回值赋值
		第二步：真正的RET返回

	defer语句执行时机：
	  	return X：    返回值=X	->	逆序运行defer	->	RET指令
		第一步：return最先执行，先将结果写入返回值中（即赋值）
		第二步：按照"先定义后执行"的顺序延迟执行多个defer语句
		第三步：真正的RET返回
*/

/*
f1函数首先仅声明了返回值类型为int，并没有指定返回值是谁
函数体首先初始化x值为5，随后+1=6，
当函数结束，执行return命令时（分三步），
第一步因为return优先级较高，所以执行return，指定返回值为x（赋值"返回值"为此时x的大小，即x=6），
第二步执行defer的立即执行函数，x++，加完之后x值为7
第三步因为f1函数没有指定返回值是谁，第一步的赋的值是6，这一步只是执行RET指令返回"返回值",即6
f1函数的defer立即执行函数修改的是x的值，不是返回值的值
*/
func f1() int {
	x := 5
	x = x + 1
	// defer定义了一个立即执行函数
	defer func() {
		x++
		println("x is ", x) // 7
	}()
	return x // 6
}

/*
f2函数首先声明了返回值为 x int,
随后执行函数体，函数体只有defer语句和return命令，分三步走
第一步执行return语句为返回值赋值5，因为函数定义的时候指定了返回值为x，所以赋值语句相当于"返回值=x=5"
第二步执行defer的立即执行函数x++，此时x=5+1=6，
第三步执行RET指令时返回的是x的值，也就是6
f2的defer函数修改的返回值的值
*/
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

/*
f3函数首先声明返回值为y int，随后执行函数体x := 5 和 return命令
return命令分三步，
第一步赋值"返回值"y=x=5，
第二步执行defer语句，defer语句将x++为6，
第三步RET指令返回y，因为赋值操作已经在第一步执行过了，defer改变的只是x的值，不是返回值y，所以也就是y=5
f3函数修改的是x的值，不是已经指定的返回值y
*/
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

/*
f4函数指定返回值为x变量，
直接执行函数体的return语句，分三步走，
第一步，指定返回值为5
第二步，执行defer的立即执行函数，传入x变量执行x++，在函数中改变的只是传入值的副本
第三步，返回返回值x，即为5
*/
func f4() (x int) {
	defer func(x int) {
		x++ // 改变的只是函数中x的副本，不影响真正x的值
		fmt.Println("defer x value: ", x)
	}(x)
	return 5 // 返回值 = x = 5
}

// 1、返回值 x = 5
// 2、defer x = 6 ，defer的return x 没有变量接收
// 3、RET 返回 x =5
func f5() (x int) {
	defer func(x int) int {
		x++
		fmt.Println(x) // 1
		return x
	}(x)
	fmt.Println(x) // 0
	return 5
}

// 传一个x的指针到匿名函数中
func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5
	// 1、返回值 = x = 5
	// 2、defer x = 6, 因为传入的是x地址，所以改变的是原来的x值
	// 3、RET 返回 x = 6
}

func main() {
	fmt.Println(intSum3(3, 4, 5))

	// 调用函数可以通过  函数名()的方式
	r := sum(1, 2)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(r)
	printName("zhi")
	fmt.Println("")

	// 调用有返回值的函数时，可以不接收其返回值

	// defer Demo
	deferDemo()

	/*
		原子，本意即为不能被进一步分割的最小粒子，而原子操作则意为“不可被中断的一个或一系列操作”
		简单理解，原子操作就是指令或指令集本身是否是一口气执行完成而没有分步或中断
	*/

	// 在Go语言中，return语句在底层实现时并不是原子操作（），它分为 给返回值赋值 和 RET指令 两步。
	// 而defer语句执行的时机就是在return语句的返回值赋值和RET指令两步之间（）

	// defer语句执行的时机
	/*
		Go语言中return语句（不是原子操作）底层实现：
			return X：    返回值=X  ->  RET指令
			第一步：返回值赋值
			第二步：真正的RET返回

		defer语句执行时机：
		  	return X：    返回值=X	->	逆序运行defer	->	RET指令
			第一步：返回值赋值
			第二步：按照"先定义后执行"的顺序延迟执行多个defer语句
			第三步：真正的RET返回
	*/

	// 一个函数中可以有多个defer语句，按照先定义后执行的顺序延迟执行多个defer语句

	// 由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。
	fmt.Println("defer经典案例输出值：")
	fmt.Printf("defer函数f1输出： %d \n", f1())
	fmt.Printf("defer函数f2输出： %d \n", f2())
	fmt.Printf("defer函数f3输出： %d \n", f3())
	fmt.Printf("defer函数f4输出： %d \n", f4())
	fmt.Printf("defer函数f5输出： %d \n", f5())
	fmt.Printf("defer函数f6输出： %d \n", f6())
	// fmt.Println("" + string(f1()))

	// 5、变量的作用域

	// 全局变量：全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效
	// 局部变量：局部变量又分为两种，函数内定义的变量无法在该函数外使用

	/*
		函数中查找变量的顺序：
		① 先在函数内部查找
		② 找不到就去函数外部查找，一直找到全局变量的范围
		同时，如果局部变量和全局变量重名，优先访问局部变量
	*/

	// 6、函数类型与变量

	/* 6.1 定义函数类型
	   使用 type 关键字来定义一个函数类型，格式如下

	   type calculation func(int, int) int

	   上面的语句定义了一个 calculation 类型，它是一种函数类型，这种函数接收两个 int 类型的参数并且返回一个int类型的返回值
	   也就是说，凡是满足这个条件的函数都是 calculation 类型的函数，例如下面的add和sub就是calculation类型

	   func add(x, y int) int {
		   return x + y
	   }

	   func sub(x, y int) int {
		   return x - y
	   }

	   add和sub都能赋值给calculation类型的变量
	   var c calculation
	   c = add
	*/

	// 6.2 函数类型变量
	/*

		var c calculation
		c = add
		fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
		fmt.Println(c(1, 2))            // 像调用add一样调用c

		f := add                        // 将函数add赋值给变量f
		fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
		fmt.Println(f(10, 20))          // 像调用add一样调用f

	*/

	// 6.3 函数可以作为参数的类型
	/*
		// 函数f3的 参数x 是一个"返回值是int类型且无参数的函数"
		// 调用f3时，只要传入参数满足 "返回值是int类型的无参数的函数" 即可
		func f3(x func() int){
			ret := x()
			fmt.Println(ret)
		}
	*/

	// 6.4 函数还可以作为返回值
	// 函数f5的参数x是一个返回值是int类型的无参函数，同时f5的返回值是一个"参数是两个int类型变量，返回值也是int变量的函数"
	/*
		func f5(x func() int) func(int, int) int {

		}
	*/

	// 7、匿名函数
	// 函数体内部不能声明带名字的函数
	// 所以可以定义一个匿名函数，将其赋值给变量进行调用
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	// 如果函数仅需调用一次，还可以简写成立即执行函数
	// 因为如果仅执行一次，那么就不必要找个变量存储它
	// 立即执行函数，在匿名函数结尾加上   括号并传入参数
	func(x, y int) {
		fmt.Printf("立即执行函数执行结果： ")
		fmt.Println(x + y)
	}(100, 200)

	// 8、闭包
	// 文件bibao{1,4}

	// 9、init()函数
	//在Go语言程序执行时会自动触发包内部init()函数的调用，
	//需要注意的是：init()函数没有参数也没有返回值，init()函数在程序运行时自动被调用执行，不能在代码中主动调用它
	/*
		包中init函数的执行时机：	全局声明 -> init() -> main()

		导入包顺序：
			main --import--> A --import--> B --import--> C
		初始化函数执行顺序：
			C.init() --> B..init() --> A.init() --> main.init()
	*/
}
