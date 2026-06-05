---
title: 內建規則 - Validator
---

# 內建驗證標籤

Validator 內建豐富的驗證標籤，覆蓋常見校驗場景。

## 基礎規則

| 標籤 | 說明 | 範例 |
|------|------|------|
| `required` | 非零值 | `validate:"required"` |
| `email` | 郵箱格式 | `validate:"email"` |
| `url` | URL 格式 | `validate:"url"` |
| `alpha` | 純字母（大小寫） | `validate:"alpha"` |
| `alphanum` | 字母 + 數字 | `validate:"alphanum"` |
| `json` | JSON 格式 | `validate:"json"` |
| `uuid` | UUID 格式 | `validate:"uuid"` |

## 字母/數字變體

| 標籤 | 說明 | 範例 |
|------|------|------|
| `uppercase` | 僅大寫字母 | `validate:"uppercase"` |
| `lowercase` | 僅小寫字母 | `validate:"lowercase"` |
| `alphanum_upper` | 大寫字母 + 數字 | `validate:"alphanum_upper"` |
| `alphanum_lower` | 小寫字母 + 數字 | `validate:"alphanum_lower"` |

## 網絡與地址

| 標籤 | 說明 | 範例 |
|------|------|------|
| `ip` | IP 地址（v4 或 v6） | `validate:"ip"` |
| `ipv4` | IPv4 地址 | `validate:"ipv4"` |
| `ipv6` | IPv6 地址 | `validate:"ipv6"` |
| `ip_addr` | IP 地址 | `validate:"ip_addr"` |
| `ip4_addr` | IPv4 地址 | `validate:"ip4_addr"` |
| `ip6_addr` | IPv6 地址 | `validate:"ip6_addr"` |
| `cidr` | CIDR 記法 | `validate:"cidr"` |
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
| `port` | 端口號 | `validate:"port"` |
| `hostname` | 主機名（RFC 952） | `validate:"hostname"` |
| `hostname_rfc1123` | 主機名（RFC 1123） | `validate:"hostname_rfc1123"` |
| `hostname_port` | 主機名:端口 | `validate:"hostname_port"` |
| `fqdn` | 完全限定域名 | `validate:"fqdn"` |
| `uri` | URI | `validate:"uri"` |
| `http_url` | HTTP(S) URL | `validate:"http_url"` |
| `https_url` | 僅 HTTPS URL | `validate:"https_url"` |
| `origin` | Web Origin | `validate:"origin"` |
| `url_encoded` | URL 編碼 | `validate:"url_encoded"` |
| `datauri` | Data URI | `validate:"datauri"` |
| `urn_rfc2141` | URN（RFC 2141） | `validate:"urn_rfc2141"` |

## 數值/長度

| 標籤 | 說明 | 範例 |
|------|------|------|
| `min=N` | 最小值/長度 | `validate:"min=3"` |
| `max=N` | 最大值/長度 | `validate:"max=100"` |
| `len=N` | 精確長度 | `validate:"len=11"` |
| `minlen=N` | 最小長度 | `validate:"minlen=6"` |
| `maxlen=N` | 最大長度 | `validate:"maxlen=20"` |
| `range=min,max` | 數值範圍 | `validate:"range=0.0,5.0"` |

## 集合/模式

| 標籤 | 說明 | 範例 |
|------|------|------|
| `in=v1,v2,...` | 列舉值 | `validate:"in=male,female"` |
| `notin=v1,v2,...` | 排除列舉 | `validate:"notin=admin,root"` |
| `pattern=regex` | 正則匹配 | `validate:"pattern=^[A-Z]"` |
| `containspecial` | 含特殊字元 | `validate:"containspecial"` |
| `strong_password` | 強密碼 | `validate:"strong_password"` |

## 邏輯組合

```go
// And — 全部滿足
combined := validator.And(
    validator.Required(),
    validator.MinLength(5),
)

// Or — 滿足其一
either := validator.Or(
    validator.Email(),
    validator.Pattern(`^\d+$`),
)

// Not — 取反
negated := validator.Not(validator.In("a", "b"))
```
