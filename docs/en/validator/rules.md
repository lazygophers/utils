---
title: Built-in Rules - Validator
---

# Built-in Validation Tags

Rich set of built-in validation tags covering common scenarios.

## Basic Rules

| Tag | Description | Example |
|-----|-------------|---------|
| `required` | Non-zero value | `validate:"required"` |
| `email` | Email format | `validate:"email"` |
| `url` | URL format | `validate:"url"` |
| `json` | JSON format | `validate:"json"` |
| `uuid` | UUID format | `validate:"uuid"` |

## Strings

| Tag | Description | Example |
|-----|-------------|---------|
| `alpha` | Letters only | `validate:"alpha"` |
| `alphaspace` | Letters + spaces | `validate:"alphaspace"` |
| `alphanum` | Letters + digits | `validate:"alphanum"` |
| `alphanumspace` | Letters + digits + spaces | `validate:"alphanumspace"` |
| `alphanumunicode` | Letters + digits (unicode) | `validate:"alphanumunicode"` |
| `alphaunicode` | Letters only (unicode) | `validate:"alphaunicode"` |
| `ascii` | ASCII characters only | `validate:"ascii"` |
| `printascii` | Printable ASCII only | `validate:"printascii"` |
| `multibyte` | Multi-byte characters | `validate:"multibyte"` |
| `boolean` | Boolean value | `validate:"boolean"` |
| `number` | Number | `validate:"number"` |
| `numeric` | Numeric | `validate:"numeric"` |
| `uppercase` | Uppercase letters only | `validate:"uppercase"` |
| `lowercase` | Lowercase letters only | `validate:"lowercase"` |
| `contains` | Contains substring | `validate:"contains=abc"` |
| `containsany` | Contains any character | `validate:"containsany=abc"` |
| `containsrune` | Contains rune | `validate:"containsrune=©"` |
| `excludes` | Excludes substring | `validate:"excludes=abc"` |
| `excludesall` | Excludes all characters | `validate:"excludesall=abc"` |
| `excludesrune` | Excludes rune | `validate:"excludesrune=©"` |
| `startswith` | Starts with | `validate:"startswith=Hello"` |
| `startsnotwith` | Does not start with | `validate:"startsnotwith=Bad"` |
| `endswith` | Ends with | `validate:"endswith=World"` |
| `endsnotwith` | Does not end with | `validate:"endsnotwith=Bad"` |

## Letter / Digit Variants

| Tag | Description | Example |
|-----|-------------|---------|
| `alphanum_upper` | Uppercase letters + digits | `validate:"alphanum_upper"` |
| `alphanum_lower` | Lowercase letters + digits | `validate:"alphanum_lower"` |

## Network & Address

| Tag | Description | Example |
|-----|-------------|---------|
| `ip` | IP address (v4 or v6) | `validate:"ip"` |
| `ipv4` | IPv4 address | `validate:"ipv4"` |
| `ipv6` | IPv6 address | `validate:"ipv6"` |
| `ip_addr` | IP address | `validate:"ip_addr"` |
| `ip4_addr` | IPv4 address | `validate:"ip4_addr"` |
| `ip6_addr` | IPv6 address | `validate:"ip6_addr"` |
| `cidr` | CIDR notation | `validate:"cidr"` |
| `cidrv4` | IPv4 CIDR | `validate:"cidrv4"` |
| `cidrv6` | IPv6 CIDR | `validate:"cidrv6"` |
| `tcp_addr` | TCP address | `validate:"tcp_addr"` |
| `tcp4_addr` | TCPv4 address | `validate:"tcp4_addr"` |
| `tcp6_addr` | TCPv6 address | `validate:"tcp6_addr"` |
| `udp_addr` | UDP address | `validate:"udp_addr"` |
| `udp4_addr` | UDPv4 address | `validate:"udp4_addr"` |
| `udp6_addr` | UDPv6 address | `validate:"udp6_addr"` |
| `unix_addr` | Unix domain socket address | `validate:"unix_addr"` |
| `uds_exists` | Unix domain socket exists | `validate:"uds_exists"` |
| `mac` | MAC address | `validate:"mac"` |
| `port` | Port number | `validate:"port"` |
| `hostname` | Hostname (RFC 952) | `validate:"hostname"` |
| `hostname_rfc1123` | Hostname (RFC 1123) | `validate:"hostname_rfc1123"` |
| `hostname_port` | Hostname:port | `validate:"hostname_port"` |
| `fqdn` | Fully Qualified Domain Name | `validate:"fqdn"` |
| `uri` | URI | `validate:"uri"` |
| `http_url` | HTTP(S) URL | `validate:"http_url"` |
| `https_url` | HTTPS-only URL | `validate:"https_url"` |
| `origin` | Web origin | `validate:"origin"` |
| `url_encoded` | URL encoded | `validate:"url_encoded"` |
| `datauri` | Data URI | `validate:"datauri"` |
| `urn_rfc2141` | URN (RFC 2141) | `validate:"urn_rfc2141"` |

## Numeric / Length

| Tag | Description | Example |
|-----|-------------|---------|
| `min=N` | Min value/length | `validate:"min=3"` |
| `max=N` | Max value/length | `validate:"max=100"` |
| `len=N` | Exact length | `validate:"len=11"` |
| `minlen=N` | Min length | `validate:"minlen=6"` |
| `maxlen=N` | Max length | `validate:"maxlen=20"` |
| `range=min,max` | Numeric range | `validate:"range=0.0,5.0"` |

## Collection / Pattern

| Tag | Description | Example |
|-----|-------------|---------|
| `in=v1,v2,...` | Enum values | `validate:"in=male,female"` |
| `notin=v1,v2,...` | Exclude enum | `validate:"notin=admin,root"` |
| `pattern=regex` | Regex match | `validate:"pattern=^[A-Z]"` |
| `containspecial` | Has special chars | `validate:"containspecial"` |
| `strong_password` | Strong password | `validate:"strong_password"` |

## Logical Composition

```go
// And — all must pass
combined := validator.And(
    validator.Required(),
    validator.MinLength(5),
)

// Or — any must pass
either := validator.Or(
    validator.Email(),
    validator.Pattern(`^\d+$`),
)

// Not — negate
negated := validator.Not(validator.In("a", "b"))
```
