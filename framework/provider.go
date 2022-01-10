package framework

//定义了如何创建一个新的实例，所有服务容器的创建服务
//它规定了所有创建服务实例的方法必须有相同的参数 interface{}数组，并且返回interface{}和错误信息这两个数据结构
type NewInstance func(...interface{}) (interface{}, error)

//定义一个服务提供者需要实现的接口
type ServiceProvider interface {
	//在服务容器中注册了一个实例化服务的方法，是否在注册的时候就实例化这个服务,需要参考IsDefer接口
	Register(Container) NewInstance

	//在调用实例化服务的时候会调用，可以把一些准备工作放在这里，如果Boot返回error，整个服务实例化就会实例化失败
	Boot(Container) error

	//决定是否在注册的时候实例化这个服务，如果不是注册的时候实例化，那就是在第一次make的时候进行实例化操作
	IsDefer() bool

	//params定义传递给NewInstance的参数，可以自定义多个
	Params(Container) []interface{}

	//代表了这个服务提供者的凭证
	Name() string
}
