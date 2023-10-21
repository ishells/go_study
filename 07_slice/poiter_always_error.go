package main

import "fmt"

/*
在循环中，`&u` 的指针地址是不变的，而 `u` 的值在每次迭代中都会被更新。
因此，在循环中我们实际上将用户结构体的指针追加到了 `newUser` 数组中，
但其实这些指针都指向最后一次迭代中的同一个结构体 `User{Name: "cc", Age: 1}`。

为了解决这个问题，可以在循环内部创建一个新的变量或利用切片截取的特性来避免指向同一地址进行追加。
以下是两种可行的写法：
1. 在循环内创建新变量
```
for _, u := range userList {
	    userCopy := u
		newUser = append(newUser, &userCopy)
}```
2. 利用切片截取
```
for i := range userList {
	    newUser = append(newUser, &userList[i])
}
*/

type User struct {
	Name string
	Age  int
}

func main() {
	userList := []User{
		User{Name: "aa", Age: 1},
		User{Name: "bb", Age: 1},
		User{Name: "cc", Age: 1},
	}

	// fmt.Println(&userList)
	var newUser []*User
	for i, u := range userList {
		fmt.Println("u: ", u)
		fmt.Println("&u: ", &u)
		newUser = append(newUser, &u)
		// newUser = append(newUser, &userList[i])
		fmt.Println(newUser)
	}
	/*
		在每次循环迭代中，u 被赋予 userList 中的一个元素。u 是一个局部变量，其作用域仅在循环内部。

		你使用 append(newUser, &u) 来将 u 的地址添加到 newUser 切片。append 函数的工作方式是，
		它会将指定的元素添加到切片的末尾，并返回一个新的切片。但由于 newUser 是切片的引用，
		你需要将新的切片分配回给 newUser，否则新切片的更改不会保存。
		在这里，你并没有将新的切片分配回给 newUser，因此 newUser 保持不变。

		在下一次迭代中，u 被赋予 userList 中的下一个元素，而现有的 newUser 切片保持不变。
		这是因为你没有将新的切片分配回给 newUser，所以它仍然引用上一次迭代中的切片。

		这导致了问题：所有的 &u 最终都指向相同的 u，因为在 append 操作中没有创建新的切片。

		解决这个问题的方法是，在循循环内部创建一个新变量 userCopy 并将其地址添加到 newUser，
		或者直接将 userList 中的元素的指针添加到 newUser，而不是直接使用 &u。这将确保 newUser 包含不同的 User 结构体的指针。

	*/

	// 第一次：cc
	// 第二次：cc
	for _, nu := range newUser {
		fmt.Println("%+v", nu.Name)
	}
}
