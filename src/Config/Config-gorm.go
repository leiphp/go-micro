package Config

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"os"
	"time"
)

type DataConfig struct {
	Mysql *struct{
		Dsn 	string
		Maxidle int
		Maxopen int
	}
	Redis *struct{
		Ip 		string
		Port 	int
	}
}

type GlobalConfig struct {
	Config *struct{
		Address string
		Path 	string
		Port 	uint64
	}
	Service *struct{
		Namespace 	string
		Name 		string
	}
	Data *DataConfig
}

var JConfig *GlobalConfig
var nacosClient config_client.IConfigClient

//初始化主配置文件(本地配置)和初始化nacos链接
func InitConfig() {
	configFile := "app.yaml"
	err := config.LoadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	JConfig = &GlobalConfig{Data:new(DataConfig)}
	err = config.Get("100txy").Scan(JConfig)
	if err != nil {
		log.Fatal(err)
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:JConfig.Config.Address,
			ContextPath:JConfig.Config.Path,
			Port:JConfig.Config.Port,
		},
	}
	nacosClient, err = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})
	if err != nil {
		log.Fatal("nacos init err",err)
	}
	//开始加重数据相关配置
	//JConfig.Data.Mysql = new(MySqlConfig)
	//这里需要跟nacos配置Data Id和Group保持一致
	listenNacos("txy-sysconfig-mysql","DEFAULT_GROUP",JConfig.Data)
}

func listenNacos(dataid string, group string, model interface{}){
	err := nacosClient.ListenConfig(vo.ConfigParam{
		DataId:   dataid,
		Group:    group,
		Content:  "",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("data:",data)
			time.Sleep(time.Second*1)//模拟卡顿3秒
			cacheFile := fmt.Sprintf("./runtime/configcache/%s-%s.yaml",group,dataid)
			file,err := os.OpenFile(cacheFile,os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
			if err != nil {
				log.Println("err1",err)
				return
			}
			defer file.Close()
			_, err = file.WriteString(data)
			if err != nil {
				log.Println("err2",err)
				return
			}
			err = config.LoadFile(cacheFile)
			if err != nil {
				log.Println("err3",err)
				return
			}
			err = config.Scan(model)
			if err != nil {
				log.Println("err4",err)
				return
			}
			fmt.Println("model",model)

		},
	})
	if err != nil {
		log.Println("listen config error,dataid:",dataid,err)
	}
}
