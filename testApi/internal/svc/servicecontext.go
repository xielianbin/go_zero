package svc

import (
	"testApi/internal/config"
	"testApi/internal/db"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Mysql  sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := db.NewMysql(c.MysqlConfig) //对上面的结构体进行依赖注入
	return &ServiceContext{
		Config: c,
		Mysql:  mysql,
	}
}
