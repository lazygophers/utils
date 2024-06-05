package network

import (
	"github.com/lazygophers/log"
	"net"
)

func GetInterfaceIpByName(name string, prev6 bool) string {
	inter, err := net.InterfaceByName("eth0")
	if err != nil {
		log.Debugf("err:%v", err)
		return ""
	}

	address, err := inter.Addrs()
	if err != nil {
		log.Errorf("err:%v", err)
		return ""
	}

	return GetInterfaceIpByAddrs(address, prev6)
}

func GetInterfaceIpByAddrs(address []net.Addr, prev6 bool) string {
	var v4 string
	for _, addr := range address {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To16() != nil {
				// v6 的地址
				if prev6 {
					return ipnet.IP.To16().String()
				}
			} else if ipnet.IP.To4() != nil {
				// v4 地址
				if !prev6 {
					return ipnet.IP.To4().String()
				} else {
					v4 = ipnet.IP.To4().String()
				}
			}
		}
	}

	return v4
}

func GetListenIp(prev6 ...bool) string {
	// 找到内网 IP
	var _prev6 bool
	_prev6 = len(prev6) > 0 && prev6[0]

	// 先尝试一下常用 eth0 网卡
	if ip := GetInterfaceIpByName("eth0", _prev6); ip != "" {
		return ip
	}

	// 先尝试一下常用 en0 网卡
	if ip := GetInterfaceIpByName("en0", _prev6); ip != "" {
		return ip
	}

	address, err := net.InterfaceAddrs()
	if err != nil {
		log.Errorf("err:%v", err)
		return ""
	}

	if ip := GetInterfaceIpByAddrs(address, _prev6); ip != "" {
		return ip
	}

	log.Error("get interface ip failed")

	return ""
}
