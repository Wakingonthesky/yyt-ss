package config

import "strconv"

type Mysql struct {
	Host     string
	Port     int
	Config   string //高级配置
	DB       string
	User     string
	Password string
	LogLevel string //日志等级
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "? " + m.Config
}
