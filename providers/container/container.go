package container

import (
	"log"
	"github.com/go-redis/redis"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type (
	Container interface {
		Redis() (*redis.Client, error)
		Database() (*gorm.DB, error)
		Logger() *log.Logger
		Response(ctx iris.Context) *Response
		Env() *Environment
		Account() *Account
		SetAccount(account *Account)
		WatchEnv() *Containers
	}
	Containers struct {
		rdx               *redis.Client // redis
		db                *gorm.DB      // gorm
		env               *Environment  // 配置
		logger            *log.Logger   // 日志
		account           *Account      // 账户
	}
)

func XContainer(env Environment) *Containers {
	var container *Containers
	container = &Containers{
		env: &env,
	}
	// 初始化日志
	container.Logger()
	// 初始化 redis
	_, err := container.Redis()
	if err != nil {
		panic(err)
	}
	// 初始化 数据库
	_, err = container.Database()
	if err != nil {
		panic(err)
	}
	return container
}
