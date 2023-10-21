package main

import (
	"fmt"
)

// 数组
// var array_name = [array_len] value_type {}
// 声明数组时 必须指定 存放元素的 类型 和 长度
var a1 = [6]int{0, 1, 2, 3, 4, 5}

func main() {
	// var 数组名 [长度]数据类型
	var array1 [3]bool
	var array2 [3]bool

	// var array9 = [10]int{1,2,3,4,5,6,7,8,8,20}
	fmt.Printf("a1:%T a2:%T", array1, array2)

	// 数组的初始化
	// 如果数组不进行初始化，默认元素都是零值（各种数据类型的初始默认值）
	fmt.Println()
	// 1、初始化方式1，根据数组长度分别进行初始化
	array1 = [3]bool{true, true, true}
	fmt.Println(array1)

	// 2、初始化方式2，自动推断数组长度进行初始化
	array6 := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(array6)

	// 3、初始化方式3，根据索引指定值
	array3 := [5]int{0: 1, 3: 6}
	fmt.Print("array3 is : ")
	fmt.Println(array3)

	// 数组的遍历
	citys := [...]string{"beijing", "shanghai", "shenzhen"}
	// 1、根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// 2、for range 进行遍历
	fmt.Println("range遍历citys：")
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组
	// [[北京 上海] [广州 深圳] [成都 重庆]]
	array4 := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Print("多维数组值：")
	fmt.Println(array4)
	// 根据索引取多维数组的值
	fmt.Print("根据索引取多维数组值：")
	fmt.Println(array4[2][1])

	// 二维数组的遍历
	fmt.Println("二维数组的遍历：")
	for _, v1 := range array4 {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	// 多维数组只有第一层可以使用...来让编译器推导数组长度(即多维数组仅最外层可以使用...)
	//支持的写法
	array7 := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(array7)
	//不支持多维数组的内层使用...
	// b := [3][...]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }

	// 数组是值类型，赋值和传参会复制整个数组。因此，改变副本的值，不会改变数组本身的值
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
	var sum int
	z1 := [...]int{1, 2, 3, 4, 5, 6}
	for _, s := range z1 {
		sum = sum + s
	}
	fmt.Println(sum)
}

func strings(array3 [5]int) {
	panic("unimplemented")
}
