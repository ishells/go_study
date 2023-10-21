package main

import "fmt"

// 方法是作用于特定类型的函数（即，方法method仅可以被特定类型所调用）

/*
	方法与函数的区别是:
		函数可以被随便调用，
		而方法规定了只有特定的接受者(Receiver)才能调用

	方法的语法如下：
		func (r ReceiverType) funcName(parameters)(return parameters){

		}
		当调用method时，会将receiver作为函数的第一个参数，相当于


		func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    		函数体
		}
		接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。
			例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
		接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
		方法名、参数列表、返回参数：具体格式与函数定义相同。

	值接收者和指针接收者区别：
		所以，receiver是值类型还是指针类型要看method的作用。
		如果要修改对象的值，就需要传递对象的指针。指针作为Receiver会对实例对象的内容发生操作
		而普通类型作为Receiver仅仅是以副本作为操作对象,并不对原实例对象发生操作。

	什么时候应该使用指针类型接收者：
		1、需要修改接收者中的值
		2、接收者是拷贝代价比较大的大对象（比如struct中有很多变量值，如果不适用指针比较消耗资源）
		3、保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者
*/

type dog struct {
	name string
	age  int
}

// dog struct的构造函数
func newDog(name string, age int) dog {
	return dog{
		name: name,
		age:  age,
	}
}

// 这里的方法wang()仅作用于特定类型dog（因为dog类型才有name值）
// 接收者代表调用该方法的具体类型变量，多用类型名的首字母小写来表示
func (d dog) wang() {
	fmt.Printf("%s在汪汪汪\n", d.name)
}

// 使用值接收者：因为传递的是普通类型，所以只是对dog的copy，不会影响原值
func (d dog) ageCopyAdd1() {
	d.age++
}

// 指针接收者：因为传递的是指针类型，所以此时修改的就是源对象的值
func (d *dog) ageRealAdd1() {
	d.age++
}

func main() {
	dog1 := newDog("小白", 6)
	dog1.wang()
	fmt.Println(dog1.age)
	dog1.ageCopyAdd1()
	fmt.Println(dog1.age)
	dog1.ageRealAdd1()
	fmt.Println(dog1.age)

}

// 自定义类型加方法
// 不能给别的包里边的类型加方法，只能给自己包里的类型添加方法

/*
package main
import "fmt"

type myInt int

func (m myInt) hello(){
	fmt.Println("我是一个int")
}

func main(){
	m := myInt(100)
	m.hello()
}

*/
