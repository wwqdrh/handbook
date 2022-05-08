# gin 实现 https 服务器

1、使用 openssl 自签证书

需要注意的是如果使用浏览器访问必须勾选信任该证书才能够继续使用

```shell
// 创建证书的步骤 ca-key.pem ca-req.csr ca-cert.pem server-key.pem server-req.csr server-cert.pem

// ca 证书颁发者
openssl genrsa -out cert/ca-key.pem 1024  // 创建私钥
openssl req -new -key cert/ca-key.pem -out cert/ca-req.csr -subj "/C=CN/ST=BJ/L=BJ/O=BJ/OU=BJ/CN=BJ"  // 创建csr证书请求
openssl x509 -req -in cert/ca-req.csr -out cert/ca-cert.pem -signkey cert/ca-key.pem -days 3650  // 生成crt证书

// 服务端私钥以及证书
openssl genrsa -out cert/server-key.pem 1024  // 创建服务端私钥
openssl req -new -out cert/server-req.csr -key cert/server-key.pem -subj "/C=CN/ST=BJ/L=BJ/O=BJ/OU=BJ/CN=BJ"
openssl x509 -req -in cert/server-req.csr -out cert/server-cert.pem -signkey cert/server-key.pem -CA cert/ca-cert.pem -CAkey cert/ca-key.pem -CAcreateserial -days 3650 // 生成crt证书

// 客户端私钥以及证书(用于双向验证)
openssl verify -CAfile cert/ca-cert.pem  cert/server-cert.pem // 确认证书
```

2、内存数据结构，字符串，判断字符串是否已经存在了，需要保证高并发情况下的数据安全

3、基于 gin 的 https 服务启动 以及 handler 的定义

接口定义

[post] /string/verify; multiform:arrs=[]

# 构建

docker build . -t wwqdrh/strverify:latest

docker run --rm -p 80:8080 --name strverify wwqdrh/strverify:latest

curl --location --request POST 'http://[集群中任意一个节点的地址]:31377/stringverify' \
--form 'arrs="2"' \
--form 'arrs="2"' \
--form 'arrs="2"'

## CICD

分成两部分讨论，CI，根据gitlab或者github的工作流workflow等自动编译生成镜像

当然在这种本地测试的环境中，这一部分完全可以由本地处理编译然后发布到镜像中

重点查看CD，持续部署，使用kubevela

### 基于jekins

向流水线中使用的 Git 仓库的分支推送代码变更，Git 仓库的 Webhook 会触发 Jenkins 中新创建的流水线。该流水线会自动构建代码镜像并推送至镜像仓库，然后对 KubeVela apiserver 发送 POST 请求，将仓库中的应用配置文件部署到 Kubernetes 集群中

### 基于gitops

相较于jekins，能够支持的功能更多

在 KubeVela Git 配置仓库以外，还需要准备一个应用代码仓库。在用户更新了应用代码仓库中的代码后，需要配置一个 CI 来自动构建镜像并推送至镜像仓库中。KubeVela 会监听镜像仓库中的最新镜像，并自动更新配置仓库中的镜像配置，最后再更新集群中的应用配置。使用户可以达成在更新代码后，集群中的配置也自动更新的效果。

CI: 也就是说需要两个仓库，一个接入CI自动构建镜像，这个可以使用act在本地自动构建并发布镜像到仓库中

CD: 另外一个仓库监听镜像的改变，然后基于kubevela进行部署