package nacos

//
//import (
//	"fmt"
//
//	"github.com/nacos-group/nacos-sdk-go/model"
//
//	"github.com/nacos-group/nacos-sdk-go/clients"
//	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
//	"github.com/nacos-group/nacos-sdk-go/common/constant"
//	"github.com/nacos-group/nacos-sdk-go/vo"
//)
//
//var ClientServer naming_client.INamingClient
//
//func RegisterServer() {
//	serverConfigs := []constant.ServerConfig{
//		{
//			IpAddr: "localhost",
//			Port:   8848,
//		},
//	}
//	clientConfig := constant.ClientConfig{
//		NamespaceId:         "public",
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		LogDir:              "/tmp/nacos/log",
//		CacheDir:            "/tmp/nacos/cache",
//		LogLevel:            "debug",
//	}
//	client, err := clients.CreateNamingClient(map[string]interface{}{
//		"serverConfigs": serverConfigs,
//		"clientConfig":  clientConfig,
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	// 注册一个服务
//	serviceName := "myapp"
//	ip := "127.0.0.1"
//	port := 8080
//	instance := &model.Instance{
//		ServiceName: serviceName,
//		Ip:          ip,
//		Port:        uint64(port),
//	}
//	_, err = client.RegisterInstance(instance)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Service registered successfully")
//}
//
//func main() {
//	// 创建一个Nacos客户端
//	serverConfigs := []constant.ServerConfig{
//		{
//			IpAddr: "localhost",
//			Port:   8848,
//		},
//	}
//	clientConfig := constant.ClientConfig{
//		NamespaceId:         "public",
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		LogDir:              "/tmp/nacos/log",
//		CacheDir:            "/tmp/nacos/cache",
//		LogLevel:            "debug",
//	}
//	client, err := clients.CreateNamingClient(map[string]interface{}{
//		"serverConfigs": serverConfigs,
//		"clientConfig":  clientConfig,
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	// 发现一个服务
//	serviceName := "myapp"
//	instances, err := client.SelectInstances(model.SelectInstancesParam{
//		ServiceName: serviceName,
//		HealthyOnly: true,
//	})
//	if err != nil {
//		panic(err)
//	}
//	for _, instance := range instances {
//		fmt.Printf("Service instance: %s:%d\n", instance.Ip, instance.Port)
//	}
//}
//func Deregister() vo.DeregisterInstanceParam {
//	instance := vo.DeregisterInstanceParam{
//		Ip:          "10.2.171.84",
//		Port:        8084,
//		ServiceName: "user",
//		//Cluster:     "your_service_cluster",
//	}
//	return instance
//}
