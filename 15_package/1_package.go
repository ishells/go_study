package main

import (
	"go_study/15_package/tools_ops"
)

/*
	一个package中的标识符（变量名/函数名/接口体/接口/）如果首字母是小写的，就表示它是私有的（）

*/

/*
	vscode 一次只能打开一个项目，所以在go项目中，被打开的根目录要有一个go.mod文件（一个项目对应一个go.mod）

	使用vscode初始化项目，打开项目文件夹，比如 /.../src/  其中src/是项目的根目录，
	在此目录下执行 go mod init src(项目名称/即模块名)，会生成一个新的go.mod文件

	当需要导入本地自定义的package时，格式为
	import （
		别名 根目录名/目录层级直至需要导入包所在的文件夹
	）
*/

func main() {
	//fmt.Println(calc.Add(10, 20))
	println(tools_ops.Add(1, 2))
}
