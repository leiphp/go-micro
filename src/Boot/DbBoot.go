package Boot

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/logger"
	"sync"
	"time"
)


//mysql相关
var mysql_db *gorm.DB

func WaitForDbReady(d time.Duration) {
	go func() {
		err := WaitForReady(d,func() error {
			return InitMysql()
		},"数据库初始化成功","数据库初始化失败")
		if err != nil {
			BootErrChan<-err
		}
	}()
}

//func init(){
//	InitMysql()
//}
func ReloadDB() error {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	if mysql_db != nil {
		err := mysql_db.Close()
		if err != nil {
			return err
		}
		mysql_db = nil
		return InitMysql()
	}
	return nil
}


func InitMysql() error {
	var err error
	mysql_db, err = gorm.Open("mysql", JConfig.Data.Mysql.Dsn)
	logger.Info("zzzz",JConfig.Data.Mysql.Dsn)
	if err != nil {
		mysql_db=nil
		 return NewFatalError(err.Error()) //这里返回致命异常
	}
	mysql_db.SingularTable(true)
	mysql_db.DB().SetMaxIdleConns(  JConfig.Data.Mysql.Maxidle)
	mysql_db.DB().SetMaxOpenConns( JConfig.Data.Mysql.Maxopen)
	mysql_db.LogMode(true)


	return nil
}

func GetDB() *gorm.DB {
	return mysql_db
}