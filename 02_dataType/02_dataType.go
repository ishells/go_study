package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

func main() {
	// 1、整型
	/*
		整型分为以下两个大类：
			按长度分为：int8、int16、int32、int64
			对应无符号整型：uint8、uint16、uint32、uint64

			uint8	无符号 8位整型 (0 到 255)
			uint16	无符号 16位整型 (0 到 65535)
			uint32	无符号 32位整型 (0 到 4294967295)
			uint64	无符号 64位整型 (0 到 18446744073709551615)
			int8	有符号 8位整型 (-128 到 127)
			int16	有符号 16位整型 (-32768 到 32767)
			int32	有符号 32位整型 (-2147483648 到 2147483647)
			int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)

			uint	32位操作系统上就是uint32，64位操作系统上就是uint64
			int		32位操作系统上就是int32，64位操作系统上就是int64
			uintptr	无符号整型，用于存放一个指针
	*/
	// 通用占位符：
	//	%b	二进制
	//	%c	该值对应的unicode码值
	//  %d	十进制
	//	%f 	浮点数
	//	%o	八进制
	//	%x	十六进制，表示为a-f
	//	%X	十六进制，表示为A-F
	//	%s	字符串
	//  %T	显示数据类型

	//	%t	布尔类型

	//	%v	值
	// 	%+v	类似于%v,输出结构体时会添加字段名
	//	%#v	值的go语法表示形式（输出更详细）
	//	%%	输出百分号

	// 2、浮点型
	/*
		Go语言支持两种浮点型数：float32和float64。
		float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32
		float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64
	*/
	fmt.Printf("保留两位小数的pi值： %.2f\n", math.Pi)

	// 3、布尔值
	// go语言以 bool 类型进行声明布尔型数据，只有 true 和 false 两个值
	/*
		注意：
			① 布尔类型变量的默认值为false。
			② Go 语言中不允许将整型强制转换为布尔型.
			③ 布尔型无法参与数值运算，也无法与其他类型进行转换。
	*/

	// 4、字符串
	/*

			① go语言中的字符串只能使用双引号"",单引号''的是字符
			② 1 字节 = 8 Bit（8个二进制位）
		   	  1个字符 'A' = 1 字节
		   	  1个utf8编码汉字 '学' = 一般占3个字节

	*/
	// 4.1 多行字符串(原样输出，带格式)
	s := `
			1
			2
			3
	`
	fmt.Println("3、多行字符串原样输出： " + s)

	// 字符串求长度 len(str)
	// len()返回的是字符串长度的字节数，与其他编程语言不同的地方

	str_len := "直接"
	fmt.Printf("使用len方法计算\"直接\"字节长度：%v\n", len(str_len))

	// 可以使用utf8库下的RuneCountInString方法获取字符长度
	fmt.Printf("使用utf8.RuneCountInString方法计算\"直接\"字符长度：%v\n", utf8.RuneCountInString(str_len))

	// 正经计算字符长度的时候，就使用utf8下的方法来计算，不要使用len()/3来进行计算（避免错误，有的表情之类的是4字节）

	// 4.2 字符串拼接  +或者fmt.Sprintf
	var name string = "zhijb"
	var age string = "22"

	word := name + age
	fmt.Println("4.1 +字符串拼接： " + word)
	ss := fmt.Sprintf("%s%s", name, age)
	fmt.Println("4.2 Sprintf字符串拼接：" + ss)

	// 4.3 字符串分割 strings.Split
	var dir_location string = "/home/zhijb/data_dir/code_dir/src"
	fmt.Println()
	fmt.Println(strings.Split(dir_location, "/"))

	// 4.4 字符串判断包含 strings.Contains
	fmt.Print("6.1 字符串是否包含字符zhijb： ")
	fmt.Println(strings.Contains(dir_location, "zhijb"))
	fmt.Print("6.2 字符串是否包含字符支付宝： ")
	fmt.Println(strings.Contains(dir_location, "zhifubao"))

	// 4.5 前缀开头判断strings.HasPrefix()	结尾判断strings.HasSuffix()
	fmt.Print("7.1 判断路径是否以/home字符串开头： ")
	fmt.Println(strings.HasPrefix(dir_location, "/home"))
	fmt.Print("7.2 判断路径是否以src结尾： ")
	fmt.Println(strings.HasSuffix(dir_location, "src"))

	// 4.6 字符串位置判断
	fmt.Print("8.1 d字符在字符串中第一次出现位置： ")
	fmt.Println(strings.Index(dir_location, "d"))
	fmt.Print("8.2 d字符在字符串中最后一次出现位置： ")
	fmt.Println(strings.LastIndex(dir_location, "d"))

	// 5、字符类型：byte 和 rune
	/*
		Go 语言的字符有以下两种：
			uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
			rune类型，代表一个 UTF-8字符

		当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32 !!
		Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾
	*/

	// 遍历字符串
	// for i := 0; i < len(s); i++ { //byte
	// 	fmt.Printf("%v(%c) ", s[i], s[i])
	// }
	fmt.Println()

	string1 := "hello世界"
	for _, r := range string1 { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()

	var rune12 rune = '你'
	fmt.Println(rune12)

	// 5、修改字符串
	/*
		要修改字符串，需要先将其转换成 []rune 或 []byte,字符用[]byte,文字使用[]rune
		完成后再转换为string。
		无论哪种转换，都会重新分配内存，并复制字节数组
	*/
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))

	// 6、类型转换
	/*
		go语言中类型转换与python类似，直接使用 data_type(变量/复杂算式/函数返回值) 进行强制转换
		并且原数据类型不会被改变

		go语言中没有隐式类型转换，即以下示例代码会报错：
			var num1 int = 13
			var num2 unint = 15
			// 下面这一句无法通过编译，因为golang是强类型语言，并且不会帮你做任何的转换
			println(a == c)
	*/
	var int_num int = 8
	fmt.Printf("%T\n", float32(int_num))
	fmt.Printf("%T\n", int_num)

	stringsss := "s"
	stringbyte := '@'
	fmt.Printf("%T --- %T", stringsss, stringbyte)
}
