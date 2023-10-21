package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// This a commit change test
func main() {
	// map 是一种无序的基于key-value的数据结构，go语言中的map是引用类型，必须初始化才能使用

	// 1、map 定义：
	// map[keyType]ValueType ,KeyType代表键的类型，ValueType代表值对应的值类型

	// map类型的变量默认初始值为nil，需要使用make()函数来分配内存。
	// 语法为:
	// make(map[keyType]ValueType, [cap])		cap即为容量大小,该参数非必须，但是最好在初始化map的时候就为其指定一个合适的容量

	var map1 map[string]int
	fmt.Println(map1 == nil) // map1还没有进行初始化（没有在内存空间中开辟空间）
	map1 = make(map[string]int, 10)
	fmt.Println(map1 == nil) // 使用make分配完内存之后map1就不是nil了
	map1["age"] = 23
	map1["height"] = 226
	fmt.Println(map1)
	fmt.Println(map1["age"])
	fmt.Println(map1["weigh"]) // 如果不存在这个key，拿到对应数值类型的零值
	// 约定成俗，将map指定key的值赋给两个变量时，第一个变量存的是该key对应value值，第二个变量用以判断该key是否为空
	value, ok := map1["age"]
	if !ok {
		fmt.Println("map1中不存在此键值")
	} else {
		// fmt.Println(value)
		fmt.Printf("map1[%v]的值是%v", value, map1["age"])
	}

	// 2、map 的遍历

	for k, v := range map1 {
		fmt.Println(k, v)
	}
	// 只遍历key值
	for k := range map1 {
		fmt.Println(k)
	}
	// 只遍历value值
	for _, v := range map1 {
		fmt.Println(v)
	}

	// 3、使用delete()方法删除
	// delete(mapName,keyName)
	// 内建函数delete按照指定的键将元素从映射中删除，若map为空或无此元素，delete不进行操作
	delete(map1, "height")
	fmt.Println(map1)

	// 4、按照指定顺序遍历map
	/*
		① 首先声明一个map容量200，使用循环对map初始化赋值，key依次增加，value来自随机数
		② 将map的所有key值取出来存到切片中
		③ 对切片进行排序
		④ 根据排序后的key值获取value值
	*/
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var map_keys = make([]string, 0, 200)
	for key := range scoreMap {
		map_keys = append(map_keys, key)
	}
	//对切片进行排序
	sort.Strings(map_keys)
	//按照排序后的key遍历map
	for _, key := range map_keys {
		fmt.Println(key, scoreMap[key])
	}

	// 5、元素为map类型的切片
	// 声明元素类型为map的切片时，切片和map都需要进行初始化

	// 使用make创建切片格式 make([]Type,length,cap)
	// 使用make()创建map时格式为 make(map[key_type]value_type,[cap])
	var slice1 = make([]map[int]string, 10, 10) // 使用make创建元素类型为map的切片

	slice1[0] = make(map[int]string, 1) // 初始化slice中的第一个map元素
	slice1[0][0] = "beijing"
	fmt.Println(slice1)

	// 6、值为切片类型的map

	// 使用make创建map格式： make(type,cap)

	var map2 = make(map[string][]int, 10) // map[string]Type，map的key类型是string，value类型是切片[]int
	map2["beijing"] = []int{10, 20, 30}
	fmt.Println(map2)
}

/* 例：
make创建map：	make(type,cap)  ->  make(map1[string]int,200)
make创建slice：	make([]T, size ,cap)  ->  make([]int, 5, 10)
make创建元素为map类型的slice：	make([]map[int]string, 5, 10) slice格式是[]type，此例Type是键为int，value为string的map
make创建元素为slice类型的map：	make(map[string][]int, 10)  map[string]Type , 此例中Type是int类型的slice
声明元素类型为map的切片时，切片和map都需要分别初始化
*/
