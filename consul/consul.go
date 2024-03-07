package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
)

var ConSuLClient *api.Client

func SonSul() {
	var err error
	api.DefaultConfig().Address = fmt.Sprintf("%s:%d", "10.2.171.70", 8500)
	ConSuLClient, err = api.NewClient(api.DefaultConfig())
	if err != nil {
		zap.S().Panic(err.Error())
	}

	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", "10.2.171.70", 8081),
		Timeout:                        "10s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "1m",
	}
	srv := &api.AgentServiceRegistration{
		Name:    "test",
		Tags:    []string{"q1mi", "hello"},
		Address: "10.2.171.70",
		Port:    8081,
		Check:   check,
	}
	err = ConSuLClient.Agent().ServiceRegister(srv)
	if err != nil {
		panic(err)
		return
	}
}

func GetClient(serverName string) (*grpc.ClientConn, error) {
	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("api.NewClient failed, err:%v\n", err)
		return nil, err
	}

	serviceMap, date, err := cc.Agent().AgentHealthServiceByName(serverName)
	if serviceMap != "passing" {
		log.Println("获取consul服务发现失败！", err)
		return nil, err
	}
	// 选一个服务机（这里选最后一个）
	var addr string
	for _, v := range date {
		addr = v.Service.Address + ":" + strconv.Itoa(v.Service.Port)
	}
	fmt.Println(addr)
	// 建立RPC连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return nil, err
	}
	return conn, err
}
