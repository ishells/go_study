package main

import "fmt"

func main() {

	// var array_test = [3]int{0, 1, 2}
	// var array9 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 8, 20}
	// for seq, val := range array9 {
	// 	fmt.Println("%d,%v", seq, val)
	// }
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d   ", j, i, i*j)
		}
		fmt.Println()
	}
}
