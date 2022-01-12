package contract

const (
	//代表生产环境
	EnvProduction="production"

	//代表测试环境
	EnvTesting= "testing"

	//代表开发环境
	EnvDevelopment="development"


	//环境变量服务字符串凭证
	EnvKey = "hade:env"
)

//定义环境变量的获取服务
type Env interface {
	//获取当前的环境
	AppEnv() string

	//判断一个环境变量是否有被设置
	IsExist(string) bool

	//获取某个环境变量,如果没有返回""
	Get(string) string

	//获取所有的环境变量
	All() map[string]string

	//为什么这里没有Set方法，因为考虑到环境变量本质代表的是当前程序运行的环境，是一个程序运行时就已经固定的
	//不应该允许程序运行中进行设置。
}