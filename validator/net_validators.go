package validator

import (
	"net"
	"net/url"
	"regexp"
	"strings"
)

// NetValidators 返回所有网络验证器注册表
func NetValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"ip":               validateIP,
		"ipv6":             validateIPv6,
		"ip_addr":          validateIP,
		"ip4_addr":         validateIPv4,
		"ip6_addr":         validateIPv6,
		"cidr":             validateCIDR,
		"cidrv4":           validateCIDRv4,
		"cidrv6":           validateCIDRv6,
		"hostname":         validateHostname,
		"hostname_rfc1123": validateHostnameRFC1123,
		"hostname_port":    validateHostnamePort,
		"fqdn":             validateFQDN,
		"tcp_addr":         validateTCPAddr,
		"tcp4_addr":        validateTCP4Addr,
		"tcp6_addr":        validateTCP6Addr,
		"udp_addr":         validateUDPAddr,
		"udp4_addr":        validateUDP4Addr,
		"udp6_addr":        validateUDP6Addr,
		"unix_addr":        validateUnixAddr,
		"uri":              validateURI,
		"http_url":         validateHTTPURL,
		"url_encoded":      validateURLEncoded,
		"datauri":          validateDataURI,
		"urn_rfc2141":      validateURNRFC2141,
	}
}

func validateIP(fl FieldLevel) bool {
	return net.ParseIP(fl.Field().String()) != nil
}


func validateIPv6(fl FieldLevel) bool {
	return net.ParseIP(fl.Field().String()) != nil && strings.Contains(fl.Field().String(), ":")
}

func validateCIDR(fl FieldLevel) bool {
	_, _, err := net.ParseCIDR(fl.Field().String())
	return err == nil
}

func validateCIDRv4(fl FieldLevel) bool {
	ip, _, err := net.ParseCIDR(fl.Field().String())
	if err != nil {
		return false
	}
	return ip.To4() != nil
}

func validateCIDRv6(fl FieldLevel) bool {
	ip, _, err := net.ParseCIDR(fl.Field().String())
	if err != nil {
		return false
	}
	return ip.To4() == nil && ip.To16() != nil
}

// hostnameRegex RFC 952: 字母/数字/连字符，不以连字符开头/结尾，每段 ≤63
var hostnameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-]{0,62}$`)

func validateHostname(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) > 253 || s == "" {
		return false
	}
	parts := strings.Split(s, ".")
	for _, p := range parts {
		if !hostnameRegex.MatchString(p) {
			return false
		}
	}
	return true
}

// hostnameRFC1123Regex 允许以数字开头
var hostnameRFC1123Regex = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\-]{0,62}$`)

func validateHostnameRFC1123(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) > 253 || s == "" {
		return false
	}
	parts := strings.Split(s, ".")
	for _, p := range parts {
		if !hostnameRFC1123Regex.MatchString(p) {
			return false
		}
	}
	return true
}

func validateHostnamePort(fl FieldLevel) bool {
	s := fl.Field().String()
	host, port, err := net.SplitHostPort(s)
	if err != nil {
		return false
	}
	if port == "" {
		return false
	}
	// 验证端口
	for _, c := range port {
		if c < '0' || c > '9' {
			return false
		}
	}
	if host == "" {
		return false
	}
	// 验证 host 是合法 IP
	if net.ParseIP(host) != nil {
		return true
	}
	// 当作 hostname 检查（分段验证）
	parts := strings.Split(host, ".")
	for _, p := range parts {
		if p == "" || !hostnameRFC1123Regex.MatchString(p) {
			return false
		}
	}
	return len(parts) > 0
}



func validateFQDN(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) > 253 || s == "" || !strings.Contains(s, ".") {
		return false
	}
	parts := strings.Split(s, ".")
	// 最后一段必须是 TLD（纯字母）
	tld := parts[len(parts)-1]
	if len(tld) < 2 {
		return false
	}
	for _, c := range tld {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}
	for _, p := range parts {
		if !hostnameRFC1123Regex.MatchString(p) {
			return false
		}
	}
	return true
}

func validateTCPAddr(fl FieldLevel) bool {
	_, err := net.ResolveTCPAddr("tcp", fl.Field().String())
	return err == nil
}

func validateTCP4Addr(fl FieldLevel) bool {
	_, err := net.ResolveTCPAddr("tcp4", fl.Field().String())
	return err == nil
}

func validateTCP6Addr(fl FieldLevel) bool {
	_, err := net.ResolveTCPAddr("tcp6", fl.Field().String())
	return err == nil
}

func validateUDPAddr(fl FieldLevel) bool {
	_, err := net.ResolveUDPAddr("udp", fl.Field().String())
	return err == nil
}

func validateUDP4Addr(fl FieldLevel) bool {
	_, err := net.ResolveUDPAddr("udp4", fl.Field().String())
	return err == nil
}

func validateUDP6Addr(fl FieldLevel) bool {
	_, err := net.ResolveUDPAddr("udp6", fl.Field().String())
	return err == nil
}

func validateUnixAddr(fl FieldLevel) bool {
	_, err := net.ResolveUnixAddr("unix", fl.Field().String())
	return err == nil
}

func validateURI(fl FieldLevel) bool {
	u, err := url.ParseRequestURI(fl.Field().String())
	return err == nil && u.Scheme != "" && u.Host != ""
}

func validateHTTPURL(fl FieldLevel) bool {
	u, err := url.Parse(fl.Field().String())
	if err != nil {
		return false
	}
	scheme := strings.ToLower(u.Scheme)
	return (scheme == "http" || scheme == "https") && u.Host != ""
}

func validateURLEncoded(fl FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return false
	}
	// URL 编码字符串不应包含未编码的特殊字符
	for i := 0; i < len(s); i++ {
		c := s[i]
		// 允许: 字母、数字、-、_、.、%
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') ||
			c == '-' || c == '_' || c == '.' || c == '%') {
			return false
		}
	}
	return true
}

func validateDataURI(fl FieldLevel) bool {
	s := fl.Field().String()
	if !strings.HasPrefix(s, "data:") {
		return false
	}
	// data:[<mediatype>][;base64],<data>
	rest := s[5:]
	// 至少要有逗号分隔
	commaIdx := strings.Index(rest, ",")
	if commaIdx < 0 {
		return false
	}
	meta := rest[:commaIdx]
	// meta 部分应包含 mediatype 或 ;base64
	return meta != "" || commaIdx == 0
}

// urnRegex URN RFC 2141
var urnRegex = regexp.MustCompile(`^urn:[a-zA-Z0-9][a-zA-Z0-9\-]{0,31}:[^\s]+$`)

func validateURNRFC2141(fl FieldLevel) bool {
	return urnRegex.MatchString(fl.Field().String())
}
