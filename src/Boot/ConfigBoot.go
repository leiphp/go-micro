package Boot

import (
	"fmt"
	"time"
)

func configIsReady() bool {
	//以下设定，哪些配置值必须有值，才算成功。有一个失败都算not ready todo弃用
	//checks := []interface{}{JConfig.Data.Mysql}
	//for _,c := range checks {
	//	cv := reflect.ValueOf(c) //主要这里，interface=nil是不成立的
		//interface{} 类型的变量包含了2个指针，一个指针指向在编译时确定的类型，另外一个指针指向实际值
	//	if cv.IsNil(){
	//		return false
	//	}
	//}
	//return true

	checks := []ConfigInterface{JConfig.Data.Mysql}
	for _,c := range checks {
		if c.IsLoad(){
			return false
		}
	}
	return true
}

//配置中心相关，初始化，支持超时参数
func WaitForConfigReady(d time.Duration) error {
	return WaitForReady(d, func() error {
		if configIsReady() {
			return nil
		}else {
			return fmt.Errorf("wait")
		}
	},"配置加载成功","配置加载失败")
}