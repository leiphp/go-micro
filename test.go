package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go-micro/src/Boot"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type JtConfig struct {
	Db struct{
		Ip 	string `json:"ip"`
		Port int	`json:"port"`
	} `json:"db"`
	Redis struct{
		Ip 	string `json:"ip"`
		Port int	`json:"port"`
	} `json:"redis"`
}

func WaitForReady() error {
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*5)
	defer cancel()
	for{
		select {
			case <-ctx.Done():
				return fmt.Errorf("init config error")
		default:
			if Boot.JConfig.Data.Mysql != nil {
				return nil
			}
		}
	}
}

func main() {
	Boot.BootInit()
	r := gin.New()
	//r.Use(CheckForReade())
	r.Handle("GET","/", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":Boot.JConfig.Data.Mysql})
	})

	go func() {
		err := r.Run(":8112")
		if err != nil {
			Boot.BootErrChan<-err
		}
	}()

	go func() {
		sig_c := make(chan os.Signal)
		signal.Notify(sig_c,syscall.SIGINT,syscall.SIGTERM)
		Boot.BootErrChan<-fmt.Errorf("%s",<-sig_c)
	}()

	getErr := <-Boot.BootErrChan
	logger.Info(getErr)
}

//网址服务读取配置，延迟处理
func main5(){
	Boot.InitConfig()
	fmt.Println(Boot.JConfig.Data.Mysql)

	err := WaitForReady()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Handle("GET","/", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":Boot.JConfig.Data.Mysql})
	})
	r.Run(":8111")
}

//封装初始化去nacos拿配置
func main4(){
	Boot.InitConfig()
	fmt.Println(Boot.JConfig.Data.Mysql)

	for{
		time.Sleep(time.Second*1)
		fmt.Println(Boot.JConfig.Data.Mysql)
	}
}

//从nacos配置中心拿配置文件
func main3(){
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:	"localhost",
			ContextPath: "/nacos",
			Port: 8848,
		},
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId:   "txy-sysconfig",
		Group:    "TXY_GROUP",
		Content:  "",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println(data)
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	for{
		time.Sleep(time.Second*1)
	}

}

//从根目录读取配置文件
func main2(){
	//configFile := "app.json"
	configFile := "app.yaml"
	err := config.LoadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	jtConfig := &JtConfig{}
	//err = config.Get("100txy").Scan(jtConfig)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(jtConfig)
	w,_ := config.Watch("100txy")
	v,_ := w.Next()//卡主
	v.Scan(jtConfig)
	fmt.Println(jtConfig)

}