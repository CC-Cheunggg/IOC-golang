package dal

import "github.com/cc-cheunggg/ioc-golang/extension/state/redis"

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type RedisConn struct {
	Redis redis.RedisIOCInterface `normal:""`
}

func (conn *RedisConn) GetConn() redis.RedisIOCInterface {
	return conn.Redis
}
