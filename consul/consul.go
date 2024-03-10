package consul

import (
	"errors"
	"fmt"
	"github.com/15733012783/mysql/nacos"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

func SonSul(Ghost string, Host string, Port int, Name string) {
	var err error
	sprintf := fmt.Sprintf("%v:%v", nacos.GoodsT.Grpc.Host, 8500)
	ConsulCli, err := api.NewClient(&api.Config{
		Address: sprintf,
	})
	if err != nil {
		log.Println(err, "服务注册失败")
		return
	}
	Srvid := uuid.New().String()
	check := &api.AgentServiceCheck{
		Interval:                       "5s",
		Timeout:                        "5s",
		GRPC:                           fmt.Sprintf("%s:%d", Ghost, Port),
		DeregisterCriticalServiceAfter: "30s",
	}
	err = ConsulCli.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      Srvid,
		Name:    Name,
		Tags:    []string{"GRPC"},
		Port:    Port,
		Address: Host,
		Check:   check,
	})
	if err != nil {
		log.Println(err, "服务注册失败")
		return
	}
	return
}

var currentIndex int

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
		log.Println("获取consul服务发现失败***！", err)
		return "", err
	}
	// 选一个服务机（这里选最后一个）
	if len(date) == 0 {
		return "", errors.New("没有可用的服务")
	}
	// 获取当前要访问的服务的索引
	currentIndex = (currentIndex + 1) % len(date)

	// 获取当前要访问的服务地址
	selectedService := date[currentIndex]
	addr := selectedService.Service.Address + ":" + strconv.Itoa(selectedService.Service.Port)
	fmt.Println(addr, "addr*******************")
	return addr, nil
}
