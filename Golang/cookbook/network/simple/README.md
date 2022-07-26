1、

如果一个耗时的handler，客户端关闭连接了，handler会停止吗

curl -i localhost:8000/hello

不会，因为golang无法终止正在运行的协程