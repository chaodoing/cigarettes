package container

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Environment 配置管理
type Environment struct {
	XMLName xml.Name     `xml:"root" json:"-"`
	handle  *viper.Viper // 配置驱动

	App struct {
		Name        string `json:"name" xml:"name"`
		Host        string `json:"host" xml:"host"`
		Port        string `json:"port" xml:"port"`
		LogLevel    string `json:"log_level" xml:"log_level"`
		Crossdomain bool   `json:"crossdomain" xml:"crossdomain"`
	} `json:"app" xml:"app"`
	Template struct {
		Directory string `json:"directory" xml:"directory"`
		Extension string `json:"extension" xml:"extension"`
		Favicon   string `json:"favicon" xml:"favicon"`
	} `json:"template" xml:"template"`
	WebStatic struct {
		AccessPath string `json:"access_path" xml:"access_path"`
		LocalPath  string `json:"local_path" xml:"local_path"`
	} `json:"web_static" xml:"web_static"`
	Upload struct {
		Maximum    int    `json:"maximum" xml:"maximum"`
		AccessPath string `json:"access_path" xml:"access_path"`
		LocalPath  string `json:"local_path" xml:"local_path"`
	} `json:"upload" xml:"upload"`
	Database struct {
		LogLevel string `json:"log_level" xml:"log_level"`
		Type     string `json:"type" xml:"type"`
		User     string `json:"user" xml:"user"`
		Password string `json:"password" xml:"password"`
		Host     string `json:"host" xml:"host"`
		Port     int    `json:"port" xml:"port"`
		Name     string `json:"name" xml:"name"`
		Charset  string `json:"charset" xml:"charset"`
	} `json:"database" xml:"database"`
	Log struct {
		Stdout   bool   `json:"stdout" xml:"stdout"`
		Path     string `json:"path" xml:"path"`
		Name     string `json:"name" xml:"name"`
		Sql      string `json:"sql" xml:"sql"`
		Size     int    `json:"size" xml:"size"`
		Day      int    `json:"day" xml:"day"`
		Backup   int    `json:"backup" xml:"backup"`
		Local    bool   `json:"local" xml:"local"`
		Compress bool   `json:"compress" xml:"compress"`
	} `json:"log" xml:"log"`
	Redis struct {
		Host     string `json:"host" xml:"host"`
		Port     string `json:"port" xml:"port"`
		Password string `json:"password" xml:"password"`
		Database int    `json:"database" xml:"database"`
		Expire   int    `json:"expire" xml:"expire"`
	} `json:"redis" xml:"redis"`
	Account struct {
		Prefix   string `json:"prefix" xml:"prefix"`
		LifeTime int    `json:"lifetime" xml:"lifetime"`
	} `json:"account" xml:"account"`
}

// Json 输出配置文件json内容
//	@return string
//	@return error
func (e Environment) Json() (string, error) {
	data, err := json.MarshalIndent(e, "", "\t")
	return string(data), err
}

// Xml 输出配置文件json内容
//	@return string error
func (e Environment) Xml() (string, error) {
	data, err := xml.MarshalIndent(e, "", "\t")
	return string(data), err
}

// Dialect 转换数据配置
//	@return dialect=mysql
//	@return schema=root:123.com@tcp(127.0.0.1:3306)/arrangement?charset=utf8mb4&parseTime=True&loc=Local
//	@return logMode=false
func (e Environment) Dialect() (dialect, schema string) {
	dialect = e.Database.Type
	schema = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local", e.Database.User, e.Database.Password, e.Database.Host, e.Database.Port, e.Database.Name, e.Database.Charset)
	return
}

// Env 加载配置文件
//	@param string appConfig 配置文件名称
//	@return container.Environment 配置数据
func Env(appConfig string) Environment {
	var (
		handle = viper.New()
		err    error
	)
	handle.SetConfigType("ini")
	handle.SetConfigFile(os.Getenv("PWD") + "/.env")
	err = handle.ReadInConfig()
	if err == nil {
		var config Environment
		err = handle.Unmarshal(&config)
		if err != nil {
			panic(err)
		}
		config.handle = handle
		return config
	} else {
		app := viper.New()
		app.SetConfigType("ini")
		app.SetConfigFile(appConfig)
		err = app.ReadInConfig()
		if err != nil {
			panic(err)
		}
		var config Environment
		err = app.Unmarshal(&config)
		if err != nil {
			panic(err)
		}
		config.handle = app
		return config
	}
}
