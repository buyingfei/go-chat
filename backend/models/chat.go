package models

import (
	"net"
	"github.com/astaxie/beego"
	"math/rand"
	"time"
	"strconv"
	"strings"
	"net/http"
)

func GetIP(req * http.Request) string {
	remoteAddr := req.RemoteAddr
	beego.Info("XRealIP",req.Header.Get("XRealIP"))
	if ip := req.Header.Get("XRealIP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("XForwardedFor"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

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

//生成随机字符串
func GetRandomString(length  int) string{
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}
