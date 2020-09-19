package common

import "gin_dev/config"

//获取带协议头的域名
func DomainName()string{
	return config.GetString("domain_name")
}