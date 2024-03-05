package nacos

import (
	"log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var ClientServer naming_client.INamingClient

func RegisterServer() {
	// 配置Nacos服务器地址和命名空间等信息
	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}
	cc := &constant.ClientConfig{
		NamespaceId:         "", //
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}

	// 创建nacos客户端
	var err error
	ClientServer, err = clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	})
	if err != nil {
		log.Println("连接nacos配置客户端失败！", err)
	}
	suesscc, err := ClientServer.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "10.2.171.84",
		Port:        8084,
		Weight:      0,
		Enable:      false,
		Healthy:     true, // 开启健康检测
		Metadata:    nil,
		ClusterName: "",
		ServiceName: "user",
		GroupName:   "DEFAULT_GROUP",
		Ephemeral:   false,
	})
	if err != nil {
		log.Println("注册服务出错！", err)
		return
	}
	if suesscc {
		log.Println("nacos服务注册成功！", err)
		return
	} else {
		log.Println("nacos服务注册失败！", err)
		return
	}
}

//func HealthCheck(ClientServer naming_client.INamingClient) {
//	for {
//		time.Sleep(time.Second * 10)
//
//		serviceName := "user"
//		instances, err := ClientServer.SelectInstances(vo.SelectInstancesParam{
//			ServiceName: serviceName,
//			HealthyOnly: true,
//		})
//		if err != nil {
//			log.Printf("Error getting instances for health check: %v", err)
//			continue
//		}
//
//		for _, instance := range instances {
//			// 这里可以添加自定义的健康检测逻辑
//			// 例如，检查实例的状态，响应时间等
//			// 如果实例不健康，可以使用 DeregisterInstance 取消注册
//			fmt.Printf("Instance: %s, Healthy: %t\n", instance.Ip, instance.Healthy)
//		}
//	}
//}

func Deregister() vo.DeregisterInstanceParam {
	instance := vo.DeregisterInstanceParam{
		Ip:          "10.2.171.84",
		Port:        8084,
		ServiceName: "user",
		//Cluster:     "your_service_cluster",
	}
	return instance
}
