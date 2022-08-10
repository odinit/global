package gredis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	Client *redis.Client
	Nil    = redis.Nil
)

// Init 初始化连接
func Init(host string, port int, password string, db, poolSize, minIdleConns int) (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Password:     password, // no password set
		DB:           db,       // use default DB
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
	})

	_, err = Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = Client.Close()
}
