// 视频35：:35

package _6_泛型

import "fmt"

// 泛型用法: 类型形参 类型约束
func compare[T int | float64](a, b T) T {
	if a > b {
		print(a)
		return a
	}
	print(b)
	return b
}

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
