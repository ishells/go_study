package main

import (
	"fmt"
)

// 指针、容量、长度

// 切片slice

// 在golang内部，slice、map、channel本身就是指针类型

// 切片的定义
// 切片基于数组类型做了一层封装（依赖数组实现），支持自动扩容
// 切片是一个引用类型，它的内部结构包括 地址 长度 和 容量。切片一般用于快速地操作一块数据集合
func main() {
	// 1、声明切片
	// 声明切片的基本语法如下(切片不用声明长度)：

	// var Name []Type

	// name := []Type{}

	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	// var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较

	// 切片仅声明之后并未分配内存地址，其值为nil，需要初始化之后才会有内存地址

	// 2、切片的本质
	// 切片本身不保存数值，本质就是对底层的数组进行操作
	// 切片的本质就是对底层数组的封装，它包含三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）
	// https://www.liwenzhou.com/posts/Go/06_slice/ 参考理解切片的长度和容量两个概念
	// 切片容量在讨论底层引用数组时使用

	// 切片的长度即为从底层数组中所获取到的元素个数，而切片的容量则是 从底层数组获取到第一个元素开始，直到 底层数组的最后一个元素
	// 举例:
	// 数组 var a  = [8]int{0,1,2,3,4,5,6,7}
	// 切片 var s1 = a[:5] 此时，s1的值为{0,1,2,3,4},切片s1的长度len为5，容量cap为8（从底层数组的0到7，容量即为8）
	// 切片 var s1 = a[3:6] 此时，s1的值为{3,4,5},切片s1的长度len为3，容量cap为5（从底层数组的3到7，容量则为5）

	// 3、切片的长度和容量
	// 内置的len()函数求长度，内置的cap()函数求容量

	// 4、切片表达式
	// go里边的切片表达式类似python中的切片操作，就不详细记笔记了
	// 简单的切片表达式		a[low:high]
	// 完整的切片表达式(字符串不支持完整切片),按照 high减low 左包右不包进行切片获取元素，而容量会被设置为 max减low		a[low:high:max]
	// 完整切片表达式需要满足的条件是0 <= low <= high <= max <= cap(a)，其他条件和简单切片表达式相同
	// max 的值不能超出底层数组的总长度，否则就会越界

	array1 := [5]int{1, 2, 3, 4, 5}
	s1 := array1[1:3:5]
	fmt.Printf("s1:%v len(s1):%v cap(s1):%v \n", s1, len(s1), cap(s1))

	// 5、使用make()函数构造切片
	// 不仅可以通过数组来构造切片，如果需要动态的创建一个切片，可以使用内置的make()函数
	// make()函数创建完切片之后会填入默认值（即各个数据类型的零值）

	// make([]T, size ,cap)
	// T:切片的元素数据类型
	// size:切片中元素的数量,即长度
	// cap:切片的容量
	array2 := make([]int, 2, 10)
	fmt.Println(array2)      //[0 0]
	fmt.Println(len(array2)) //2
	fmt.Println(cap(array2)) //10

	// 6、判断切片是否为空

	// 要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断

	// 切片不能直接比较
	// 切片之间是不能比较的，不能使用==操作符来判断两个切片是否含有全部相等元素。
	// 切片唯一合法的比较操作是和 nil 比较
	// 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
	// 但是我们不能说一个长度和容量都是0的切片一定是nil
	// 所以要判断一个切片是否是空的，要使用len(s) == 0来判断，不应该使用s == nil来判断

	// 下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容
	slice1 := make([]int, 3) //[0 0 0]
	slice2 := slice1         //将slice1直接赋值给slice2，slice1和slice2共用一个底层数组
	slice2[0] = 100
	fmt.Println(slice1) //[100 0 0]
	fmt.Println(slice2) //[100 0 0]
	/*
	   由于切片是引用类型，所以slice1和slice2其实都指向了同一块内存地址。修改slice2的同时slice1的值也会发生变化
	*/

	// 下面的代码演示了切片对数组的引用，修改切片的值并不会影响底层数组的值
	array3 := [6]int{0, 1, 2, 3, 4, 5}
	slice4 := array3

	slice5 := slice4
	slice4[0] = 100
	slice5[1] = 200
	// 因为 slice4、slice5都是相当于引用的底层数组值，所以不管修改它们谁的值，都不会互相影响
	// 修改之后 array3:[0 1 2 3 4 5] slice4:[100 1 2 3 4 5] slice5:[0 200 2 3 4 5]
	fmt.Printf("array3:%v slice4:%v slice5:%v \n", array3, slice4, slice5)

	// 7、切片的遍历和数组一样，支持 索引遍历 和 for range 遍历

	// 8、append()方法为切片增加元素
	// 内建函数append()可以为切片动态添加元素。

	// 调用append()函数必须用原来的切片变量接受返回值
	// 举例说明，如创业公司xxx初始就2个人，只需要2个工位，
	// 后来公司规模变大，就需要变更办公地址以容纳更多的员工，但是公司名称还是xxx，仅仅只是办公地址变了

	// append()可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加 … 符号）
	var slice6 []int
	slice6 = append(slice6, 1)       // [1]
	slice6 = append(slice6, 2, 3, 4) // [1 2 3 4]
	slice7 := []int{5, 6, 7}
	slice6 = append(slice6, slice7...) // [1 2 3 4 5 6 7]

	// 注意：通过var声明的零值切片可以在append()函数直接使用，无需初始化
	var slice8 []int
	slice8 = append(slice8, 1, 2, 3)

	// append() 方法可为切片动态添加元素，每个切片都指向一个底层数组
	// 当当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换
	// append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	// 从输出结果可以看出来
	// ① append()函数将元素追加到切片的最后并返回该切片。
	// ② 切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。

	// 9、切片的扩容策略
	// 查看$GOROOT/src/runtime/slice.go源码
	// https://www.liwenzhou.com/posts/Go/06_slice/#autoid-2-5-0
	/*
		首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
		否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
		否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
		如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）
	*/

	// 10、使用copy()函数复制切片

	// copy()复制切片，copy(dst, src)

	// copy()函数不会主动扩容目的切片，如果目的切片容量不足也无法完成copy()
	slice9 := []int{1, 2, 3, 4, 5}
	slice10 := make([]int, 5, 5)
	slice11 := slice9
	copy(slice10, slice9) //使用copy()函数将切片slice9中的元素复制到切片slice10
	fmt.Println(slice9)   //[1 2 3 4 5]
	fmt.Println(slice10)  //[1 2 3 4 5]
	fmt.Println(slice11)  //[1 2 3 4 5]
	slice9[0] = 1000
	fmt.Println(slice9)  //[1000 2 3 4 5]
	fmt.Println(slice10) //[1 2 3 4 5]  copy()复制切片的值不受原值的影响
	fmt.Println(slice11) //[1000 2 3 4 5]

	// 11、Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素
	// 从切片中删除元素
	slice12 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	slice10 = append(slice12[:2], slice12[3:]...)
	fmt.Println(slice12) //[30 31 33 34 35 36 37 37]

	x1 := [...]int{1, 3, 5} // 数组
	// [:]就是从头取到尾，左包右也包
	s2 := x1[:] // 切片，从数组x1的开始切到最后
	fmt.Println(s2, len(s2), cap(s2))
	/*
		1、切片不保存具体的值
		2、切片对应一个底层数组
		3、底层数组都是占用一块连续的内存
	*/
	fmt.Printf("%p\n", &x1)
	fmt.Printf("%p\n", &s2[0])
	fmt.Printf("%p\n", &s2)
	s2 = append(s2[:1], s2[2:]...) // 修改了切片的底层数组（非原始数组），将s2数组索引为0的元素和索引为2的元素拼接起来，相当于删除了索引为2的元素
	fmt.Printf("数组x1地址：%p 切片s2地址：%p \n", &x1, &s2)
	fmt.Printf("数组x1值:%v 切片s2值:%v \n", x1, s2) // 数组x1:[1 5 5] 切片s2:[1 5]
	x1[0] = 100
	fmt.Printf("数组x1值:%v 切片s2值:%v \n", x1, s2) // 数组x1值:[100 5 5] 切片s2值:[100 5]
	// * 总结起来一句话，修改数组的值会同时修改引用该数组的切片值，而仅仅修改切片的值则不会修改其引用的原数组值（只是修改了切片对应内存地址的底层数组的值） *

	// 使用append修改或删除切片的值，会（相当于挪动元素位置）修改该切片所引用的原数组值，具体看下例子：

	example_array := [...]int{1, 2, 3, 4, 5, 6}
	example_slice := example_array[:]
	// 删除掉切片索引为2处的元素
	fmt.Printf("未删除切片索引为2处的元素前 example_slice值为 %v\n", example_slice)
	example_slice = append(example_slice[:2], example_slice[3:]...)
	fmt.Printf("删掉索引2处元素后example_slice切片值为:%v \n", example_slice) // [1 2 4 5 6]
	fmt.Printf("删掉索引2处元素后example_array数组值为:%v \n", example_array) // [1 2 4 5 6 6] 前面的append操作相当于把数组后三个元素往前推了三位，其他元素不变

	// 切片真实面试题（非常nice）
	var a_test = make([]int, 5, 10)
	fmt.Printf("使用make创建完长度为5的切片之后，切片已经有5个int默认值了，即0 \n")
	fmt.Println(a_test)
	for i := 0; i < 10; i++ {
		a_test = append(a_test, i)
	}
	fmt.Println("a_test值是: ", a_test)

	// 12、指针

	/*
		指针使用流程：
			定义指针变量。
			为指针变量赋值。
			访问指针变量中指向地址的值。
	*/

	/*
		go 语言中没有指针操作，只有两个符号：
		① & 获得地址
		② * 根据地址取值
	*/

	var a_a int = 20     /* 声明实际变量 */
	var int_pointer *int /* 声明指针变量 */

	int_pointer = &a_a /* 指针变量的存储地址 */
	fmt.Printf("a_a 变量的地址是: %x\n", &a_a)
	/* 指针变量的存储地址 */
	fmt.Printf("int_pointer 变量储存的指针地址: %x\n", int_pointer)
	/* 使用指针访问值 */
	fmt.Printf("*int_pointer 变量的值: %d\n", *int_pointer)

	// 使用 & 取地址
	num := 23
	mem_add := &num
	fmt.Println(mem_add)
	fmt.Printf("mem_add变量的类型是%T, 代表int类型的指针 \n", mem_add)

	num1 := *mem_add
	fmt.Println(num1)
	fmt.Printf("%T \n", num1)

	// 指针类型数据的零值是 nil，很多类型数据的零值都是nil

	// 使用 new(Type) 函数申请一个内存地址,new()函数申请的是内存地址，如果需要赋值还需要结合*符号使用
	mem_add3 := new(int)
	*mem_add3 = 100
	print(*mem_add3)

	// 使用 make() 函数分配内存，区别于 new() 函数，make() 只能用于slice、map、chan的内存创建
	// func make(t Type, size ...IntegerType)

	// make() 函数创建slice、map的格式不同！

	/* 例：
	make创建map：	make(type,cap)  ->  make(map1[string]int,200)
	make创建slice：	make([]T, size ,cap)  ->  make([]int, 5, 10)
	make创建元素为map类型的slice：	make([]map[int]string, 5, 10) slice格式是[]type，此例Type是"键为int，value为string"的map
	make创建元素为slice类型的map：	make(map[string][]int, 10)  map[string]Type , 此例中Type是int类型的slice
	声明元素类型为map的切片时，切片和map都需要分别初始化
	*/

	/* make和new的区别：
	① make和new都是用来申请内存的
	② new很少用，一般用来给基本数据类型申请内存，例如string、int，new()返回的是对应类型的指针（*string）
	③ make用来给slice、map、chan申请内存，make函数返回的是对应的这三个类型本身
	*/

}
