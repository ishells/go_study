package main

import "fmt"

// 反转切片 /
func reverse(s []int) []int {
	l := len(s)
	newS := make([]int, l)

	for k, v := range s {
		newS[l-k-1] = v
	}
	return newS
}

func main() {
	fmt.Println(reverse([]int{1, 2, 3, 4, 5, 6}))
}
