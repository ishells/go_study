package main

import (
	"fmt"
	"strings"
)

// 闭包进阶

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		// 判断字符串后缀
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	// jpgFunc、txtFunc仅仅只是用来接收makeSuffixFunc函数的返回值，类型为一个匿名函数
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")

	// 如果不是以固定结尾的，就在后边加上固定字符串
	fmt.Printf("%T\n", jpgFunc)
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}
