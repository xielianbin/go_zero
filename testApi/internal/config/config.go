package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MysqlConfig MysqlConfig
}
type MysqlConfig struct {
	DataSource     string
	ConnectTimeout int64
}
