微服务相关，有些自有库是单独存放的，这里存放的是实现示例



[参考]: 

python实战grpc: https://www.jianshu.com/p/43fdfeb105ff?from=timeline&isappinstalled=0

protobuf3语言指南: https://blog.csdn.net/u011518120/article/details/54604615





grpc 使用 protobuf 进行数据传输. protobuf 是一种数据交换格式, 由三部分组成:



\- proto 文件: 使用的 proto 语法的文本文件, 用来定义数据格式

\- protoc: protobuf 编译器(compile), 将 proto 文件编译成不同语言的实现, 这样不同语言中的数据就可以和 protobuf 格式的数据进行交互

\- protobuf 运行时(runtime): protobuf 运行时所需要的库, 和 protoc 编译生成的代码进行交互



1、定义proto文件

2、使用grpcio-tools(python下的protoc编译器进行编译)

python -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I. hello.proto



--python_out=. : 编译生成处理 protobuf 相关的代码的路径, 这里生成到当前目录

--grpc_python_out=. : 编译生成处理 grpc 相关的代码的路径, 这里生成到当前目录

-I. helloworld.proto : proto 文件的路径, 这里的 proto 文件在当前目录



hello_pb2.py: 用来和protobuf数据进行交互

hello_pb2_grpc.py: 用来和grpc进行交互



\# grpc支持四种通信方式

\- 客服端一次请求, 服务器一次应答

\- 客服端一次请求, 服务器多次应答(流式)

\- 客服端多次请求(流式), 服务器一次应答

\- 客服端多次请求(流式), 服务器多次应答(流式)