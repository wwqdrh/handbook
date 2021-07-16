## 步骤

1、使用 openssl 自签证书

2、编写 https 服务端 api 实现，实现数组数据的编解码处理

3、编写客户端代码

4、编写单元测试

## 启动测试

`go test -v ./server/...`

包含两个测试用例

- TestHttpsServerClient
- TestHttpsServerClientWithClear (清除服务端缓存)

![结果](./result.png)

## log.1

1、不能包含两个包名: 移除main包
2、server 中存在竞态代码: sync.RWMutex 使用读写锁处理
3、测试中应该固定 server 和 client 的启动顺序:
    - 1)、服务端测试时先判断端口是否已经启用?
    - 2)、添加重试机制，如果失败等待1s，3次机会

## log.2

1、修改所有未经过处理error的代码
