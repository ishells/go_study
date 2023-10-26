// package main

// import "fmt"

// type person struct {
// 	name string
// }

// type dog struct {
// 	name string
// }

// func speak(s ...interface{}) {
// 	fmt.Printf("%v 在说话", s)
// }

// func main() {
// 	p1 := person{name: "zhi"}
// 	d1 := dog{name: "cai"}
// 	speak(p1)
// 	speak(d1)

// }

// 参考：https://www.liwenzhou.com/posts/Go/12-interface/

// test

package main

import "fmt"

type Person struct {
	Name string
}

type Personer interface {
	Sing()
	Move()
	Eat()
}

func (p *Person) Sing() {
	fmt.Println("Sing....")
}

func (p *Person) Move() {
	fmt.Println("Move....")
}

func (p *Person) Eat() {
	fmt.Println("Eat....")
}

var _ Personer = &Person{}
var _ Personer = (*Person)(nil)

func main() {
	var person Personer
	person = &Person{
		Name: "zhi",
	}
	person.Sing()
}

/*

1、接口的定义
	接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节

	接口是一种由程序员来定义的类型，一个接口类型就是一组方法的集合，它规定了需要实现的所有方法。
	相较于使用结构体类型，当我们使用接口类型说明相比于它是什么更关心它能做什么

	type 接口名	interface{
		方法名1（参数1,参数2...）(返回值1，返回值2...)
		方法名2（参数1,参数2...）(返回值1，返回值2...)
		...
	}

	接口名：Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有关闭操作的接口叫closer等。接口名最好要能突出该接口的类型含义。
	方法名：当接口名首字母是大写且其包含的方法名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
	参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

	举个例子，定义一个包含Write方法的Writer接口。

	type Writer interface{
    	Write([]byte) error
	}
	当你看到一个Writer接口类型的值时，你不知道它是什么，唯一知道的就是可以通过调用它的Write方法来做一些事情。


	用来给变量/参数/返回值等设置类型

2、接口的实现
	接口就是规定了一个需要实现的方法列表，在 Go 语言中一个类型只要实现了接口中规定的所有方法，那么我们就称它实现了这个接口。
	我们定义的Singer接口类型，它包含一个Sing方法。

	// Singer 接口
	type Singer interface {
		Sing()
	}
	我们有一个Bird结构体类型如下。
	type Bird struct {}
	因为Singer接口只包含一个Sing方法，所以只需要给Bird结构体添加一个Sing方法就可以满足Singer接口的要求。

	// Sing Bird类型的Sing方法
	func (b Bird) Sing() {
		fmt.Println("汪汪汪")
	}
	这样就称为Bird实现了Singer接口

2.1 为什么要使用接口？
	可见interface_example.go示例

3、指针接收者实现接口和值接收者实现接口区别
	使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量；而使用指针接收者实现接口时，只有结构体指针类型的变量可以赋值给该接口变量


4、结构体实现多个接口和接口嵌套
// src/io/io.go

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// ReadWriter 是组合Reader接口和Writer接口形成的新接口类型
type ReadWriter interface {
	Reader
	Writer
}

// ReadCloser 是组合Reader接口和Closer接口形成的新接口类型
type ReadCloser interface {
	Reader
	Closer
}

// WriteCloser 是组合Writer接口和Closer接口形成的新接口类型
type WriteCloser interface {
	Writer
	Closer
}

对于这种由多个接口类型组合形成的新接口类型，同样只需要实现新接口类型中规定的所有方法就算实现了该接口类型。


接口也可以作为结构体的一个字段，我们来看一段Go标准库sort源码中的示例。
// src/sort/sort.go

// Interface 定义通过索引对元素排序的接口类型
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}


// reverse 结构体中嵌入了Interface接口
type reverse struct {
    Interface
}
通过在结构体中嵌入一个接口类型，从而让该结构体类型实现了该接口类型，并且还可以改写该接口的方法。
// Less 为reverse类型添加Less方法，重写原Interface接口类型的Less方法
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}
Interface类型原本的Less方法签名为 Less(i, j int) bool，此处重写为r.Interface.Less(j, i)，即通过将索引参数交换位置实现反转。

在这个示例中还有一个需要注意的地方是reverse结构体本身是不可导出的（结构体类型名称首字母小写），sort.go中通过定义一个可导出的Reverse函数来让使用者创建reverse结构体实例。
func Reverse(data Interface) Interface {
	return &reverse{data}
}
这样做的目的是保证得到的reverse结构体中的Interface属性一定不为nil，否者r.Interface.Less(j, i)就会出现空指针panic。



5、空接口
	空接口是指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值

	通常我们在使用空接口类型时不必使用type关键字声明，可以像下面的代码一样直接使用interface{}
	var x interface{}  // 声明一个空接口类型变量x

	空接口做函数的参数可以接受任意类型的参数，结合  变长标识... 使函数可以接收任意长度的任意类型参数

6、接口值
	学习内容来自：https://www.liwenzhou.com/posts/Go/12-interface/
	由于接口类型的值可以是任意一个实现了该接口的类型值（即任意类型都可以实现interface），
	所以接口值除了需要记录具体值之外，还需要记录这个值属于的类型。
	也就是说接口值由“类型”和“值”组成，鉴于这两部分会根据存入值的不同而发生变化，我们称之为接口的动态类型和动态值
           类型type
		  /
	接口值
		  \
		   值value

通过一个示例来加深对接口值的理解。下面的示例代码中，定义了一个Mover接口类型和两个实现了该接口的Dog和Car结构体类型。
type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

func (d *Dog) Move() {
	fmt.Println("狗在跑~")
}

type Car struct {
	Brand string
}

func (c *Car) Move() {
	fmt.Println("汽车在跑~")
}
首先，我们创建一个Mover接口类型的变量m。
var m Mover
此时，接口变量m是接口类型的零值，也就是它的类型和值部分都是nil
我们可以使用m == nil来判断此时的接口值是否为空。
fmt.Println(m == nil)  // true
**注意：**我们不能对一个空接口值调用任何方法，否则会产生panic。
m.Move() // panic: runtime error: invalid memory address or nil pointer dereference
接下来，我们将一个*Dog结构体指针赋值给变量m
m = &Dog{Name: "旺财"}
此时，接口值m的动态类型会被设置为*Dog，动态值为结构体变量的拷贝。
然后，我们给接口变量m赋值为一个*Car类型的值。
var c *Car
m = c
这一次，接口值m的动态类型为*Car，动态值为nil。
**注意：**此时接口变量m与nil并不相等，因为它只是动态值的部分为nil，而动态类型部分保存着对应值的类型。
fmt.Println(m == nil) // false


接口值是支持相互比较的，当且仅当接口值的动态类型和动态值都相等时才相等。

var (
	x Mover = new(Dog)
	y Mover = new(Car)
)
fmt.Println(x == y) // false
但是有一种特殊情况需要特别注意，如果接口值保存的动态类型相同，但是这个动态类型不支持互相比较（比如切片），那么对它们相互比较时就会引发panic。

var z interface{} = []int{1, 2, 3}
fmt.Println(z == z) // panic: runtime error: comparing uncomparable type []int




7、类型断言
	接口值可能被赋值为任意类型的值（因为接口可以被任意类型实现，且接口值分为“类型”和“值”两部分），那我们如何从接口值获取其存储的具体数据呢？
	我们可以借助标准库fmt包的格式化打印获取到接口值的动态类型
	var m Mover

	m = &Dog{Name: "旺财"}
	fmt.Printf("%T\n", m) // *main.Dog

	m = new(Car)
	fmt.Printf("%T\n", m) // *main.Car
	而fmt包内部其实是使用反射的机制在程序运行时获取到动态类型的名称

	而想要从接口值中获取到对应的实际值需要使用类型断言，其语法格式如下。
	x.(T)
	x：表示接口类型的变量
	T：表示断言x可能是的类型
	该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败

	举个例子：

	var m Mover = &Dog{Name: "旺财"}
	v, ok := m.(*Dog)
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "富贵" // 变量v是*Dog类型
	} else {
		fmt.Println("类型断言失败")
	}

	如果对一个接口值有多个实际类型需要判断，推荐使用switch语句来实现
	// justifyType 对传入的空接口类型变量x进行类型断言
	func justifyType(x interface{}) {
		switch v := x.(type) {
		case string:
			fmt.Printf("x is a string，value is %v\n", v)
		case int:
			fmt.Printf("x is a int is %v\n", v)
		case bool:
			fmt.Printf("x is a bool is %v\n", v)
		default:
			fmt.Println("unsupport type！")
		}
	}

	由于接口类型变量能够动态存储不同类型值的特点，所以很多初学者会滥用接口类型（特别是空接口）来实现编码过程中的便捷。
	只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。切记不要为了使用接口类型而增加不必要的抽象，导致不必要的运行时损耗。

	在 Go 语言中接口是一个非常重要的概念和特性，使用接口类型能够实现代码的抽象和解耦，也可以隐藏某个功能的内部实现，
	但是缺点就是在查看源码的时候，不太方便查找到具体实现接口的类型。

	请牢记接口是一种类型，一种抽象的类型。区别于那些具体的类型（整型、数组、结构体类型等），它是一个只要求实现特定方法的抽象类型。

	小技巧： 下面的代码可以在程序编译阶段验证某一结构体是否满足特定的接口类型
	// 摘自gin框架routergroup.go
	type IRouter interface{ ... }

	type RouterGroup struct { ... }

	var _ IRouter = &RouterGroup{}  // 确保RouterGroup实现了接口IRouter
	// 上面的代码中也可以使用var _ IRouter = (*RouterGroup)(nil)进行验证

*/
