package main

import "fmt"

// 单个声明变量
var name string
var age int
var city string = "北京"

// 批量声明
var (
	school string
	height int
)
var company, org string

// 单一常量声明
const name66 = "zhijb"

// 批量常量声明
const (
	webName = "ishells"
	pi      = 3.141592
)

// 同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 100
	n2
	n3
)

// iota
const (
	iota_1 = iota // 0
	iota_2        // 1
	iota_3        // 2
)

// iota在const关键字出现时被重置为0 const中每新增一行常量声明将使iota计数一次（值+1）
// 使用 _ 匿名变量跳过iota的值
const (
	iota_4 = iota // 0
	iota_5        // 1
	_             // 2,匿名变量不保存，丢弃2
	iota_6        // 3
)

const (
	iota_7 = iota // 0
	num    = 100  // 100
	iota_8 = iota // 2
	iota_9        // iota_9=iota=2+1=3
)

/*
定义数量级 （这里的<<表示左移操作，
1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。
同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
*/
const (
	_  = iota             // 0,匿名变量丢弃
	KB = 1 << (10 * iota) // 1左移10位，变成二进制10000000000，十进制1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

// 多个iota定义在一行
const (
	a, b = iota + 1, iota + 2 //1, 2
	c, d = iota + 1, iota + 2 //c=1+1,d=1+2
	e, f                      //e=iota+1=2+1=3,f=iota+2=2+2=4
)
const (
	g, h = iota, iota
	i, j // 1,1
)

func main() {
	// var sex string
	// 初始化变量
	// name = "zhijb"
	// age = 88
	// school = "xcu"
	// height = 170
	fmt.Println("hello world")
	// fmt.Print(name)                   // 直接打印输出
	// fmt.Println("school is" + school) // 输出后换行
	// fmt.Printf("age is： %d", age)     // 格式化输出

	fmt.Println("n1: ", n1)
	fmt.Println("n2: ", n2)
	fmt.Println("n3: ", n3)

	fmt.Println("iota_1: ", iota_1)
	fmt.Println("iota_2: ", iota_2)
	fmt.Println("iota_3: ", iota_3)

	fmt.Println("iota_4: ", iota_4)
	fmt.Println("iota_5: ", iota_5)
	fmt.Println("iota_6: ", iota_6)

	fmt.Println("iota_7: ", iota_7)
	fmt.Println("iota_8: ", iota_8)
	fmt.Println("iota_9: ", iota_9)
	fmt.Println("KB: ", KB)
	fmt.Println("MB: ", MB)
	fmt.Println("GB: ", GB)
	fmt.Println("TB: ", TB)
	fmt.Println("PB: ", PB)

	fmt.Println("i: ", i)
	fmt.Println("j: ", j)

}

/*
    0、函数外只能放置标识符（变量、常量、函数、类型）的声明，不能放置任何语句
    // 程序的入口函数
   		func main(){
       		fmt.Println("Hello world")
   		}

	// 变量
		单变量声明：
		第一种，指定变量类型，如果没有初始化，则变量值为该类型默认值(零值)
		var v_name v_type
		v_name = value
		第二种，根据值自行判断变量类型
		var v_name = value
		第三种，省略var，使用 :=   (这种方式只能用在函数内部)
		v_name := value


		多变量非全局变量声明：
		var vname1, vname2, vname3 v_type
		vname1, vname2, vname3 = v1, v2, v3

		var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断
		vname1, vname2, vname3 := v1, v2, v3

		多变量全局变量声明：
		var(
			v_name1 v_type1
			v_name2 v_type2
		)

		1、 GO语言中的变量需要先声明才能使用，
			声明变量时需要指定变量类型，
			不像Python之类的解释性语言可以直接使用变量

		2、变量声明时，如果没有进行初始化，会自动对变量对应内存区域执行初始化操作
			例如：
			整型和浮点型变量的默认值为 0
			字符串型变量的默认值为 空字符串
			布尔型变量的默认值为 false
			切片、函数、指针变量的默认为	nil

		3、同一作用域内不支持变量的重复声明

		4、匿名变量：
			在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量，匿名变量用一个下划线_表示
			匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。


	// 常量
		1、常量就是恒定不变的值，多用于定义程序运行期间不会改变的那些值。
		  常量的声明和变量的声明非常相似，关键字是 const
		  常量定义时必须赋值，定义之后不能重复赋值

		2、常量批量定义
		批量定义时，如果未给定值，与上一行保持一致
		const(
			n1 = 100
			n2		// 100
			n3		// 100
		)

		iota：
			iota是go语言的常量计数器，只能在常量的表达式中使用
			iota在const关键字出现时被重置为0,在const中每新增一行常量声明将使iota计数一次（即值+1）
		例：

*/

/*
	变量声明易错点：
		1、变量声明后未使用
		2、变量类型不匹配
		3、同一作用域下，变量只能声明一次
*/
