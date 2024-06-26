package main

import "fmt"

/*
	panic/recover 模式
	Recover捕获异常
	通常来说，不应该对panic异常做任何处理，但有时，也许我们可以从异常中恢复，至少我们可以在程序崩溃前，做一些操作。举个例子，当web服务器遇到不可预料的严重问题时，在崩溃前应该将所有的连接关闭；如果不做任何处理，会使得客户端一直处于等待状态。如果web服务器还在开发阶段，服务器甚至可以将异常信息反馈到客户端，帮助调试。

	如果在deferred函数中调用了内置函数recover，并且定义该defer语句的函数发生了panic异常，recover会使程序从panic中恢复，并返回panic value。导致panic异常的函数不会继续运行，但能正常返回。在未发生panic时调用recover，recover会返回nil。

	让我们以语言解析器为例，说明recover的使用场景。考虑到语言解析器的复杂性，即使某个语言解析器目前工作正常，也无法肯定它没有漏洞。因此，当某个异常出现时，我们不会选择让解析器崩溃，而是会将panic异常当作普通的解析错误，并附加额外信息提醒用户报告此错误。

	func Parse(input string) (s *Syntax, err error) {
		defer func() {
			if p := recover(); p != nil {
				err = fmt.Errorf("internal error: %v", p)
			}
		}()
		// ...parser...
	}
	deferred函数帮助Parse从panic中恢复。在deferred函数内部，panic value被附加到错误信息中；并用err变量接收错误信息，返回给调用者。我们也可以通过调用runtime.Stack往错误信息中添加完整的堆栈调用信息。

	不加区分的恢复所有的panic异常，不是可取的做法；因为在panic之后，无法保证包级变量的状态仍然和我们预期一致。比如，对数据结构的一次重要更新没有被完整完成、文件或者网络连接没有被关闭、获得的锁没有被释放。此外，如果写日志时产生的panic被不加区分的恢复，可能会导致漏洞被忽略。

	虽然把对panic的处理都集中在一个包下，有助于简化对复杂和不可以预料问题的处理，但作为被广泛遵守的规范，你不应该试图去恢复其他包引起的panic。公有的API应该将函数的运行失败作为error返回，而不是panic。同样的，你也不应该恢复一个由他人开发的函数引起的panic，比如说调用者传入的回调函数，因为你无法确保这样做是安全的。

	有时我们很难完全遵循规范，举个例子，net/http包中提供了一个web服务器，将收到的请求分发给用户提供的处理函数。很显然，我们不能因为某个处理函数引发的panic异常，杀掉整个进程；web服务器遇到处理函数导致的panic时会调用recover，输出堆栈信息，继续运行。这样的做法在实践中很便捷，但也会引起资源泄漏，或是因为recover操作，导致其他问题。

	基于以上原因，安全的做法是有选择性的recover。换句话说，只恢复应该被恢复的panic异常，此外，这些异常所占的比例应该尽可能的低。为了标识某个panic是否应该被恢复，我们可以将panic value设置成特殊类型。在recover时对panic value进行检查，如果发现panic value是特殊类型，就将这个panic作为errror处理，如果不是，则按照正常的panic进行处理（在下面的例子中，我们会看到这种方式）。

	下面的例子是title函数的变形，如果HTML页面包含多个<title>，该函数会给调用者返回一个错误（error）。在soleTitle内部处理时，如果检测到有多个<title>，会调用panic，阻止函数继续递归，并将特殊类型bailout作为panic的参数。

	// soleTitle returns the text of the first non-empty title element
	// in doc, and an error if there was not exactly one.
	func soleTitle(doc *html.Node) (title string, err error) {
		type bailout struct{}
		defer func() {
			switch p := recover(); p {
			case nil:       // no panic
			case bailout{}: // "expected" panic
				err = fmt.Errorf("multiple title elements")
			default:
				panic(p) // unexpected panic; carry on panicking
			}
		}()
		// Bail out of recursion if we find more than one nonempty title.
		forEachNode(doc, func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "title" &&
				n.FirstChild != nil {
				if title != "" {
					panic(bailout{}) // multiple titleelements
				}
				title = n.FirstChild.Data
			}
		}, nil)
		if title == "" {
			return "", fmt.Errorf("no title element")
		}
		return title, nil
	}
	在上例中，deferred函数调用recover，并检查panic value。当panic value是bailout{}类型时，deferred函数生成一个error返回给调用者。当panic value是其他non-nil值时，表示发生了未知的panic异常，deferred函数将调用panic函数并将当前的panic value作为参数传入；此时，等同于recover没有做任何操作。（请注意：在例子中，对可预期的错误采用了panic，这违反了之前的建议，我们在此只是想向读者演示这种机制。）

	有些情况下，我们无法恢复。某些致命错误会导致Go在运行时终止程序，如内存不足。
*/

// Go语言中(目前版本1.17.1)是没有异常机制的，但是使用 panic/recover 模式来处理错误
// panic 可以在任何地方引发，但是 recover 只在 defer 调用的函数中有效。

// func funcA() {
// 	fmt.Println("func A")
// }

// func funcB() {
// 	panic("panic in B")
// }

// func funcC() {
// 	fmt.Println("func C")
// }
// func main() {
// 	funcA()
// 	funcB()
// 	funcC()
// }

// 以上程序输出：
/*
func A
panic: panic in B

goroutine 1 [running]:
main.funcB(...)
        .../code/go_study/10_panicRecover/panicRecover.go:13
main.main()
        .../code/go_study/10_panicRecover/panicRecover.go:21 +0xa5
exit status 2
*/

// 程序运行期间 funcB 中引发了 panic 导致程序崩溃异常退出了。
// 这个时候，就可以通过 recover 将程序恢复回来，继续向后执行

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	// defer定义了一个匿名立即执行函数
	defer func() {
		err := recover()
		//如果程序出现了panic错误,可以通过recover进行恢复
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}

/*
	注意：
	1、recover()必须搭配defer使用。
	2、defer一定要在可能引发panic的语句之前定义。
*/
