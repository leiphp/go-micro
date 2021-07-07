package Boot

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/logger"
	"time"
)

var BootErrChan chan error

//启动函数，用于各种初始化，首先完成配置中心的初始化
func BootInit(){
	BootErrChan = make(chan error)
	InitConfig()
	go func() {
		err := WaitForConfigReady(time.Second*8)
		fmt.Println("WaitForConfigReady err：",err)
		if err != nil {
			BootErrChan<-err
		}else{
			WaitForDbReady(time.Second*5) //数据库初始化
		}
	}()
}

//服务器已经准备好了，后面要不断扩展
func ServerIsReady() bool {
	if configIsReady() && mysql_db != nil {
		return true
	}
	return false
}

//通用的超时控制函数
func WaitForReady(d time.Duration, f func() error,success string,fail string) error {
	ctx, cancel := context.WithTimeout(context.Background(),d)
	defer cancel()
	for{
		select {
			case<-ctx.Done():
				return fmt.Errorf(fail)
		default:
			err := f()
			//logger.Info("f err:",err)
			if err == nil {//没有错误，则直接返回
				logger.Info(success)
				return nil
			}else if IsFatalError(err){//如果是致命异常，则无需再次循环
				return fmt.Errorf("出现致命异常：%s",err.Error())
			}
		}
	}
}