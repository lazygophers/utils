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

## Hash / Encoding

| Tag | Description | Example |
|-----|-------------|---------|
| `md4` | MD4 hash | `validate:"md4"` |
| `md5` | MD5 hash | `validate:"md5"` |
| `sha256` | SHA256 hash | `validate:"sha256"` |
| `sha384` | SHA384 hash | `validate:"sha384"` |
| `sha512` | SHA512 hash | `validate:"sha512"` |
| `ripemd128` | RIPEMD-128 hash | `validate:"ripemd128"` |
| `ripemd160` | RIPEMD-160 hash | `validate:"ripemd160"` |
| `tiger128` | TIGER128 hash | `validate:"tiger128"` |
| `tiger160` | TIGER160 hash | `validate:"tiger160"` |
| `tiger192` | TIGER192 hash | `validate:"tiger192"` |
| `hexadecimal` | Hexadecimal string | `validate:"hexadecimal"` |
| `base64` | Base64 string | `validate:"base64"` |
| `base64url` | Base64URL string | `validate:"base64url"` |
| `base64rawurl` | Base64RawURL string | `validate:"base64rawurl"` |

## UUID Variants

| Tag | Description | Example |
|-----|-------------|---------|
| `uuid_rfc4122` | UUID (RFC 4122) | `validate:"uuid_rfc4122"` |
| `uuid3` | UUID v3 | `validate:"uuid3"` |
| `uuid3_rfc4122` | UUID v3 (RFC 4122) | `validate:"uuid3_rfc4122"` |
| `uuid4` | UUID v4 | `validate:"uuid4"` |
| `uuid4_rfc4122` | UUID v4 (RFC 4122) | `validate:"uuid4_rfc4122"` |
| `uuid5` | UUID v5 | `validate:"uuid5"` |
| `uuid5_rfc4122` | UUID v5 (RFC 4122) | `validate:"uuid5_rfc4122"` |

## Colors

| Tag | Description | Example |
|-----|-------------|---------|
| `hexcolor` | Hex color | `validate:"hexcolor"` |
| `rgb` | RGB color | `validate:"rgb"` |
| `rgba` | RGBA color | `validate:"rgba"` |
| `hsl` | HSL color | `validate:"hsl"` |
| `hsla` | HSLA color | `validate:"hsla"` |
| `cmyk` | CMYK color | `validate:"cmyk"` |

## Identification

| Tag | Description | Example |
|-----|-------------|---------|
| `isbn` | ISBN number (10 or 13 digit) | `validate:"isbn"` |
| `isbn10` | ISBN-10 | `validate:"isbn10"` |
| `isbn13` | ISBN-13 | `validate:"isbn13"` |
| `issn` | ISSN number | `validate:"issn"` |
| `credit_card` | Credit card number (Luhn check) | `validate:"credit_card"` |
| `luhn_checksum` | Luhn checksum | `validate:"luhn_checksum"` |
| `ein` | U.S. Employer Identification Number | `validate:"ein"` |
| `ssn` | Social Security Number | `validate:"ssn"` |

## Address / Crypto

| Tag | Description | Example |
|-----|-------------|---------|
| `btc_addr` | Bitcoin address | `validate:"btc_addr"` |
| `btc_addr_bech32` | Bitcoin Bech32 address | `validate:"btc_addr_bech32"` |
| `eth_addr` | Ethereum address | `validate:"eth_addr"` |

## Geographic / ISO

| Tag | Description | Example |
|-----|-------------|---------|
| `latitude` | Latitude (-90 to 90) | `validate:"latitude"` |
| `longitude` | Longitude (-180 to 180) | `validate:"longitude"` |
| `timezone` | IANA timezone | `validate:"timezone"` |
| `iso3166_1_alpha2` | ISO 3166-1 alpha-2 country code | `validate:"iso3166_1_alpha2"` |
| `iso3166_1_alpha3` | ISO 3166-1 alpha-3 country code | `validate:"iso3166_1_alpha3"` |
| `iso3166_1_alpha_numeric` | ISO 3166-1 numeric country code | `validate:"iso3166_1_alpha_numeric"` |
| `iso3166_2` | ISO 3166-2 subdivision code | `validate:"iso3166_2"` |
| `iso4217` | ISO 4217 currency code | `validate:"iso4217"` |
| `postcode_iso3166_alpha2` | Postcode | `validate:"postcode_iso3166_alpha2"` |
| `postcode_iso3166_alpha2_field` | Postcode (cross-field) | `validate:"postcode_iso3166_alpha2_field=Country"` |

## Other Formats

| Tag | Description | Example |
|-----|-------------|---------|
| `semver` | Semantic version | `validate:"semver"` |
| `ulid` | ULID | `validate:"ulid"` |
| `cve` | CVE identifier | `validate:"cve"` |
| `jwt` | JSON Web Token | `validate:"jwt"` |
| `html` | Contains HTML tags | `validate:"html"` |
| `html_encoded` | HTML encoded | `validate:"html_encoded"` |
| `mongodb` | MongoDB ObjectID | `validate:"mongodb"` |
| `mongodb_connection_string` | MongoDB connection string | `validate:"mongodb_connection_string"` |
| `cron` | Cron expression | `validate:"cron"` |
| `spicedb` | SpiceDb reference | `validate:"spicedb"` |
| `datetime` | Datetime | `validate:"datetime=2006-01-02"` |
| `e164` | E.164 phone number | `validate:"e164"` |
| `bic` | BIC (ISO 9362) | `validate:"bic"` |
| `bic_iso_9362_2014` | BIC (ISO 9362:2014) | `validate:"bic_iso_9362_2014"` |
| `bcp47_language_tag` | BCP 47 language tag | `validate:"bcp47_language_tag"` |
| `bcp47_strict_language_tag` | BCP 47 language tag (strict) | `validate:"bcp47_strict_language_tag"` |

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

## Comparisons

| Tag | Description | Example |
|-----|-------------|---------|
| `gt=N` | Strictly greater than (value/length) | `validate:"gt=5"` |
| `gte=N` | Greater than or equal (same as `min`) | `validate:"gte=5"` |
| `lt=N` | Strictly less than (value/length) | `validate:"lt=100"` |
| `lte=N` | Less than or equal (same as `max`) | `validate:"lte=100"` |
| `eq_ignore_case=V` | Equals ignoring case | `validate:"eq_ignore_case=hello"` |
| `ne_ignore_case=V` | Not equal ignoring case | `validate:"ne_ignore_case=hello"` |

## File System

| Tag | Description | Example |
|-----|-------------|---------|
| `dir` | Existing directory | `validate:"dir"` |
| `dirpath` | Valid directory path | `validate:"dirpath"` |
| `file` | Existing file | `validate:"file"` |
| `filepath` | Valid file path | `validate:"filepath"` |
| `image` | Image file extension | `validate:"image"` |

## Conditional Required / Excluded

| Tag | Description | Example |
|-----|-------------|---------|
| `required_unless=F=V` | Required unless field equals value | `validate:"required_unless=Role=admin"` |
| `required_with_all=F1,F2` | Required when all fields have values | `validate:"required_with_all=FirstName,LastName"` |
| `required_without_all=F1,F2` | Required when all fields are empty | `validate:"required_without_all=Email,Phone"` |
| `excluded_if=F=V` | Must be empty when condition met | `validate:"excluded_if=Type=none"` |
| `excluded_unless=F=V` | Must be empty unless condition met | `validate:"excluded_unless=Active=true"` |
| `excluded_with=F` | Must be empty when any field has value | `validate:"excluded_with=Other"` |
| `excluded_with_all=F1,F2` | Must be empty when all fields have values | `validate:"excluded_with_all=A,B"` |
| `excluded_without=F` | Must be empty when any field is empty | `validate:"excluded_without=Email"` |
| `excluded_without_all=F1,F2` | Must be empty when all fields are empty | `validate:"excluded_without_all=A,B"` |

## Cross-Field Comparison

| Tag | Description | Example |
|-----|-------------|---------|
| `gtfield=F` | Greater than another field | `validate:"gtfield=Min"` |
| `gtefield=F` | Greater than or equal to another field | `validate:"gtefield=Min"` |
| `ltfield=F` | Less than another field | `validate:"ltfield=Max"` |
| `ltefield=F` | Less than or equal to another field | `validate:"ltefield=Max"` |
| `eqcsfield=S.F` | Equals cross-struct field (dotted path) | `validate:"eqcsfield=Inner.Val"` |
| `necsfield=S.F` | Not equal cross-struct field | `validate:"necsfield=Inner.Val"` |
| `gtcsfield=S.F` | Greater than cross-struct field | `validate:"gtcsfield=Inner.Val"` |
| `gtecsfield=S.F` | Greater than or equal cross-struct field | `validate:"gtecsfield=Inner.Val"` |
| `ltcsfield=S.F` | Less than cross-struct field | `validate:"ltcsfield=Inner.Val"` |
| `ltecsfield=S.F` | Less than or equal cross-struct field | `validate:"ltecsfield=Inner.Val"` |
| `fieldcontains=C` | Field value contains specified chars | `validate:"fieldcontains=@"` |
| `fieldexcludes=C` | Field value excludes specified chars | `validate:"fieldexcludes=@"` |

## Miscellaneous

| Tag | Description | Example |
|-----|-------------|---------|
| `oneof=A,B,C` | One of enumerated values | `validate:"oneof=red,green,blue"` |
| `unique` | Slice/array elements are unique | `validate:"unique"` |
| `isdefault` | Must be the zero value | `validate:"isdefault"` |
| `validateFn` | Calls Validate() error method | `validate:"validateFn"` |
| `iscolor` | Valid color (hex/rgb/rgba/hsl/hsla) | `validate:"iscolor"` |
| `country_code` | Valid country code (ISO 3166-1) | `validate:"country_code"` |
