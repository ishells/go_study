package tools_ops

import "fmt"

var name string = "ops"

func init() {
	fmt.Println("当代码被import时自动执行init函数。。。")
	fmt.Println("全局变量声明在init函数之前执行：" + name)
}
func Add(x, y int) int {
	b := struct {
		Name    string
		Age     int
		Address string
	}{
		Name:    "ops",
		Age:     20,
		Address: "henan",
	}
	fmt.Println(b)
	return x + y
}
