package network

import (
	"net"
	"net/http"
	"strings"
)

func RealIpFromHeader(header http.Header) string {
	val := header.Get("Cf-Connecting-Ip")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = header.Get("Cf-Pseudo-Ipv4")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = header.Get("Cf-Connecting-Ipv6")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = header.Get("Cf-Pseudo-Ipv6")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = header.Get("Fastly-Client-Ip")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = header.Get("True-Client-Ip")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = header.Get("X-Real-IP")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = header.Get("X-Client-IP")
	if val != "" {
		if !IsLocalIp(val) {
			return val
		}
	}

	val = header.Get("X-Original-Forwarded-For")
	if val != "" {
		for _, v := range strings.Split(val, ",") {
			if !IsLocalIp(v) {
				return v
			}
		}
		if !IsLocalIp(val) {
			return val
		}
	}

	val = header.Get("X-Forwarded-For")
	if val != "" {
		for _, v := range strings.Split(val, ",") {
			if net.ParseIP(v) != nil {
				return v
			}
		}
		if net.ParseIP(val) != nil {
			return val
		}
	}

	val = header.Get("X-Forwarded")
	if val != "" {
		for _, v := range strings.Split(val, ",") {
			if net.ParseIP(v) != nil {
				return v
			}
		}
		if net.ParseIP(val) != nil {
			return val
		}
	}

	val = header.Get("Forwarded-For")
	if val != "" {
		for _, v := range strings.Split(val, ",") {
			if net.ParseIP(v) != nil {
				return v
			}
		}
		if net.ParseIP(val) != nil {
			return val
		}
	}

	val = header.Get("Forwarded")
	if val != "" {
		for _, v := range strings.Split(val, ",") {
			if net.ParseIP(v) != nil {
				return v
			}
		}
		if net.ParseIP(val) != nil {
			return val
		}
	}

	return ""
}
