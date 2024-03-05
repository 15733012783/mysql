package consul

import (
	"fmt"
	"github.com/google/uuid"
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

func GetClient(serverName string) (*grpc.ClientConn, error) {
	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("api.NewClient failed, err:%v\n", err)
		return nil, err
	}
	// 返回的是一个 map[string]*api.AgentService
	// 其中key是服务ID，值是注册的服务信息
	serviceMap, err := cc.Agent().ServicesWithFilter("Service==`hello`")
	if err != nil {
		fmt.Printf("query service from consul failed, err:%v\n", err)
		return nil, err
	}
	// 选一个服务机（这里选最后一个）
	var addr string
	for k, v := range serviceMap {
		fmt.Printf("%s:%#v\n", k, v)
		addr = v.Address + ":" + strconv.Itoa(v.Port)
	}

	// 建立RPC连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return nil, err
	}
	return conn, err
}
