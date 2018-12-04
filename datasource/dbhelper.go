package datasource

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/liukunxin/superstar-db/conf"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock 		  sync.Mutex
)

//主库，单例
func InstanceMaster() *xorm.Engine  {
	if masterEngine != nil{
		return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	c := conf.MasterDbConfig
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err :=xorm.NewEngine(conf.DriverName, driverSource)
	if err!=nil{
		log.Fatalf("dbhelper.DbInstanceMaster,", err)
		return nil
	}

	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(false)
	engine.SetTZDatabase(conf.SysTimeLocation)

	//性能优化的时候才考虑，加上本机的SQL缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(),1000)
	engine.SetDefaultCacher(cacher)
	masterEngine = engine
	return engine
}

//从库，单例
func InstanceSlave() *xorm.Engine  {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		 return slaveEngine
	}
	c := conf.SlaveDbConfig
	engine, err := xorm.NewEngine(conf.DriverName,
					fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
						c.User, c.Pwd, c.Host, c.Port, c.DbName))
	if err	!= nil {
		log.Fatal("dbhelper", "DbInstanceMaster", err)
	}
	engine.SetTZLocation(conf.SysTimeLocation)
	slaveEngine = engine
	return engine
}
