package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json

/*
	1、序列化：		把go语言中的结构体变量 --> json格式的字符串
	2、反序列化：	把json格式的字符串	--> go语言中能够识别的结构体变量
*/

type Person struct {
	// 结构体标签（Tag）
	/*
		    结构体标签的作用就是，如果在json、db、ini配置文件需要该变量值为特定大小写格式的话，
			使用反射来为其指定特定的字段名称，格式为`json:"name"`样式

			Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。
			Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
			`key1:"value1" key2:"value2"`

			结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。
			同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔

			注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。
			结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。
			例如不要在key和value之间添加空格。
	*/
	Name    string `json:"name" db:"name" ini:"name"`
	Age     int    `json:"age"` // 通过指定tag实现json序列化该字段时的key
	address string // 私有不能被json包访问
}

func main() {
	p1 := Person{
		Name: "zhi",
		Age:  22,
	}
	// 序列化
	jsonData, marshalErr := json.Marshal(p1)
	// func Marshal(v interface{})([]byte[],error){}
	// 如果结构体中变量命名为小写name、age，
	// 在本package main包中可以随便调用，但是在使用json方法时，相当于在json包中调用person struct，将无法调用
	// 所以person struct结构体的变量命名只有大写开头才能被调用

	if marshalErr != nil {
		fmt.Printf("marshal failed, err:%v\n", marshalErr)
		return
	}
	fmt.Printf("%v\n", string(jsonData))

	// 反序列化
	jsonStr := `{"name":"ishells","age":22}`
	p2 := Person{}
	if json.Unmarshal([]byte(jsonStr), &p2) != nil {
		fmt.Println("unmarshal failed")
	}
	fmt.Printf("%v\n", p2)

}
