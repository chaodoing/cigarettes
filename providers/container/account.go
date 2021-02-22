package container

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Key     string        `json:"key" xml:"key"`
	Auth    string        `json:"auth" xml:"auth"`
	Life    time.Duration `json:"life" xml:"life"`
	Expires time.Duration `json:"expires" xml:"expires"`
	Prefix  string        `json:"prefix" xml:"prefix"`
	Json    string        `json:"json" xml:"json"`
	rdx     *redis.Client
}

func uid() (key string) {
	key = fmt.Sprintf("%v", uuid.NewV4())
	return
}

// Accounts 实例化账户信息存储
//	@param *redis.Client rdx redis 客户端
//	@param *Environment env 配置管理
//	@return *Account    账户数据
func Accounts(rdx *redis.Client, env *Environment) *Account {
	return &Account{
		Life:    time.Duration(env.Account.LifeTime) * time.Second,
		Expires: time.Duration(env.Redis.Expire) * time.Second,
		Prefix:  env.Account.Prefix,
		rdx:     rdx,
	}
}

// Profiles 获取账户信息
//	@param interface data
//	@return err 错误信息
func (s *Account) Profiles(data interface{}) (err error) {
	err = json.Unmarshal([]byte(s.Json), &data)
	return
}

// Logout 退出登录
//	@return err 错误信息
func (s *Account) Logout() error {
	return s.rdx.Del(s.Key).Err()
}

func (s *Account) Authorization(Authorization string) (*Account, error) {
	const prefix = "Bearer "
	if !strings.EqualFold(Authorization, "") {
		if strings.HasPrefix(Authorization, prefix) {
			s.Auth = strings.TrimLeft(Authorization, prefix)
		} else {
			s.Auth = Authorization
		}
	}
	var (
		uuidKey []byte
		err     error
	)
	uuidKey, err = base64.StdEncoding.DecodeString(s.Auth)
	if err != nil {
		return nil, err
	}
	s.Key = string(uuidKey)
	value, err := s.rdx.Get(s.Key).Result()
	if err != nil {
		if strings.EqualFold(value, "") {
			return nil, errors.New("用户登录信息错误")
		}
		return nil, err
	}
	s.Json = value
	err = s.rdx.Expire(s.Key, s.Life).Err()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// BasicAuthCache HttpBasic 认证
//	@param string string auth
//	@param string interface{} data
func (s *Account) BasicAuthCache(auth string, data interface{}) error {
	const prefix = "Basic "
	s.Key = auth
	s.Auth = strings.TrimPrefix(auth, prefix)
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	s.Json = string(jsonByte)
	err = s.rdx.Set(s.Auth, s.Json, s.Life).Err()
	if err != nil {
		return err
	}
	return nil
}

// Cache 缓存账户信息
func (s *Account) Cache(account interface{}) (string, error) {
	authorization := fmt.Sprintf("%v:%v", s.Prefix, uid())
	data, err := json.Marshal(account)
	if err != nil {
		return "", err
	}
	err = s.rdx.Set(authorization, string(data), s.Life).Err()
	if err != nil {
		return "", err
	}
	authorization = base64.StdEncoding.EncodeToString([]byte(authorization))
	return fmt.Sprintf("Bearer %v", authorization), nil
}

// Store 存储Json数据
func (s *Account) Set(key string, value interface{}, expires ...time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var exp time.Duration
	if len(expires) > 0 {
		exp = expires[0] * time.Second
	} else {
		exp = s.Expires
	}
	_, err = s.rdx.Set(key, string(data), exp).Result()
	if err != nil {
		return err
	} else {
		return nil
	}
}

// Get 获取存储的信息
//	@param string key
//	@param 引用    data
func (s *Account) Get(key string, data interface{}) error {
	var (
		value string
		err   error
	)
	value, err = s.rdx.Get(key).Result()
	if err != nil {
		return err
	}
	err = s.rdx.Expire(key, s.Expires).Err()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(value), &data)
}
