package main

import (
	"fmt"
)

//
//import (
//	"fmt"
//)
//
///*
//var _ UserService = (*UserServiceImpl)(nil) 表达了什么意义
//这行代码是在声明一个接口，并检查一个叫做 UserServiceImpl 的结构体是否满足该接口。
//
//在 Go 编程语言中，接口是定义行为的一种方式，而这行代码则是在进行编译时类型检查。具体来说，它在声明一个名为 UserService 的接口，并使用 *UserServiceImpl 的指针类型进行赋值，这里的 nil 表示 UserServiceImpl 的实例不存在。
//
//这种方式通常被用于确保某个结构体满足某个接口的所有方法，以避免在运行时出现类型错误或缺少实现的错误。如果结构体没有实现接口中的所有方法，编译器会报错。如果结构体实现了接口中的所有方法，则该结构体就可以被视为实现了该接口，可以被赋值给接口类型的变量，从而可以通过接口来调用该结构体的方法。
//*/
//
//type Charger interface {
//	Charge()
//}
//
//type IPhone struct {
//	Name string
//}
//
//type Huawei struct {
//	Name string
//}
//
//func (iphone IPhone) Charge() {
//	fmt.Println("iphone充电。。")
//}
//
//func (hw Huawei) Charge() {
//	fmt.Println("hw充电。。")
//}
//
//func main() {
//	hw := Huawei{}
//	fmt.Printf("%T \n", hw)
//	var charge Charger
//	fmt.Printf("%T \n", charge)
//	iphone := IPhone{}
//	hw.Charge()
//	iphone.Charge()
//
//}

//package main
//
//func charge(deviceType string) func(string, string) string {
//	return func(chargeType, phoneType string) string {
//		return deviceType + " " + phoneType + " " + chargeType + " " + "chargeType..."
//	}
//}
//
//func main() {
//	android := charge("android")
//	ios := charge("ios")
//
//	println(android("type-c", "huawei"))
//	println(ios("lightning", "iphone 14"))
//}

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type Animal struct {
	Name string
}

type Dog struct {
	Feet int
	*Animal
}

func newDog(name string, fleet int) *Dog {
	return &Dog{
		Feet: fleet,
		Animal: &Animal{
			Name: name,
		},
	}
}

func (animal *Animal) move() {
	fmt.Printf("%s move...\n", animal.Name)
}

func (dog *Dog) run() {
	fmt.Printf("%s run...\n", dog.Name)
}

func init() {
	fmt.Println("This is init func test")
}

// 类型约束（泛型特性），(16_泛型) 具体用法事后可详细查阅资料学习
func compare[T int | float64](a, b T) T {
	if a > b {
		print(a)
		return a
	}
	print(b)
	return b
}

func main() {
	// dog1 := newDog("dog", 4)
	// dog1.move()
	// dog1.run()
	// compare(1.6, 2.6)

	// user1 := &User{
	// 	Name:   "ops",
	// 	Age:    20,
	// 	Gender: "man",
	// }
	// jsonData, marshalErr := json.Marshal(*user1)
	// if marshalErr != nil {
	// 	fmt.Println("marshal error")
	// } else {
	// 	fmt.Printf("%v\n", string(jsonData))
	// }

	// user2 := User{}
	// jsonStr := `{"name":"ishells","age":22,"gender":"man"}`
	// if json.Unmarshal([]byte(jsonStr), &user2) != nil {
	// 	fmt.Println("Unmalshal failed")
	// } else {
	// 	fmt.Printf("%v", user2)
	// }
}
