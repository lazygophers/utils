package network

import "net/netip"

func IsLocalIp(ip string) bool {
	i, err := netip.ParseAddr(ip)
	if err != nil {
		return false
	}

	return i.IsPrivate() || i.IsLoopback() || i.IsLinkLocalUnicast()
}
