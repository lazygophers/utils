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
| `json` | JSON 格式 | `validate:"json"` |
| `uuid` | UUID 格式 | `validate:"uuid"` |

## 字符串

| 標籤 | 說明 | 範例 |
|------|------|------|
| `alpha` | 僅字母 | `validate:"alpha"` |
| `alphaspace` | 僅字母 + 空格 | `validate:"alphaspace"` |
| `alphanum` | 僅字母 + 數字 | `validate:"alphanum"` |
| `alphanumspace` | 僅字母 + 數字 + 空格 | `validate:"alphanumspace"` |
| `alphanumunicode` | 字母 + 數字（含 Unicode） | `validate:"alphanumunicode"` |
| `alphaunicode` | 僅字母（含 Unicode） | `validate:"alphaunicode"` |
| `ascii` | 僅 ASCII 字元 | `validate:"ascii"` |
| `printascii` | 僅可列印 ASCII | `validate:"printascii"` |
| `multibyte` | 多位元組字元 | `validate:"multibyte"` |
| `boolean` | 布林值 | `validate:"boolean"` |
| `number` | 數字 | `validate:"number"` |
| `numeric` | 數值 | `validate:"numeric"` |
| `uppercase` | 僅大寫字母 | `validate:"uppercase"` |
| `lowercase` | 僅小寫字母 | `validate:"lowercase"` |
| `contains` | 包含子串 | `validate:"contains=abc"` |
| `containsany` | 包含任一字元 | `validate:"containsany=abc"` |
| `containsrune` | 包含指定 rune | `validate:"containsrune=©"` |
| `excludes` | 不包含子串 | `validate:"excludes=abc"` |
| `excludesall` | 不包含任一字元 | `validate:"excludesall=abc"` |
| `excludesrune` | 不包含指定 rune | `validate:"excludesrune=©"` |
| `startswith` | 以指定字串開頭 | `validate:"startswith=Hello"` |
| `startsnotwith` | 不以指定字串開頭 | `validate:"startsnotwith=Bad"` |
| `endswith` | 以指定字串結尾 | `validate:"endswith=World"` |
| `endsnotwith` | 不以指定字串結尾 | `validate:"endsnotwith=Bad"` |

## 字母/數字變體

| 標籤 | 說明 | 範例 |
|------|------|------|
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

## 雜湊/編碼

| 標籤 | 說明 | 範例 |
|------|------|------|
| `md4` | MD4 雜湊 | `validate:"md4"` |
| `md5` | MD5 雜湊 | `validate:"md5"` |
| `sha256` | SHA256 雜湊 | `validate:"sha256"` |
| `sha384` | SHA384 雜湊 | `validate:"sha384"` |
| `sha512` | SHA512 雜湊 | `validate:"sha512"` |
| `ripemd128` | RIPEMD-128 雜湊 | `validate:"ripemd128"` |
| `ripemd160` | RIPEMD-160 雜湊 | `validate:"ripemd160"` |
| `tiger128` | TIGER128 雜湊 | `validate:"tiger128"` |
| `tiger160` | TIGER160 雜湊 | `validate:"tiger160"` |
| `tiger192` | TIGER192 雜湊 | `validate:"tiger192"` |
| `hexadecimal` | 十六進位字串 | `validate:"hexadecimal"` |
| `base64` | Base64 字串 | `validate:"base64"` |
| `base64url` | Base64URL 字串 | `validate:"base64url"` |
| `base64rawurl` | Base64RawURL 字串 | `validate:"base64rawurl"` |

## UUID 變體

| 標籤 | 說明 | 範例 |
|------|------|------|
| `uuid_rfc4122` | UUID（RFC 4122） | `validate:"uuid_rfc4122"` |
| `uuid3` | UUID v3 | `validate:"uuid3"` |
| `uuid3_rfc4122` | UUID v3（RFC 4122） | `validate:"uuid3_rfc4122"` |
| `uuid4` | UUID v4 | `validate:"uuid4"` |
| `uuid4_rfc4122` | UUID v4（RFC 4122） | `validate:"uuid4_rfc4122"` |
| `uuid5` | UUID v5 | `validate:"uuid5"` |
| `uuid5_rfc4122` | UUID v5（RFC 4122） | `validate:"uuid5_rfc4122"` |

## 顏色

| 標籤 | 說明 | 範例 |
|------|------|------|
| `hexcolor` | 十六進位顏色 | `validate:"hexcolor"` |
| `rgb` | RGB 顏色 | `validate:"rgb"` |
| `rgba` | RGBA 顏色 | `validate:"rgba"` |
| `hsl` | HSL 顏色 | `validate:"hsl"` |
| `hsla` | HSLA 顏色 | `validate:"hsla"` |
| `cmyk` | CMYK 顏色 | `validate:"cmyk"` |

## 證件/編號

| 標籤 | 說明 | 範例 |
|------|------|------|
| `isbn` | ISBN 編號（10 或 13 位） | `validate:"isbn"` |
| `isbn10` | ISBN-10 | `validate:"isbn10"` |
| `isbn13` | ISBN-13 | `validate:"isbn13"` |
| `issn` | ISSN 編號 | `validate:"issn"` |
| `credit_card` | 信用卡號（Luhn 校驗） | `validate:"credit_card"` |
| `luhn_checksum` | Luhn 校驗和 | `validate:"luhn_checksum"` |
| `ein` | 美國雇主識別號 | `validate:"ein"` |
| `ssn` | 美國社會安全號 | `validate:"ssn"` |

## 地址/加密

| 標籤 | 說明 | 範例 |
|------|------|------|
| `btc_addr` | 比特幣地址 | `validate:"btc_addr"` |
| `btc_addr_bech32` | 比特幣 Bech32 地址 | `validate:"btc_addr_bech32"` |
| `eth_addr` | 以太坊地址 | `validate:"eth_addr"` |

## 地理/ISO

| 標籤 | 說明 | 範例 |
|------|------|------|
| `latitude` | 緯度（-90 ~ 90） | `validate:"latitude"` |
| `longitude` | 經度（-180 ~ 180） | `validate:"longitude"` |
| `timezone` | IANA 時區 | `validate:"timezone"` |
| `iso3166_1_alpha2` | ISO 3166-1 二字母國家代碼 | `validate:"iso3166_1_alpha2"` |
| `iso3166_1_alpha3` | ISO 3166-1 三字母國家代碼 | `validate:"iso3166_1_alpha3"` |
| `iso3166_1_alpha_numeric` | ISO 3166-1 數字國家代碼 | `validate:"iso3166_1_alpha_numeric"` |
| `iso3166_2` | ISO 3166-2 地區代碼 | `validate:"iso3166_2"` |
| `iso4217` | ISO 4217 貨幣代碼 | `validate:"iso4217"` |
| `postcode_iso3166_alpha2` | 郵遞區號 | `validate:"postcode_iso3166_alpha2"` |
| `postcode_iso3166_alpha2_field` | 郵遞區號（跨欄位） | `validate:"postcode_iso3166_alpha2_field=Country"` |

## 其他格式

| 標籤 | 說明 | 範例 |
|------|------|------|
| `semver` | 語義化版本號 | `validate:"semver"` |
| `ulid` | ULID | `validate:"ulid"` |
| `cve` | CVE 編號 | `validate:"cve"` |
| `jwt` | JSON Web Token | `validate:"jwt"` |
| `html` | 包含 HTML 標籤 | `validate:"html"` |
| `html_encoded` | HTML 編碼 | `validate:"html_encoded"` |
| `mongodb` | MongoDB ObjectID | `validate:"mongodb"` |
| `mongodb_connection_string` | MongoDB 連接字串 | `validate:"mongodb_connection_string"` |
| `cron` | Cron 運算式 | `validate:"cron"` |
| `spicedb` | SpiceDb 參照 | `validate:"spicedb"` |
| `datetime` | 日期時間 | `validate:"datetime=2006-01-02"` |
| `e164` | E.164 電話號碼 | `validate:"e164"` |
| `bic` | BIC 代碼（ISO 9362） | `validate:"bic"` |
| `bic_iso_9362_2014` | BIC 代碼（ISO 9362:2014） | `validate:"bic_iso_9362_2014"` |
| `bcp47_language_tag` | BCP 47 語言標籤 | `validate:"bcp47_language_tag"` |
| `bcp47_strict_language_tag` | BCP 47 語言標籤（嚴格） | `validate:"bcp47_strict_language_tag"` |

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

## 比較運算

| 標籤 | 說明 | 範例 |
|------|------|------|
| `gt=N` | 嚴格大於（值/長度） | `validate:"gt=5"` |
| `gte=N` | 大於等於（同 `min`） | `validate:"gte=5"` |
| `lt=N` | 嚴格小於（值/長度） | `validate:"lt=100"` |
| `lte=N` | 小於等於（同 `max`） | `validate:"lte=100"` |
| `eq_ignore_case=V` | 忽略大小寫相等 | `validate:"eq_ignore_case=hello"` |
| `ne_ignore_case=V` | 忽略大小寫不等 | `validate:"ne_ignore_case=hello"` |

## 檔案系統

| 標籤 | 說明 | 範例 |
|------|------|------|
| `dir` | 已存在的目錄 | `validate:"dir"` |
| `dirpath` | 合法的目錄路徑 | `validate:"dirpath"` |
| `file` | 已存在的檔案 | `validate:"file"` |
| `filepath` | 合法的檔案路徑 | `validate:"filepath"` |
| `image` | 圖片副檔名 | `validate:"image"` |

## 條件必填/排除

| 標籤 | 說明 | 範例 |
|------|------|------|
| `required_unless=F=V` | 除非指定欄位等於某值，否則必填 | `validate:"required_unless=Role=admin"` |
| `required_with_all=F1,F2` | 所有指定欄位有值時必填 | `validate:"required_with_all=FirstName,LastName"` |
| `required_without_all=F1,F2` | 所有指定欄位無值時必填 | `validate:"required_without_all=Email,Phone"` |
| `excluded_if=F=V` | 條件滿足時必須為空 | `validate:"excluded_if=Type=none"` |
| `excluded_unless=F=V` | 除非條件滿足，否則必須為空 | `validate:"excluded_unless=Active=true"` |
| `excluded_with=F` | 任一指定欄位有值時必須為空 | `validate:"excluded_with=Other"` |
| `excluded_with_all=F1,F2` | 所有指定欄位有值時必須為空 | `validate:"excluded_with_all=A,B"` |
| `excluded_without=F` | 任一指定欄位無值時必須為空 | `validate:"excluded_without=Email"` |
| `excluded_without_all=F1,F2` | 所有指定欄位無值時必須為空 | `validate:"excluded_without_all=A,B"` |

## 跨欄位比較

| 標籤 | 說明 | 範例 |
|------|------|------|
| `gtfield=F` | 大於指定欄位 | `validate:"gtfield=Min"` |
| `gtefield=F` | 大於等於指定欄位 | `validate:"gtefield=Min"` |
| `ltfield=F` | 小於指定欄位 | `validate:"ltfield=Max"` |
| `ltefield=F` | 小於等於指定欄位 | `validate:"ltefield=Max"` |
| `eqcsfield=S.F` | 等於跨結構體欄位（點分路徑） | `validate:"eqcsfield=Inner.Val"` |
| `necsfield=S.F` | 不等於跨結構體欄位 | `validate:"necsfield=Inner.Val"` |
| `gtcsfield=S.F` | 大於跨結構體欄位 | `validate:"gtcsfield=Inner.Val"` |
| `gtecsfield=S.F` | 大於等於跨結構體欄位 | `validate:"gtecsfield=Inner.Val"` |
| `ltcsfield=S.F` | 小於跨結構體欄位 | `validate:"ltcsfield=Inner.Val"` |
| `ltecsfield=S.F` | 小於等於跨結構體欄位 | `validate:"ltecsfield=Inner.Val"` |
| `fieldcontains=C` | 欄位值包含指定字元 | `validate:"fieldcontains=@"` |
| `fieldexcludes=C` | 欄位值不含指定字元 | `validate:"fieldexcludes=@"` |

## 雜項

| 標籤 | 說明 | 範例 |
|------|------|------|
| `oneof=A,B,C` | 列舉值之一 | `validate:"oneof=red,green,blue"` |
| `unique` | 切片/陣列元素唯一 | `validate:"unique"` |
| `isdefault` | 必須是零值 | `validate:"isdefault"` |
| `validateFn` | 呼叫 Validate() error 方法 | `validate:"validateFn"` |
| `iscolor` | 有效顏色（hex/rgb/rgba/hsl/hsla） | `validate:"iscolor"` |
| `country_code` | 有效國家代碼（ISO 3166-1） | `validate:"country_code"` |
