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

	//appID
	AppID() string

	//加载新的AppConfig,可以为对应的函数转为小写下划线，例子：ConfigFolder -》 config_folder
	LoadAppConfig(kv map[string]string)

	//定义业务代码所在的目录，用于监控文件变更使用
	AppFolder() string

	//存放部署的时候创建的文件夹
	DeployFolder() string
}