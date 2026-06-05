---
title: 内置规则 - Validator
---

# 内置验证标签

Validator 内置丰富的验证标签，覆盖常见校验场景。

## 基础规则

| 标签 | 说明 | 示例 |
|------|------|------|
| `required` | 非零值 | `validate:"required"` |
| `email` | 邮箱格式 | `validate:"email"` |
| `url` | URL 格式 | `validate:"url"` |
| `alpha` | 纯字母（大小写） | `validate:"alpha"` |
| `alphanum` | 字母 + 数字 | `validate:"alphanum"` |
| `json` | JSON 格式 | `validate:"json"` |
| `uuid` | UUID 格式 | `validate:"uuid"` |

## 字母/数字变体

| 标签 | 说明 | 示例 |
|------|------|------|
| `uppercase` | 仅大写字母 | `validate:"uppercase"` |
| `lowercase` | 仅小写字母 | `validate:"lowercase"` |
| `alphanum_upper` | 大写字母 + 数字 | `validate:"alphanum_upper"` |
| `alphanum_lower` | 小写字母 + 数字 | `validate:"alphanum_lower"` |

## 网络与地址

| 标签 | 说明 | 示例 |
|------|------|------|
| `ip` | IP 地址（v4 或 v6） | `validate:"ip"` |
| `ipv4` | IPv4 地址 | `validate:"ipv4"` |
| `ipv6` | IPv6 地址 | `validate:"ipv6"` |
| `ip_addr` | IP 地址 | `validate:"ip_addr"` |
| `ip4_addr` | IPv4 地址 | `validate:"ip4_addr"` |
| `ip6_addr` | IPv6 地址 | `validate:"ip6_addr"` |
| `cidr` | CIDR 记法 | `validate:"cidr"` |
| `cidrv4` | IPv4 CIDR | `validate:"cidrv4"` |
| `cidrv6` | IPv6 CIDR | `validate:"cidrv6"` |
| `tcp_addr` | TCP 地址 | `validate:"tcp_addr"` |
| `tcp4_addr` | TCPv4 地址 | `validate:"tcp4_addr"` |
| `tcp6_addr` | TCPv6 地址 | `validate:"tcp6_addr"` |
| `udp_addr` | UDP 地址 | `validate:"udp_addr"` |
| `udp4_addr` | UDPv4 地址 | `validate:"udp4_addr"` |
| `udp6_addr` | UDPv6 地址 | `validate:"udp6_addr"` |
| `unix_addr` | Unix 域套接字地址 | `validate:"unix_addr"` |
| `uds_exists` | Unix 域套接字已存在 | `validate:"uds_exists"` |
| `mac` | MAC 地址 | `validate:"mac"` |
| `port` | 端口号 | `validate:"port"` |
| `hostname` | 主机名（RFC 952） | `validate:"hostname"` |
| `hostname_rfc1123` | 主机名（RFC 1123） | `validate:"hostname_rfc1123"` |
| `hostname_port` | 主机名:端口 | `validate:"hostname_port"` |
| `fqdn` | 完全限定域名 | `validate:"fqdn"` |
| `uri` | URI | `validate:"uri"` |
| `http_url` | HTTP(S) URL | `validate:"http_url"` |
| `https_url` | 仅 HTTPS URL | `validate:"https_url"` |
| `origin` | Web Origin | `validate:"origin"` |
| `url_encoded` | URL 编码 | `validate:"url_encoded"` |
| `datauri` | Data URI | `validate:"datauri"` |
| `urn_rfc2141` | URN（RFC 2141） | `validate:"urn_rfc2141"` |

## 数值/长度

| 标签 | 说明 | 示例 |
|------|------|------|
| `min=N` | 最小值/长度 | `validate:"min=3"` |
| `max=N` | 最大值/长度 | `validate:"max=100"` |
| `len=N` | 精确长度 | `validate:"len=11"` |
| `minlen=N` | 最小长度 | `validate:"minlen=6"` |
| `maxlen=N` | 最大长度 | `validate:"maxlen=20"` |
| `range=min,max` | 数值范围 | `validate:"range=0.0,5.0"` |

## 集合/模式

| 标签 | 说明 | 示例 |
|------|------|------|
| `in=v1,v2,...` | 枚举值 | `validate:"in=male,female"` |
| `notin=v1,v2,...` | 排除枚举 | `validate:"notin=admin,root"` |
| `pattern=regex` | 正则匹配 | `validate:"pattern=^[A-Z]"` |
| `containspecial` | 含特殊字符 | `validate:"containspecial"` |
| `strong_password` | 强密码 | `validate:"strong_password"` |

## 逻辑组合

```go
// And — 全部满足
combined := validator.And(
    validator.Required(),
    validator.MinLength(5),
)

// Or — 满足其一
either := validator.Or(
    validator.Email(),
    validator.Pattern(`^\d+$`),
)

// Not — 取反
negated := validator.Not(validator.In("a", "b"))
```
