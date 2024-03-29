package gateway

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
)

func InitConfig() Routes{
	configFile := "gateway.yaml"
	err := config.LoadFile(configFile)
	if err!= nil {
		log.Fatal(err)
	}
	r := make(Routes,0)
	err = config.Get("routes").Scan(&r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}