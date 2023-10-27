package main

import (
	"fmt"
	"reflect"
)

// type Runner interface {
// 	Run()
// 	Cry()
// }

type Student struct {
	Name string
	Age  int
}

// func (stu Student) Run() {}
// func (stu Student) Cry() {}

func newStudent(name string, age int) *Student {
	return &Student{
		Name: name,
		Age:  age,
	}
}

// 使用Elem()方法reflect包反射的指针值
func reflectElemSetValue(i interface{}) interface{} {
	value := reflect.ValueOf(i)
	fmt.Println("value.Elem.Kind()值：", value.Elem().Kind())

	switch {
	case value.Kind() != reflect.Ptr:
		fmt.Println("error: 参数必须是指针")
		return nil
	case value.Elem().Kind() == reflect.Int:
		value.Elem().SetInt(200)
		return i
	default:
		return nil
	}
}

func main() {

	// 判断 Struct 是否实现了 interface
	// var _ Runner = &Student{}
	// var _ Runner = (*Student)(nil)

	newStudent("ops", 22)
	stu1 := Student{
		Name: "ops1",
		Age:  22,
	}
	fmt.Println(reflect.TypeOf(stu1).Name())
	fmt.Println(reflect.TypeOf(stu1).Kind())
	fmt.Println(reflect.TypeOf(stu1).NumField())

	for i := 0; i < reflect.TypeOf(stu1).NumField(); i++ {
		field := reflect.TypeOf(stu1).Field(i)
		fmt.Printf("name:%s index:%d type:%v \n", field.Name, field.Index, field.Type)
	}
	testNum := 100
	valueAfter := reflectElemSetValue(&testNum)
	fmt.Println(*valueAfter.(*int))
	// test
	// test1
}
