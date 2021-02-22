package container

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis"
	"github.com/kataras/iris/v12"
	"github.com/natefinch/lumberjack"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

)

// SetAccount 设置用户信息
//	@param *Account account 账户
func (c *Containers) SetAccount(account *Account) {
	c.account = account
}


// Redis
func (c *Containers) Redis() (*redis.Client, error) {
	if c.rdx != nil {
		return c.rdx, nil
	} else {
		var addr = fmt.Sprintf("%v:%v", c.env.Redis.Host, c.env.Redis.Port)
		rdb := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: c.env.Redis.Password,
			DB:       c.env.Redis.Database,
		})
		pong, err := rdb.Ping().Result()
		if err != nil || !strings.EqualFold(pong, "PONG") {
			return nil, err
		}
		c.rdx = rdb
		return c.rdx, nil
	}
}

// Database 数据库连接
func (c *Containers) Database() (*gorm.DB, error) {
	if c.db != nil {
		return c.db, nil
	}
	dialect, schema := c.env.Dialect()
	driver := mysql.New(mysql.Config{
		DriverName: dialect,
		DSN:        schema,
	})
	var (
		db  *gorm.DB
		err error
	)
	logx := log.New(os.Stdout, "", log.LstdFlags)
	if !c.env.Log.Stdout {
		logx.SetOutput(&lumberjack.Logger{
			Filename:   os.Getenv("PWD") + "/" + c.env.Log.Path + "/" + c.env.Log.Sql,
			MaxAge:     c.env.Log.Day,
			MaxSize:    c.env.Log.Size,
			MaxBackups: c.env.Log.Backup,
			LocalTime:  c.env.Log.Local,
			Compress:   c.env.Log.Compress,
		})
	}
	var loglevel = map[string]int{
		"silent": 1,
		"error":  2,
		"warn":   3,
		"info":   4,
	}
	logs := logger.New(logx, logger.Config{
		SlowThreshold: time.Second,
		Colorful:      c.env.Log.Stdout,
		LogLevel:      logger.LogLevel(loglevel[strings.ToLower(c.env.Database.LogLevel)]),
	})
	db, err = gorm.Open(driver, &gorm.Config{
		Logger:      logs,
		PrepareStmt: false,
	})
	if err != nil {
		return nil, err
	}
	c.db = db
	return db, nil
}

// Logger 获取日志
func (c *Containers) Logger() *log.Logger {
	if c.logger != nil {
		return c.logger
	}
	logx := log.New(os.Stdout, "", log.LstdFlags)
	if !c.env.Log.Stdout {
		w := &lumberjack.Logger{
			Filename:   os.Getenv("PWD") + "/" + c.env.Log.Path + "/" + c.env.Log.Name,
			MaxAge:     c.env.Log.Day,
			MaxSize:    c.env.Log.Size,
			MaxBackups: c.env.Log.Backup,
			LocalTime:  c.env.Log.Local,
			Compress:   c.env.Log.Compress,
		}
		logx.SetOutput(w)
		log.SetOutput(w)
	}
	c.logger = logx
	return c.logger
}

// Environment 获取配置变量
func (c *Containers) Env() *Environment {
	return c.env
}

// WatchEnv 监听配置文件变化
//	@return *Containers
func (c *Containers) WatchEnv() *Containers {
	c.env.handle.WatchConfig()
	c.env.handle.OnConfigChange(func(in fsnotify.Event) {
		err := c.env.handle.Unmarshal(c.env)
		if err != nil {
			c.logger.Println(fmt.Printf("配置文件重新加载:%v [失败] ", in.Name))
		} else {
			c.rdx = nil
			c.db = nil
			c.logger.Println(fmt.Printf("配置文件重新加载:%v ", in.Name))
		}
	})
	return c
}

// Account 账户存储
func (c *Containers) Account() *Account {
	if c.account != nil {
		return c.account
	} else {
		c.account = Accounts(c.rdx, c.env)
		return c.account
	}
}

// Response 数据响应器
func (c *Containers) Response(ctx iris.Context) *Response {
	return &Response{
		ctx: ctx,
	}
}
