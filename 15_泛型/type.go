// 视频35：:35

package _6_泛型

import "fmt"

/*
什么是泛型?

泛型允许程序员在强类型程序设计语言中编写代码时使用一些以后才指定的类型，在实例化时作为参数指明这些类型。ーー换句话说，在编写某些代码或数据结构时先不提供值的类型，而是之后再提供。

泛型是一种独立于所使用的特定类型的编写代码的方法。使用泛型可以编写出适用于一组类型中的任何一种的函数和类型。

// 反转切片 /
func reverse(s []int) []int {
	l := len(s)
	newS := make([]int, l)

	for k, v := range s {
		newS[l-k-1] = v
	}
	return newS
}

如上述切片反转函数，只能接收[]int类型的参数，如果想支持[]float64类型参数、想支持[]xxx就需要额外定义一个reverseXXX的函数，

一遍一遍地编写相同的功能是低效的，实际上这个反转切片的函数并不需要知道切片中元素的类型，但为了适用不同的类型我们把一段代码重复了很多遍。
Go1.18之前我们可以尝试使用反射去解决上述问题，但是使用反射在运行期间获取变量类型会降低代码的执行效率并且失去编译期的类型检查，同时大量的反射代码也会让程序变得晦涩难懂。

类似这样的场景就非常适合使用泛型。从Go1.18开始，使用泛型就能够编写出适用所有元素类型的“普适版”reverse函数

func reverseWithGenerics[T any](s []T) []T {
	l := len(s)
	newS := make([]T, l)

	for k, v := range s {
		newS[l-i-1] = v
	}
	return newS
}

*/

/*
泛型为Go语言添加了三个新的重要特性:
	函数和类型的类型参数。
	将接口类型定义为类型集，包括没有方法的类型。
	类型推断，它允许在调用函数时在许多情况下省略类型参数。

学习函数的过程已经知道，定义函数时可以指定形参，调用函数时需要传入实参（下面例子中的a,b就是形参，调用sum的1,2就是实参）
func sum(a, b int)int{
	return a+b
}
sum(1,2)

Go1.18后，GO语言的 函数参数和类型 支持使用 类型参数，类型参数列表看起来像普通的参数列表，只不过它使用方括号[] 而不是圆括号()
*/

// 泛型用法: 类型形参 类型约束
func compare[T int | float64](a, b T) T {
	// 其中，"T int | float64" 统称类型参数列表，"T" 称为类型形参，  "int | float64" 称为类型约束
	if a > b {
		print(a)
		return a
	}
	print(b)
	return b
}

// 类型实例化
/// 这个compare函数就同时支持int和float64两种类型，也就是说当调用compare函数时，既可以传入int类型参数，也可以传入float64参数
/// compare[int](10, 20)  调用时 "[int]" 称为类型实参
/// 向 compare函数提供类型参数（在本例子中为int和float64）称为实例化，
/*
	类型实例化分两步进行：
		1、首先编译器在整个泛型函数或类型中将所有类型形参替换为它们各自的类型实参
		2、其次，编译器验证每个类型参数是否满足相应的约束
	在成功实例化之后，我们将得到一个非泛型函数，它可以像任何其它函数一样被调用，例如
	fcompare := compare[float64] // 类型实例化，编译器生成T=float64的min函数
	num = fcompare(1.2, 2.3)
*/

// 类型参数的使用
/// 除了函数中支持使用类型参数列表外，如切片、Map、struct等的类型也可以使用类型参数列表
/*
type Slice[T int |string] []T

type Map[K int | string,V float32 | float64] map[K]V

type Tree[T interface{}] struct {
	left,right *Tree[T]
	value T
}
在上述泛型类型中，T、K、V都属于类型形参，类型形参后面是类型约束，类型实参需要满足对应的类型约束

泛型类型可以有方法，例如为上面的Tree实现一个查找元素的Lookup方法。
func (t *Tree[T]) Lookup(t T) *Tree[T] {...}

要使用泛型类型，就必须进行实例化。Tree[String] 是使用类型实参string实例化Tree的示例。
var stringTree Tree[string]
*/

// 类型约束
/*  普通函数中的每个参数都有一个类型; 该类型定义一系列值的集合。
    类似于参数列表中每个参数都有对应的参数类型，类型参数列表中每个类型参数都有一个类型约束。
    类型约束定义了一个类型集——只有在这个类型集中的类型才能用作类型实参。

	Go语言中的类型约束是接口类型。

// 类型约束接口可以直接在类型参数列表中使用。
// 类型约束字面量，通常外层interface{}可省略
func min[T interface{ int | float64 }](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// 作为类型约束使用的接口类型可以事先定义并支持复用。
// 事先定义好的类型约束类型
type Value interface {
	int | float64
}
func min[T Value](a, b T) T {
	if a <= b {
		return a
	}
	return b
}


在使用类型约束时，如果省略了外层的interface{}会引起歧义，那么就不能省略。例如：

type IntPtrSlice [T *int] []T  // T*int ?

type IntPtrSlice[T *int,] []T  // 只有一个类型约束时可以添加逗号`,`
type IntPtrSlice[T interface{ *int }] []T // 使用interface{}包裹

*/

// 类型集（以下看不太懂，回头有时间再研究 https://www.liwenzhou.com/posts/Go/generics/）
/*

**Go1.18开始,接口类型的定义也发生了改变，由过去的接口类型定义方法集（method set）变成了接口类型定义类型集（type set）。**也就是说，接口类型现在可以用作值的类型，也可以用作类型约束。
可见图 https://www.liwenzhou.com/images/Go/generics/type-set.png

把接口类型当做类型集相较于方法集有一个优势: 我们可以显式地向集合添加类型，从而以新的方式控制类型集。

Go语言扩展了接口类型的语法，让我们能够向接口中添加类型。例如
type V interface {
	int | string | bool
}
上面的代码就定义了一个包含 int、 string 和 bool 类型的类型集。
（图片地址 https://www.liwenzhou.com/images/Go/generics/type-set-2.png）

从 Go 1.18 开始，一个接口不仅可以嵌入其他接口，还可以嵌入任何类型、类型的联合或共享相同底层类型的无限类型集合。

当用作类型约束时，由接口定义的类型集精确地指定允许作为相应类型参数的类型。
	|符号

	T1 | T2表示类型约束为T1和T2这两个类型的并集，例如下面的Integer类型表示由Signed和Unsigned组成。

	type Integer interface {
		Signed | Unsigned
	}

	~符号

	~T表示所以底层类型是T的类型，例如~string表示所有底层类型是string的类型集合。

	type MyString string  // MyString的底层类型是string
	注意：~符号后面只能是基本类型。


接口作为类型集是一种强大的新机制，是使类型约束能够生效的关键。目前，使用新语法表的接口只能用作类型约束。

any接口
空接口在类型参数列表中很常见，在Go 1.18引入了一个新的预声明标识符，作为空接口类型的别名。

// src/builtin/builtin.go

type any = interface{}
由此，我们可以使用如下代码：

func foo[S ~[]E, E any]() {
	// ...
}
constrains
https://pkg.go.dev/golang.org/x/exp/constraints 包提供了一些常用类型。

类型推断
最后一个新的主要语言特征是类型推断。从某些方面来说，这是语言中最复杂的变化，但它很重要，因为它能让人们在编写调用泛型函数的代码时更自然。

函数参数类型推断
对于类型参数，需要传递类型参数，这可能导致代码冗长。回到我们通用的 min函数：

func min[T int | float64](a, b T) T {
	if a <= b {
		return a
	}
	return b
}
类型形参T用于指定a和b的类型。我们可以使用显式类型实参调用它：

var a, b, m float64
m = min[float64](a, b) // 显式指定类型实参
在许多情况下，编译器可以从普通参数推断 T 的类型实参。这使得代码更短，同时保持清晰。

var a, b, m float64

m = min(a, b) // 无需指定类型实参
这种从实参的类型推断出函数的类型实参的推断称为函数实参类型推断。函数实参类型推断只适用于函数参数中使用的类型参数，而不适用于仅在函数结果中或仅在函数体中使用的类型参数。例如，它不适用于像 MakeT [ T any ]() T 这样的函数，因为它只使用 T 表示结果。

约束类型推断
Go 语言支持另一种类型推断，即约束类型推断。接下来我们从下面这个缩放整数的例子开始：

// Scale 返回切片中每个元素都乘c的副本切片
func Scale[E constraints.Integer](s []E, c E) []E {
    r := make([]E, len(s))
    for i, v := range s {
        r[i] = v * c
    }
    return r
}
这是一个泛型函数适用于任何整数类型的切片。

现在假设我们有一个多维坐标的 Point 类型，其中每个 Point 只是一个给出点坐标的整数列表。这种类型通常会实现一些业务方法，这里假设它有一个String方法。

type Point []int32

func (p Point) String() string {
    b, _ := json.Marshal(p)
    return string(b)
}
由于一个Point其实就是一个整数切片，我们可以使用前面编写的Scale函数：

func ScaleAndPrint(p Point) {
    r := Scale(p, 2)
    fmt.Println(r.String()) // 编译失败
}
不幸的是，这代码会编译失败，输出r.String undefined (type []int32 has no field or method String的错误。

问题是Scale函数返回类型为[]E的值，其中E是参数切片的元素类型。当我们使用Point类型的值调用Scale（其基础类型为[]int32）时，我们返回的是[]int32类型的值，而不是Point类型。这源于泛型代码的编写方式，但这不是我们想要的。

为了解决这个问题，我们必须更改 Scale 函数，以便为切片类型使用类型参数。

func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
    r := make(S, len(s))
    for i, v := range s {
        r[i] = v * c
    }
    return r
}
我们引入了一个新的类型参数S，它是切片参数的类型。我们对它进行了约束，使得基础类型是S而不是[]E，函数返回的结果类型现在是S。由于E被约束为整数，因此效果与之前相同：第一个参数必须是某个整数类型的切片。对函数体的唯一更改是，现在我们在调用make时传递S，而不是[]E。

现在这个Scale函数，不仅支持传入普通整数切片参数，也支持传入Point类型参数。

这里需要思考的是，为什么不传递显式类型参数就可以写入 Scale 调用？也就是说，为什么我们可以写 Scale(p, 2)，没有类型参数，而不是必须写 Scale[Point, int32](p, 2) ？

新 Scale 函数有两个类型参数——S 和 E。在不传递任何类型参数的 Scale(p, 2) 调用中，如上所述，函数参数类型推断让编译器推断 S 的类型参数是 Point。但是这个函数也有一个类型参数 E，它是乘法因子 c 的类型。相应的函数参数是2，因为2是一个非类型化的常量，函数参数类型推断不能推断出 E 的正确类型(最好的情况是它可以推断出2的默认类型是 int，而这是错误的，因为Point 的基础类型是[]int32)。相反，编译器推断 E 的类型参数是切片的元素类型的过程称为约束类型推断。

约束类型推断从类型参数约束推导类型参数。当一个类型参数具有根据另一个类型参数定义的约束时使用。当其中一个类型参数的类型参数已知时，约束用于推断另一个类型参数的类型参数。

通常的情况是，当一个约束对某种类型使用 ~type 形式时，该类型是使用其他类型参数编写的。我们在 Scale 的例子中看到了这一点。S 是 ~[]E，后面跟着一个用另一个类型参数写的类型[]E。如果我们知道了 S 的类型实参，我们就可以推断出E的类型实参。S 是一个切片类型，而 E是该切片的元素类型。





*/

type CusNumT interface {
	// 支持uint8、int32、float64与int64及其衍生类型
	// ～ 表示支持类型的衍生类型
	// ｜ 表示取并集
	// 多行之间取交集
	uint8 | int32 | float64 | ~int64
	int32 | float64 | ~int64 | uint64
	//Get() string
}

// MyInt64 为int64的衍生类型，是具有基础类型int64的新类型，与int64是不同的类型
type MyInt64 int64

// MyInt32 为int32的别名
type MyInt32 = int32

// 别名与衍生类型区别，衍生类型是不同的类型，所以存在类型转换的过程，别名仅仅是名称不同

func getMaxCusNum[T CusNumT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func CusNumTCase() {
	var a, b int32 = 3, 4
	var a1, b1 MyInt32 = a, b
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(a, b))
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(a1, b1))

	var c, d float64 = 5, 6
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(c, d))

	var e, f int64 = 7, 8
	var g, h MyInt64 = 7, 8
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(e, f))
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(g, h))

}

func main() {
	//compare(1.6, 2.6)
	CusNumTCase()
}
