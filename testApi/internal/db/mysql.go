package db

import (
	"context"
	"testApi/internal/config"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func NewMysql(mysqlConfig config.MysqlConfig) sqlx.SqlConn {
	mysql := sqlx.NewMysql(mysqlConfig.DataSource) //得到这样一个连接
	db, err := mysql.RawDB()                       //得到db数据库，便于设置它的参数
	if err != nil {
		panic(err)
	}
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(mysqlConfig.ConnectTimeout))
	defer cancel()
	err = db.PingContext(cxt) //ping一下数据库，看看能否连接成功
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return mysql
}
