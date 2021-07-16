package util

import (
	"fmt"
	"log"

	consulApi "github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
)

// go consul 客户端库 githug.com/hashicorp/consul

// consul服务中心，用来注册服务监听健康值等等
// 使用docker启动
// docker run -d -name=consul1.8.6 -p 8500:8500 consul:1.8.6 agent -server -bootstrap -ui -client 0.0.0.0
// -server 代表以服务端的方式启动
// -bootstrap 指定自己为leader，不需要选举
// -ui 启动一个内置管理web界面
// -client 指定客户端可以访问的IP，设置为0.0.0,0则任意访问，否则默认主机可以访问

// 手动注册
// curl --request PUT --data @userservice.json localhost:8500/v1/agent/service/register

// 手动注销服务
// curl --request PUT http://localhost:8500/v1/agent/service/deregister/userservice

// 服务注册、 服务发现

var ConsulClient *consulApi.Client
var ServiceID string
var ServiceName string
var ServicePort int

func init() { // 引入包的时候自动执行
	config := consulApi.DefaultConfig()
	config.Address = "172.27.104.48:8500"
	client, err := consulApi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	ConsulClient = client
	ServiceID = "userservice" + uuid.NewV1().String()
}

func SetServiceNameAndPort(name string, port int) {
	ServiceName = name
	ServicePort = port
}

func RegService() {
	reg := consulApi.AgentServiceRegistration{}
	reg.ID = ServiceID
	reg.Name = ServiceName
	reg.Address = "172.27.104.48"
	reg.Port = ServicePort
	reg.Tags = []string{"primary"}

	check := consulApi.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = fmt.Sprintf("http://%s:%d/health", reg.Address, ServicePort)

	reg.Check = &check
	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}

func UnRegService() {
	_ = ConsulClient.Agent().ServiceDeregister(ServiceID)
}
