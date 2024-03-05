package consul

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"log"
)

var ConSuLClient *api.Client

func SonSul() {
	var err error
	ConSuLClient, err = api.NewClient(api.DefaultConfig())
	if err != nil {
		return
	}
	api.DefaultConfig().Address = fmt.Sprintf("%s:%d", "127.0.0.1", 8500)
	err = ConSuLClient.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.NewString(),
		Name:    "test",
		Tags:    []string{"GRPC"},
		Port:    8081,
		Address: "10.2.171.70",
	})

	check := &api.AgentServiceCheck{
		GRPC:     fmt.Sprintf("%s:%d", "10.2.171.70", 8081), // 这里一定是外部可以访问的地址
		Timeout:  "10s",                                     // 超时时间
		Interval: "10s",                                     // 运行检查的频率
		// 指定时间后自动注销不健康的服务节点
		// 最小超时时间为1分钟，收获不健康服务的进程每30秒运行一次，因此触发注销的时间可能略长于配置的超时时间。
		DeregisterCriticalServiceAfter: "1m",
	}
	srv := &api.AgentServiceRegistration{
		Name:    "test",                    // 服务名称
		Tags:    []string{"q1mi", "hello"}, // 为服务打标签
		Address: "10.2.171.70",
		Port:    8081,
		Check:   check,
	}
	if err != nil {
		zap.S().Panic(err.Error())
	}
	err = ConSuLClient.Agent().ServiceRegister(srv)
	if err != nil {
		panic(err)
		return
	}
}

func GetClient(serverName string) (string, int, error) {
	name, data, err := ConSuLClient.Agent().AgentHealthServiceByName(serverName)
	if name != "passing" {
		log.Println("获取consul服务发现失败！", err)
		return "", 0, nil
	}
	var Address string
	var Port int
	for _, val := range data {
		Address = val.Service.Address
		Port = val.Service.Port
	}
	log.Println("端口：lianjie", Address, Port)
	return Address, Port, nil
}
