package consul

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

func SonSul(Address string, Port int, Add string) {
	var err error
	ConsulCli, err := api.NewClient(&api.Config{
		Address: Add,
	})
	if err != nil {
		return
	}
	Srvid := uuid.New().String()
	check := &api.AgentServiceCheck{
		Interval:                       "5s",
		Timeout:                        "5s",
		GRPC:                           fmt.Sprintf("%s:%d", Address, Port),
		DeregisterCriticalServiceAfter: "30s",
	}
	err = ConsulCli.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      Srvid,
		Name:    "user_srv",
		Tags:    []string{"GRPC"},
		Port:    Port,
		Address: Address,
		Check:   check,
	})
	if err != nil {
		return
	}
	return
}

func GetClient(serverName, Address string) (string, error) {
	cc, err := api.NewClient(&api.Config{
		Address: Address,
	})
	if err != nil {
		fmt.Printf("api.NewClient failed, err:%v\n", err)
		return "", err
	}
	serviceMap, date, err := cc.Agent().AgentHealthServiceByName(serverName)
	if serviceMap != "passing" {
		log.Println("获取consul服务发现失败！", err)
		return "", err
	}
	// 选一个服务机（这里选最后一个）
	var addr string
	for _, v := range date {
		addr = v.Service.Address + ":" + strconv.Itoa(v.Service.Port)
	}
	return addr, nil
}
