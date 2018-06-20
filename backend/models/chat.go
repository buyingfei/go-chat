package models

import (
	"net"
	"github.com/astaxie/beego"
)

func GetToken() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		beego.Error("get IP error :", err)
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return "ip" + ipnet.IP.String()
			}
		}
	}
	return  ""
}
