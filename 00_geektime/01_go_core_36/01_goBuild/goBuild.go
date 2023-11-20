package main

// https://github.com/hyper0x/go_command_tutorial/blob/master/0.3.md

/*

go build命令一些可选项的用途和用法
	在运行go build命令的时候，默认不会编译目标代码包所依赖的那些代码包。当然，如果被依赖的代码包的归档文件不存在，或者源码文件有了变化，那它还是会被编译。

	如果要强制编译它们，可以在执行命令的时候加入标记-a。此时，不但目标代码包总是会被编译，它依赖的代码包也总会被编译，即使依赖的是标准库中的代码包也是如此。

	另外，如果不但要编译依赖的代码包，还要安装它们的归档文件，那么可以加入标记-i。

	那么我们怎么确定哪些代码包被编译了呢？有两种方法。

	运行go build命令时加入标记-x，这样可以看到go build命令具体都执行了哪些操作。另外也可以加入标记-n，这样可以只查看具体操作而不执行它们。
	运行go build命令时加入标记-v，这样可以看到go build命令编译的代码包的名称。它在与-a标记搭配使用时很有用。
	下面再说一说与Go源码的安装联系很紧密的一个命令：go get。

	命令go get会自动从一些主流公用代码仓库（比如GitHub）下载目标代码包，并把它们安装到环境变量GOPATH包含的第1工作区的相应目录中。如果存在环境变量GOBIN，那么仅包含命令源码文件的代码包会被安装到GOBIN指向的那个目录。

	最常用的几个标记有下面几种。

	-u：下载并安装代码包，不论工作区中是否已存在它们。
	-d：只下载代码包，不安装代码包。
	-fix：在下载代码包后先运行一个用于根据当前Go语言版本修正代码的工具，然后再安装代码包。
	-t：同时下载测试所需的代码包。
	-insecure：允许通过非安全的网络协议下载和安装代码包。HTTP就是这样的协议。

	$ go get -d github.com/hyper-carrot/go_lib/logging
	现在，让我们再来看一下Lib工作区的目录结构：
	/home/hc/golang/lib:
		bin/
		pkg/
		src/
			github.com/
			hyper-carrot/
			go_lib/
				logging/
		...
	我们可以看到，go get命令只将代码包下载到了Lib工作区的src目录，而没有进行后续的编译和安装动作。这个加入-d标记的结果。

	再来看-fix标记。我们知道，绝大多数计算机编程语言在进行升级和演进过程中，不可能保证100%的向后兼容（Backward Compatibility）。
	在计算机世界中，向后兼容是指在一个程序或者代码库在更新到较新的版本后，用旧的版本程序创建的软件和系统仍能被正常操作或使用，或在旧版本的代码库的基础上编写的程序仍能正常编译运行的能力。Go语言的开发者们已想到了这点，并提供了官方的代码升级工具——fix。fix工具可以修复因Go语言规范变更而造成的语法级别的错误。

	标记-u的意图和执行的动作都比较简单。我们在执行go get命令时加入-u标记就意味着，如果在本地工作区中已存在相关的代码包，那么就是用对应的代码版本控制系统的更新命令更新它，并进行编译和安装。这相当于强行更新指定的代码包及其依赖包。


*/

/*

自定义代码包远程导入路径

如果你想把你编写的（被托管在不同的代码托管网站上的）代码包的远程导入路径统一起来，或者不希望让你的代码包中夹杂某个代码托管网站的域名，那么你可以选择自定义你的代码包远程导入路径。这种自定义的实现手段叫做“导入注释”。导入注释的写法示例如下：

package analyzer // import "hypermind.cn/talon/analyzer"
代码包analyzer实际上属于我的一个网络爬虫项目。这个项目的代码被托管在了Github网站上。它的网址是：https://github.com/hyper-carrot/talon。如果用标准的导入路径来下载analyzer代码包的话，命令应该这样写go get github.com/hyper-carrot/talon/analyzer。不过，如果我们像上面的示例那样在该代码包中的一个源码文件中加入导入注释的话，这样下载它就行不通了。我们来看一看这个导入注释。

导入注释的写法如同一条代码包导入语句。不同的是，它出现在了单行注释符//的右边，因此Go语言编译器会忽略掉它。另外，它必须出现在源码文件的第一行语句（也就是代码包声明语句）的右边。只有符合上述这两个位置条件的导入注释才是有效的。再来看其中的引号部分。被双引号包裹的应该是一个符合导入路径语法规则的字符串。其中，hypermind.cn是我自己的一个域名。实际上，这也是用来替换掉我想隐去的代码托管网站域名及部分路径（这里是github.com/hyper-carrot）的那部分。在hypermind.cn右边的依次是我的项目的名称以及要下载的那个代码包的相对路径。这些与其标准导入路径中的内容都是一致的。为了清晰起见，我们再来做下对比。

github.com/hyper-carrot/talon/analyzer // 标准的导入路径
hypermind.cn           /talon/analyzer // 导入注释中的导入路径
你想用你自己的域名替换掉标准导入路径中的哪部分由你自己说了算。不过一般情况下，被替换的部分包括代码托管网站的域名以及你在那里的用户ID就可以了。这足以达到我们最开始说的那两个目的。

虽然我们在talon项目中的所有代码包中都加入了类似的导入注释，但是我们依然无法通过go get hypermind.cn/talon/analyzer命令来下载这个代码包。因为域名hypermind.cn所指向的网站并没有加入相应的处理逻辑。具体的实现步骤应该是这样的：

编写一个可处理HTTP请求的程序。这里无所谓用什么编程语言去实现。当然，我推荐你用Go语言去做。

将这个处理程序与hypermind.cn/talon这个路径关联在一起，并总是在作为响应的HTML文档的头中写入下面这行内容：

<meta name="go-import" content="hypermind.cn/talon git https://github.com/hyper-carrot/talon">
hypermind.cn/talon/analyzer熟悉HTML的读者都应该知道，这行内容会被视为HTML文档的元数据。它实际上go get命令的文档中要求的写法。它的模式是这样的：

<meta name="go-import" content="import-prefix vcs repo-root">
实际上，content属性中的import-prefix的位置上应该填入我们自定义的远程代码包导入路径的前缀。这个前缀应该与我们的处理程序关联的那个路径相一致。而vcs显然应该代表与版本控制系统有关的标识。还记得表0-2中的主命令列吗？这里的填入内容就应该该列中的某一项。在这里，由于talon项目使用的是Git，所以这里应该填入git。至于repo-root，它应该是与该处理程序关联的路径对应的Github网站的URL。在这里，这个路径是hypermind.cn/talon，那么这个URL就应该是https://github.com/hyper-carrot/talon。后者也是talon项目的实际网址。

好了，在我们做好上述处理程序之后，go get hypermind.cn/talon/analyzer命令的执行结果就会是正确的。analyzer代码包及其依赖包中的代码会被下载到GOPATH环境变量中的第一个工作区目录的src子目录中，然后被编译并安装。

注意，具体的代码包源码存放路径会是/home/hc/golang/lib/src/hypermind.cn/talon/analyzer。也就是说，存放路径（包括代码包源码文件以及相应的归档文件的存放路径）会遵循导入注释中的路径（这里是hypermind.cn/talon/analyzer），而不是原始的导入路径（这里是github.com/hyper-carrot/talon/analyzer）。另外，我们只需在talon项目的每个代码包中的某一个源码文件中加入导入注释，但这些导入注释中的路径都必须是一致的。在这之后，我们就只能使用hypermind.cn/talon/作为talon项目中的代码包的导入路径前缀了。一个反例如下：

hc@ubt:~$ go get github.com/hyper-carrot/talon/analyzer
package github.com/hyper-carrot/talon/analyzer: code in directory /home/hc/golang/lib/src/github.com/hyper-carrot/talon/analyzer expects import "hypermind.cn/talon/analyzer"
与自定义的代码包远程导入路径有关的内容我们就介绍到这里。从中我们也可以看出，Go语言为了让使用者的项目与代码托管网站隔离所作出的努力。只要你有自己的网站和一个不错的域名，这就很容易搞定并且非常值得。这会在你的代码包的使用者面前强化你的品牌，而不是某个代码托管网站的。当然，使你的代码包导入路径整齐划一是最直接的好处。

*/

/*
命令源码文件的用途是什么，怎样编写它？

这里，我给出你一个参考的回答：命令源码文件是程序的运行入口，是每个可独立运行的程序必须拥有的。我们可以通过构建或安装，生成与其对应的可执行文件，后者一般会与该命令源码文件的直接父目录同名。

如果一个源码文件声明属于main包，并且包含一个无参数声明且无结果声明的main函数，那么它就是命令源码文件。 就像下面这段代码：

package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

当需要模块化编程时，我们往往会将代码拆分到多个文件，甚至拆分到不同的代码包中。
但无论怎样，对于一个独立的程序来说，命令源码文件永远只会也只能有一个。如果有与命令源码文件同包的源码文件，那么它们也应该声明属于main包。

*/
// 使用 flag 库为命令源文件设置参数
import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, name, "default name", "There needs a name string")
}

func main() {
	flag.Parse()
	fmt.Printf("HELLO,%s", name)
}
