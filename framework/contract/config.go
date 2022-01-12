package contract

import (
	"time"
)

const(
	ConfigKey = "hade:config"
)

//定义了配置文件服务，读取配置文件，支持点分割ide路径读取
//例如：.Get("app.name")表示从app文件中读取name属性
//建议使用yaml属性

type Config interface {
	//检查一个属性是否存在
	IsExist(key string) bool

	Get(key string) interface{}

	GetBool(key string) bool

	GetInt(key string) int

	GetFloat64(key string) float64

	GetTime(key string) time.Time

	GetString(key string) string

	GetIntSlice(key string) []int

	GetStringSlice(key string) []string

	GetStringMap(key string) map[string]interface{}

	GetStringMapString(key string) map[string]string
	
	GetStringMapStringSlice(key string) map[string][]string

	//加载配置到某个对象
	Load(key string,val interface{}) error
}