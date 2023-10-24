// package main

// import "fmt"

// type Zhifubao struct {
// 	// 支付宝
// }

// // 支付宝的支付方法 Pay()
// func (z *Zhifubao) Pay(amount float64) {
// 	fmt.Printf("使用支付宝付款：%.2f 元 \n", amount)
// }

// // 结账函数
// func checkout(obj *Zhifubao) {
// 	obj.Pay(100)
// }

// func main() {
// 	checkout(&Zhifubao{})
// }

/*
在上面代码中，假设电商系统中最开始只设计了支付宝一种支付方式，
随着业务发展，需要添加微信支付

实际上我们并不关心用户选择的是什么支付方式，我们只关心调用Pay方法时能否正常运行，
在这种场景下我们没必要按照最开始只有支付宝的逻辑在创建一个Pay函数，然后再为其创建一个Checkout结账函数，太麻烦（如果以后有更多支付方式加入进来呢？）

我们可以将具体的支付方式抽象为一个Payer接口类型，即任何实现了Pay方法的都可以称为Payer类型，
此时只需要修改原始Checkout函数，它接收一个Payer类型参数，这样就能够在不修改原有函数调用的基础上支持新的支付方式
*/
package main

import "fmt"

type Zhifubao struct {
	// 支付宝
}

// 新增微信struct
type Weixin struct {
	// 微信
}

type Payer interface {
	Pay(float64)
}

// 支付宝的支付方法 Pay()
func (z *Zhifubao) Pay(amount float64) {
	fmt.Printf("使用支付宝付款：%.2f 元 \n", amount)
}

func (w *Weixin) Pay(amount float64) {
	fmt.Printf("使用微信付款：%.2f 元 \n", amount)
}

// Checkout 支付宝结账
// func CheckoutWithZFB(obj *ZhiFuBao) {
// 	// 支付100元
// 	obj.Pay(100)
// }

// // Checkout 微信支付结账
// func CheckoutWithWX(obj *WeChat) {
// 	// 支付100元
// 	obj.Pay(100)
// }

// 修改原有Checkout结账函数，
func Checkout(p Payer) {
	p.Pay(100)
}

func main() {
	Checkout(&Zhifubao{})
	Checkout(&Weixin{})
}
