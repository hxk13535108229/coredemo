package contract

const AppKey = "hade:app"

type App interface {
	//定义当前版本
	Version() string

	//定义项目基础地址
	BaseFolder() string

	//定义了配置文件的路径
	ConfigFolder() string

	//定义了日志所在路径
	LogFolder() string

	//定义业务自己的服务提供者地址
	ProviderFolder() string

	//定义业务自己定义的中间件
	MiddlewareFolder() string

	//定义业务定义的命令
	CommandFolder() string

	//定义业务的运行中间态信息
	RuntimeFolder() string

	//存放测试所需要的信息
	TestFolder() string
}