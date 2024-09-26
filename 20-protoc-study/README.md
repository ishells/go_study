`Message`
- *message*: protobuf中定义一个消息类型式是通过关键字message字段指定的。 
  - 消息就是需要传输的数据格式的定义
    - message关键字类似于C++中的class,JAVA中的class,go中的struct
    - 在消息中承载的数据分别对应于每一个字段,其中每个字段都有一个名字和一种类型
    - 一个proto文件中可以定义多个消息类型

`字段规则`
- *required*: 消息体中必填字段,不设置会导致编码异常。在protobuf2中使用,在Eprotobuf3中被删去
- *optional*: 消息体中可选字段。protobuf3没有了required,optional等说明关键字,都默认为optional
- *repeate*: 消息体中可重复字段,重复的值的顺序会被保留在go中重复的会被定义为切片（即如果想在go中声明使用一个切片，在proto文件中就可以使用repeate关键字声明一个变量）
    ```golang
    # 例
    message HelloRequest{
        string requestName=1;
        // int64 age = 2;
        repeated int64 age = 2;
    }
    # 对应生成的golang代码中的
    type HelloRequest struct {
        RequestName string `protobuf:"bytes,1,opt,name=requestName,proto3" json:"requestName,omitempty"`
        Age []int64 `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
    }
    ```

`消息号`
- 在消息体的定义中,每个字段都必须要有一个唯一的标识号,标记只号是(1,2^29-1]范围内的一个整数。

`嵌套消息`
- 可以在其他消息类型中定义、使用消息类型,在下面的例子中,person消息就定义在Personinfo消息内如
    ```golang
    message PersonInfo {
        message Person {
            string name = 1;
            int32 height = 2;
            repeated int32 age = 3;
        }
        repeated Person person = 1;
    }
    ```

- 如果要在它的父消息类型的外部重用这个消息类型，需要PersonInfo.Person的形式使用它，如：
    ```golang
    message PersonMessage {
        PersonInfo.Person person = 1;
    }
    ```

`服务定义`
- 如果想要将消息类型用在RPC系统中,可以在**.proto文件中定义一个RPC服务接口*,protocol buffer编译器将会根据所选择的不同语言生成服务接口代码及存根。
    ```golang
    service SearchService {
        # rpc 服务函数名 (参数) 返回 (返回参数)
        rpc Search (SearchRequest) returns (SearchResponse)
    }
    ```

`认证-安全传输`
- gRPC是一个典型的C/S模型,需要开发客户端和服务端,客户端与服多端需要达成协议,使用某一个确认的传输协议来传输数据,gRPC通常默认是使用protobuf来作为传输协议,当然也是可以使用其他自定义的。
    ```
    client  编译并转化为代码    <---各种传输协议默认是protobuf，也可以是其他的传输等--->    编译并转化成代码    Server 
    ```
- 那么,客户端与服务端要通信之前,客户端如何知道自己的数据是发给哪一个明确的服务端呢?反过来,服务端是不是也需要有一种方式来弄清楚自己的数据要返回给谁呢?
  
- 那么就不得不提gRPC的认证，此处说到的认证,不是用户的身份认证,而是指多个server和多个client之间,如何识别对方是谁,并且可以安全的进行数据传输
  - SSL/TLS认证方式(采用http2协议)
  - 基于Token的认证方式(基于安全连接)
  - 不采用任何措施的连接,这是不安全的连接(默认采用http1)
  - 自定义的身份认证
- 客户端和服务端之间调用,我们可以通过加入证书的方式,实现调用的安全性
- TLS(Transport Layer Security,安全传输层TCP协议之」上的协议,服务于应用层,它的前身是SSL(Secure Socket Layer, 安全套接字层),它实现了将应用层的报文进行加密后再交由TCP进行行转输的功能。
- TLS协议主要解决如下三个网络安全问题。
  - 保密(message privacy),保密通过加密encryption实现,所有信息都加密传输,第三方无法嗅探;
  - 完整性(message integrity),通过MAC校验机制,一旦被篡改,通信双双方会立刻发现;
  - 认证(mutual authentication),双方认证,双方都可以配备证书,防止份被冒充;
  
  `生产环境可以购买证书或者使用一些平台发放的免费证书`
  - key: 服务器上的私钥文件,用于对发送给客户端数据的加密,以及对从客户端接收到数据的解密。
  - csr: 证书签名请求文件,用于提交给证书颁发机构(CA)对证书签名。
  - crt: 由证书颁发机构(CA)签名后的证书,或者是开发者自签名的证书,包含证书持有人的信息,持有人的公钥,以及签署者的签名等信息
  - pem: 是基于Base64编码的证书格式,扩展名包括PEM、CRT和CER.

`token认证`
```golang
// gRPC提供我们的一个接口,这个接口中有两个方法,接口位于credentials包下,这个接口需要客户端来实现
type PerRPCCredentials interface {
    GetRequestMetadata(ctx context.Context, uri ...string)(map[string]string, error)
    RequireTransportSecurity() bool
}
```
- 第一个方法作用是获取元数据信息,也就是客户端提供的key,value对,context用于控制超时和取消,uri是请求入口处的ur
- 第二个方法的作用是否需要基于TLS认证进行安全传输,如果返回值是true,则必须加上TLS验证,返回值是false则不用。

`gRPC将各种认证方式浓缩统一到一个凭证(credentials)上,可以单独使用一种凭证,比如只使用TLS凭证或者只使用自定义凭证,也可以多
种凭证组合,gRPC提供统一的API验证机制,使研发人员使用方便,这这也是gRPC设计的巧妙之处`


`流程`
1. *编写proto文件*

2. *使用protoc编译器生成对应代码*
   1. protoc编译器将proto文件编译成golang代码,生成xxxx.pb.go文件
   2. protoc编译器将proto文件编译成grpc代码,生成xxxx_grpc.pb.go文件

3. *服务端编写*
   1. 创建gRPCServer对象,你可以理解为它是Server端的抽象对象
   2. 将server(其包含需要被调用的服务端接口)注册到gRPCServer的内部注册中心。这样可以在接受到请求时,通过内部的服务发现,发现该服务端接口转接进行逻辑处理
   3. 创建Listen,监听TCP端口
   4. gRPCServer开始lis.Accept,直到Stop
4. *客户端编写*
   1. 创建与给定目标(服务端)的连接交互
   2. 创建server的客户端对象
   3. 发送RPC请求,等待同步响应,得到回调后返回响应结果
   4. 输出响应结果

`xxx.pb.go 和 xxx_grpc.pb.go 文件的区别`
- `xxx.pb.go`: 这个文件包含了由 `.proto` 文件定义的数据结构（消息体）的 Go 语言版本以及与这些结构相关的序列化和反序列化代码。它由 `protoc` 编译器使用 `protoc-gen-go` 插件生成。
- `xxx_grpc.pb.go`: 这个文件包含了 GRPC 服务接口的 Go 语言实现桩。如果 `.proto` 文件中定义了服务 (`service`)，那么 `xxx_grpc.pb.go` 文件中将包含这些服务接口的客户端和服务器端代码。这需要使用 `protoc-gen-go-grpc` 插件来生成。

`跨语言的 grpc 拷贝服务端代码到其他语言的目录下的做法`
- 对于跨语言使用 GRPC 通信，比如服务端是 Go 语言，而客户端可能是 Python、Java 或其他支持的语言。每种语言的实现只需保证有一致的 `.proto` 文件。首先确保每种语言都安装了 `protoc` 并相应语言的插件。然后根据每个语言的构建系统和目录结构，分别生成所需的代码。
- 服务端 (Go)
    ```bash
    protoc --go_out=. --go-grpc_out=. myservice.proto
    ```
- 客户端 (Python)
    ```bash
    protoc --python_out=. --grpc_python_out=. myservice.proto
    ```
- 客户端 (Java)
    ```bash
    protoc --java_out=. --grpc_java_out=. myservice.proto
    ```