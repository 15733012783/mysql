package model

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type T struct {
	Username  string `json:"Username"  yaml:"Username"`
	Password  string `json:"Password"  yaml:"Password"`
	Host      string `json:"Host"      yaml:"Host"`
	Port      string `json:"Port"      yaml:"Port"`
	Mysqlbase string `json:"Mysqlbase" yaml:"Mysqlbase"`
}

var NaCosT T

func NaCosConfig(Group, DataId string, Port int) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return
	}
	config, err3 := client.GetConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  Group,
	})
	if err3 != nil {
		return
	}
	json.Unmarshal([]byte(config), &NaCosT)
	yaml.Unmarshal([]byte(config), &NaCosT)
}

func updateDbConnection(config string) {
	s, _ := db.DB()
	if s != nil {
		s.Close()
	}
	db, err = gorm.Open(mysql.Open(config), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}

//func ListenConfig() {
//	//Listen config change,key=dataId+group+namespaceId.
//	err = client.ListenConfig(vo.ConfigParam{
//		DataId: "test-data",
//		Group:  "test-group",
//		OnChange: func(namespace, group, dataId, data string) {
//			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
//			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", NaCosT.Username,
//				NaCosT.Password, NaCosT.Host, NaCosT.Port, NaCosT.Mysqlbase)
//			updateDbConnection(dsn)
//		},
//	})
//}
