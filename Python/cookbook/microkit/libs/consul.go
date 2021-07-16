package main

import "C"
import (
	"fmt"
	"log"
	"strings"
	"unsafe"

	consulApi "github.com/hashicorp/consul/api"
)

func main() {}

var consulMapping map[uintptr]*Consul = make(map[uintptr]*Consul, 10)

// Consul 实例
type Consul struct {
	client  *consulApi.Client
	address string
}

func (t *Consul) _register(serviceID string, serviceName string, serverAddress string, servicePort int, tags ...string) {
	reg := consulApi.AgentServiceRegistration{}
	reg.ID = serviceID
	reg.Name = serviceName
	reg.Address = serverAddress
	reg.Port = servicePort
	if len(tags) == 0 {
		reg.Tags = []string{"primary"}
	} else {
		reg.Tags = tags
	}

	check := consulApi.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = fmt.Sprintf("http://%s:%d/health", reg.Address, servicePort)

	reg.Check = &check
	err := t.client.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}

func (t *Consul) _deregister(serviceID string) {
	_ = t.client.Agent().ServiceDeregister(serviceID)
}

/**
export c part
*/

//export consulFactory
func consulFactory(_consulAddress *C.char) C.long {
	consulAddress := C.GoString(_consulAddress)

	config := consulApi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulApi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	consul := Consul{client: client, address: consulAddress}
	consulPtr := uintptr(unsafe.Pointer(&consul))
	consulMapping[consulPtr] = &consul
	return C.long(consulPtr)
}

//export consulRegister
func consulRegister(
	_consulPtr C.long,
	_serviceID, _serviceName,
	_serverAddress *C.char,
	_servicePort C.int,
	_tags *C.char,
) {
	consul := consulMapping[uintptr(_consulPtr)]
	serviceID := C.GoString(_serviceID)
	serviceName := C.GoString(_serviceName)
	serverAddress := C.GoString(_serverAddress)
	serverPort := int(_servicePort)

	var tags []string = make([]string, 0)
	for _, value := range strings.Split(C.GoString(_tags), ",") {
		tags = append(tags, strings.Trim(value, " "))
	}
	consul._register(serviceID, serviceName, serverAddress, serverPort, tags...)
}

//export consulDeregister
func consulDeregister(_consulPtr C.long, _serviceID *C.char) {
	consul := consulMapping[uintptr(_consulPtr)]
	serviceID := C.GoString(_serviceID)
	consul._deregister(serviceID)
}
