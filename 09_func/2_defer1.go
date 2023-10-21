package main

import "fmt"

func calculation(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calculation("AA", x, calculation("A", x, y))
	x = 10
	defer calculation("BB", x, calculation("B", x, y))
	y = 20
}

/*
	1. x:=1
	2. y:=2
	3. defer calculation("AA",1,calculation("A",1,2))
	4. calculation("A",1,2)  // print("A",1,2,3),return 3
	5. defer calculation("AA",1,3)
	6. x = 10
	7. defer calculation("BB",10,calculation("B",10,2))
	8. calculation("B",10,2) // print("B",10,2,12),return 12
	9. defer calculation("BB",10,12)
	10. y = 20
	11. 函数体执行完毕，逆序执行defer语句
	12. calculation("BB",10,12)  // print("BB",10,12,22)
	13. calculation("AA",1,3,4)  // print("AA",1,3,4)

	结果输出：
	A 1 2 3
	B 10 2 12
	BB 10 12 22
	AA 1 3 4
*/
