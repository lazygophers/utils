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
| `json` | JSON 格式 | `validate:"json"` |
| `uuid` | UUID 格式 | `validate:"uuid"` |

## 字符串

| 标签 | 说明 | 示例 |
|------|------|------|
| `alpha` | 仅字母 | `validate:"alpha"` |
| `alphaspace` | 仅字母 + 空格 | `validate:"alphaspace"` |
| `alphanum` | 仅字母 + 数字 | `validate:"alphanum"` |
| `alphanumspace` | 仅字母 + 数字 + 空格 | `validate:"alphanumspace"` |
| `alphanumunicode` | 字母 + 数字（含 Unicode） | `validate:"alphanumunicode"` |
| `alphaunicode` | 仅字母（含 Unicode） | `validate:"alphaunicode"` |
| `ascii` | 仅 ASCII 字符 | `validate:"ascii"` |
| `printascii` | 仅可打印 ASCII | `validate:"printascii"` |
| `multibyte` | 多字节字符 | `validate:"multibyte"` |
| `boolean` | 布尔值 | `validate:"boolean"` |
| `number` | 数字 | `validate:"number"` |
| `numeric` | 数值 | `validate:"numeric"` |
| `uppercase` | 仅大写字母 | `validate:"uppercase"` |
| `lowercase` | 仅小写字母 | `validate:"lowercase"` |
| `contains` | 包含子串 | `validate:"contains=abc"` |
| `containsany` | 包含任一字符 | `validate:"containsany=abc"` |
| `containsrune` | 包含指定 rune | `validate:"containsrune=©"` |
| `excludes` | 不包含子串 | `validate:"excludes=abc"` |
| `excludesall` | 不包含任一字符 | `validate:"excludesall=abc"` |
| `excludesrune` | 不包含指定 rune | `validate:"excludesrune=©"` |
| `startswith` | 以指定字符串开头 | `validate:"startswith=Hello"` |
| `startsnotwith` | 不以指定字符串开头 | `validate:"startsnotwith=Bad"` |
| `endswith` | 以指定字符串结尾 | `validate:"endswith=World"` |
| `endsnotwith` | 不以指定字符串结尾 | `validate:"endsnotwith=Bad"` |

## 字母/数字变体

| 标签 | 说明 | 示例 |
|------|------|------|
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

## 哈希/编码

| 标签 | 说明 | 示例 |
|------|------|------|
| `md4` | MD4 哈希 | `validate:"md4"` |
| `md5` | MD5 哈希 | `validate:"md5"` |
| `sha256` | SHA256 哈希 | `validate:"sha256"` |
| `sha384` | SHA384 哈希 | `validate:"sha384"` |
| `sha512` | SHA512 哈希 | `validate:"sha512"` |
| `ripemd128` | RIPEMD-128 哈希 | `validate:"ripemd128"` |
| `ripemd160` | RIPEMD-160 哈希 | `validate:"ripemd160"` |
| `tiger128` | TIGER128 哈希 | `validate:"tiger128"` |
| `tiger160` | TIGER160 哈希 | `validate:"tiger160"` |
| `tiger192` | TIGER192 哈希 | `validate:"tiger192"` |
| `hexadecimal` | 十六进制字符串 | `validate:"hexadecimal"` |
| `base64` | Base64 字符串 | `validate:"base64"` |
| `base64url` | Base64URL 字符串 | `validate:"base64url"` |
| `base64rawurl` | Base64RawURL 字符串 | `validate:"base64rawurl"` |

## UUID 变体

| 标签 | 说明 | 示例 |
|------|------|------|
| `uuid_rfc4122` | UUID（RFC 4122） | `validate:"uuid_rfc4122"` |
| `uuid3` | UUID v3 | `validate:"uuid3"` |
| `uuid3_rfc4122` | UUID v3（RFC 4122） | `validate:"uuid3_rfc4122"` |
| `uuid4` | UUID v4 | `validate:"uuid4"` |
| `uuid4_rfc4122` | UUID v4（RFC 4122） | `validate:"uuid4_rfc4122"` |
| `uuid5` | UUID v5 | `validate:"uuid5"` |
| `uuid5_rfc4122` | UUID v5（RFC 4122） | `validate:"uuid5_rfc4122"` |

## 颜色

| 标签 | 说明 | 示例 |
|------|------|------|
| `hexcolor` | 十六进制颜色 | `validate:"hexcolor"` |
| `rgb` | RGB 颜色 | `validate:"rgb"` |
| `rgba` | RGBA 颜色 | `validate:"rgba"` |
| `hsl` | HSL 颜色 | `validate:"hsl"` |
| `hsla` | HSLA 颜色 | `validate:"hsla"` |
| `cmyk` | CMYK 颜色 | `validate:"cmyk"` |

## 证件/编号

| 标签 | 说明 | 示例 |
|------|------|------|
| `isbn` | ISBN 编号（10 或 13 位） | `validate:"isbn"` |
| `isbn10` | ISBN-10 | `validate:"isbn10"` |
| `isbn13` | ISBN-13 | `validate:"isbn13"` |
| `issn` | ISSN 编号 | `validate:"issn"` |
| `credit_card` | 信用卡号（Luhn 校验） | `validate:"credit_card"` |
| `luhn_checksum` | Luhn 校验和 | `validate:"luhn_checksum"` |
| `ein` | 美国雇主识别号 | `validate:"ein"` |
| `ssn` | 美国社会安全号 | `validate:"ssn"` |

## 地址/加密

| 标签 | 说明 | 示例 |
|------|------|------|
| `btc_addr` | 比特币地址 | `validate:"btc_addr"` |
| `btc_addr_bech32` | 比特币 Bech32 地址 | `validate:"btc_addr_bech32"` |
| `eth_addr` | 以太坊地址 | `validate:"eth_addr"` |

## 地理/ISO

| 标签 | 说明 | 示例 |
|------|------|------|
| `latitude` | 纬度（-90 ~ 90） | `validate:"latitude"` |
| `longitude` | 经度（-180 ~ 180） | `validate:"longitude"` |
| `timezone` | IANA 时区 | `validate:"timezone"` |
| `iso3166_1_alpha2` | ISO 3166-1 二字母国家代码 | `validate:"iso3166_1_alpha2"` |
| `iso3166_1_alpha3` | ISO 3166-1 三字母国家代码 | `validate:"iso3166_1_alpha3"` |
| `iso3166_1_alpha_numeric` | ISO 3166-1 数字国家代码 | `validate:"iso3166_1_alpha_numeric"` |
| `iso3166_2` | ISO 3166-2 地区代码 | `validate:"iso3166_2"` |
| `iso4217` | ISO 4217 货币代码 | `validate:"iso4217"` |
| `postcode_iso3166_alpha2` | 邮政编码 | `validate:"postcode_iso3166_alpha2"` |
| `postcode_iso3166_alpha2_field` | 邮政编码（跨字段） | `validate:"postcode_iso3166_alpha2_field=Country"` |

## 其他格式

| 标签 | 说明 | 示例 |
|------|------|------|
| `semver` | 语义化版本号 | `validate:"semver"` |
| `ulid` | ULID | `validate:"ulid"` |
| `cve` | CVE 编号 | `validate:"cve"` |
| `jwt` | JSON Web Token | `validate:"jwt"` |
| `html` | 包含 HTML 标签 | `validate:"html"` |
| `html_encoded` | HTML 编码 | `validate:"html_encoded"` |
| `mongodb` | MongoDB ObjectID | `validate:"mongodb"` |
| `mongodb_connection_string` | MongoDB 连接字符串 | `validate:"mongodb_connection_string"` |
| `cron` | Cron 表达式 | `validate:"cron"` |
| `spicedb` | SpiceDb 引用 | `validate:"spicedb"` |
| `datetime` | 日期时间 | `validate:"datetime=2006-01-02"` |
| `e164` | E.164 电话号码 | `validate:"e164"` |
| `bic` | BIC 代码（ISO 9362） | `validate:"bic"` |
| `bic_iso_9362_2014` | BIC 代码（ISO 9362:2014） | `validate:"bic_iso_9362_2014"` |
| `bcp47_language_tag` | BCP 47 语言标签 | `validate:"bcp47_language_tag"` |
| `bcp47_strict_language_tag` | BCP 47 语言标签（严格） | `validate:"bcp47_strict_language_tag"` |

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

## 比较运算

| 标签 | 说明 | 示例 |
|------|------|------|
| `gt=N` | 严格大于（值/长度） | `validate:"gt=5"` |
| `gte=N` | 大于等于（同 `min`） | `validate:"gte=5"` |
| `lt=N` | 严格小于（值/长度） | `validate:"lt=100"` |
| `lte=N` | 小于等于（同 `max`） | `validate:"lte=100"` |
| `eq_ignore_case=V` | 忽略大小写相等 | `validate:"eq_ignore_case=hello"` |
| `ne_ignore_case=V` | 忽略大小写不等 | `validate:"ne_ignore_case=hello"` |

## 文件系统

| 标签 | 说明 | 示例 |
|------|------|------|
| `dir` | 已存在的目录 | `validate:"dir"` |
| `dirpath` | 合法的目录路径 | `validate:"dirpath"` |
| `file` | 已存在的文件 | `validate:"file"` |
| `filepath` | 合法的文件路径 | `validate:"filepath"` |
| `image` | 图片文件扩展名 | `validate:"image"` |

## 条件必填/排除

| 标签 | 说明 | 示例 |
|------|------|------|
| `required_unless=F=V` | 除非指定字段等于某值，否则必填 | `validate:"required_unless=Role=admin"` |
| `required_with_all=F1,F2` | 所有指定字段有值时必填 | `validate:"required_with_all=FirstName,LastName"` |
| `required_without_all=F1,F2` | 所有指定字段无值时必填 | `validate:"required_without_all=Email,Phone"` |
| `excluded_if=F=V` | 条件满足时必须为空 | `validate:"excluded_if=Type=none"` |
| `excluded_unless=F=V` | 除非条件满足，否则必须为空 | `validate:"excluded_unless=Active=true"` |
| `excluded_with=F` | 任一指定字段有值时必须为空 | `validate:"excluded_with=Other"` |
| `excluded_with_all=F1,F2` | 所有指定字段有值时必须为空 | `validate:"excluded_with_all=A,B"` |
| `excluded_without=F` | 任一指定字段无值时必须为空 | `validate:"excluded_without=Email"` |
| `excluded_without_all=F1,F2` | 所有指定字段无值时必须为空 | `validate:"excluded_without_all=A,B"` |

## 跨字段比较

| 标签 | 说明 | 示例 |
|------|------|------|
| `gtfield=F` | 大于指定字段 | `validate:"gtfield=Min"` |
| `gtefield=F` | 大于等于指定字段 | `validate:"gtefield=Min"` |
| `ltfield=F` | 小于指定字段 | `validate:"ltfield=Max"` |
| `ltefield=F` | 小于等于指定字段 | `validate:"ltefield=Max"` |
| `eqcsfield=S.F` | 等于跨结构体字段（点分路径） | `validate:"eqcsfield=Inner.Val"` |
| `necsfield=S.F` | 不等于跨结构体字段 | `validate:"necsfield=Inner.Val"` |
| `gtcsfield=S.F` | 大于跨结构体字段 | `validate:"gtcsfield=Inner.Val"` |
| `gtecsfield=S.F` | 大于等于跨结构体字段 | `validate:"gtecsfield=Inner.Val"` |
| `ltcsfield=S.F` | 小于跨结构体字段 | `validate:"ltcsfield=Inner.Val"` |
| `ltecsfield=S.F` | 小于等于跨结构体字段 | `validate:"ltecsfield=Inner.Val"` |
| `fieldcontains=C` | 字段值包含指定字符 | `validate:"fieldcontains=@"` |
| `fieldexcludes=C` | 字段值不含指定字符 | `validate:"fieldexcludes=@"` |

## 杂项

| 标签 | 说明 | 示例 |
|------|------|------|
| `oneof=A,B,C` | 枚举值之一 | `validate:"oneof=red,green,blue"` |
| `unique` | 切片/数组元素唯一 | `validate:"unique"` |
| `isdefault` | 必须是零值 | `validate:"isdefault"` |
| `validateFn` | 调用 Validate() error 方法 | `validate:"validateFn"` |
| `iscolor` | 有效颜色（hex/rgb/rgba/hsl/hsla） | `validate:"iscolor"` |
| `country_code` | 有效国家代码（ISO 3166-1） | `validate:"country_code"` |
