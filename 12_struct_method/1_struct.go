package main

import "fmt"

/*
	结构体常遇到问题：
	1、结构体内部字段要大写！不大写外部无法访问，比如json包调用时 json.Marshal(p1)，p1为实例化的struct变量
*/

/*
	1、结构体：

	结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）

	type 结构体名 struct{
		变量名1	变量类型1
		变量名2	变量类型2
			。。。。
		变量名n	变量类型n
	}

	结构体初始化：
	① 使用键值对初始化
	p5 := person{
		name: "zhi",
		city: "beijing",
		age: 18
	}

	var p7 person
	p7.name = "ishells"
	...
	② 使用值的列表初始化
	p6 := &person{
		"ishells",
		"bj",
		18
	}

	匿名结构体： 多用于临时场景，即结构体中的成员变量不需要定义名称
	匿名结构体不常用！
	var 变量名 struct{
		x string
		y int
	}

*/

// 2、构造函数

// 构造函数作用就是为了节省实例化struct时重复性的重复代码
// 每次实例化struct都需要进行赋值等操作，而如果有了构造函数，每次实例化struct时只用调用构造函数即可（函数本就是为了减少重复性代码）

type person struct {
	name string
	age  int
}

// 构造函数：约定成俗命名使用new开头，newStructName
// 当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
// 构造（结构体变量的）函数，返回值是对应的结构体类型
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

//func main() {
//	u1 := newPerson("user1", 21)
//	u2 := newPerson("user2", 22)
//	fmt.Println(*u1, *u2)
//}

// 3、嵌套结构体
// 一个结构体中可以嵌套包含另一个结构体或结构体指针

// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

func newUser(name string, gender string, province string, city string) *User {
	return &User{
		Name:   name,
		Gender: gender,
		Address: Address{
			Province: province,
			City:     city,
		},
	}
}

func main() {
	user1 := User{
		Name:   "小王子",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City:     "威海",
		},
	}

	// 通过构造函数初始化嵌套结构体
	user2 := newUser("ops", "nan", "shanghai", "shanghai")
	fmt.Println(*user2)

	fmt.Println(user1)
	fmt.Println(user1.Name, user1.Address.City)
}

// 4、嵌套匿名结构体
/*
	//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address //匿名字段
}

func main() {
	var user2 User
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
	user2.City = "威海"                // 匿名字段同样可以省略不用
	// 嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名
	fmt.Printf("user2=%#v\n", user2)
    //	user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}
}
*/

// 5、结构体的继承（嵌套匿名结构体实现）

/*
Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}

*/
