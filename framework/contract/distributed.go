package contract

import "time"


//定义字符串凭证
const DistributedKey = "hade:distributed"

//分布式服务
type Distributed interface {
	//分布式选择器，所有节点对某个服务进行抢占，只选择其中一个节点
	Select(serviceName string,appID string,holdTime time.Duration) (selectAppID string,err error)
}