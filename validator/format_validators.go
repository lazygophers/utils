package validator

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 预编译正则表达式
var (
	hexcolorRegex    = regexp.MustCompile(`^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`)
	rgbRegex         = regexp.MustCompile(`^rgb\(\s*(\d{1,3})\s*,\s*(\d{1,3})\s*,\s*(\d{1,3})\s*\)$`)
	rgbaRegex        = regexp.MustCompile(`^rgba\(\s*(\d{1,3})\s*,\s*(\d{1,3})\s*,\s*(\d{1,3})\s*,\s*(\d*\.?\d+)\s*\)$`)
	hslRegex         = regexp.MustCompile(`^hsl\(\s*(\d{1,3})\s*,\s*(\d{1,3})%\s*,\s*(\d{1,3})%\s*\)$`)
	hslaRegex        = regexp.MustCompile(`^hsla\(\s*(\d{1,3})\s*,\s*(\d{1,3})%\s*,\s*(\d{1,3})%\s*,\s*(\d*\.?\d+)\s*\)$`)
	cmykRegex        = regexp.MustCompile(`^cmyk\(\s*(\d{1,3})%\s*,\s*(\d{1,3})%\s*,\s*(\d{1,3})%\s*,\s*(\d{1,3})%\s*\)$`)
	e164Regex        = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	einRegex         = regexp.MustCompile(`^\d{2}-\d{7}$`)
	ssnRegex         = regexp.MustCompile(`^\d{3}-\d{2}-\d{4}$`)
	ethAddrRegex     = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
	btcAddrRegex     = regexp.MustCompile(`^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$`)
	btcBech32Regex   = regexp.MustCompile(`^bc1[aqpzry9x8gf2tvdw0s3jn54khce6mua7l]{6,90}$`)
	semverRegex      = regexp.MustCompile(`^v?(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	ulidRegex        = regexp.MustCompile(`^[0-7][0-9A-HJKMNP-TV-Z]{25}$`)
	cveRegex         = regexp.MustCompile(`^CVE-\d{4}-\d{4,}$`)
	jwtRegex         = regexp.MustCompile(`^[A-Za-z0-9+_-]+\.[A-Za-z0-9+_-]+\.[A-Za-z0-9+_-]+$`)
	bicRegex         = regexp.MustCompile(`^[A-Z]{6}[A-Z0-9]{2}([A-Z0-9]{3})?$`)
	issnRegex        = regexp.MustCompile(`^\d{4}-\d{3}[\dXx]$`)
	isoAlpha2Regex   = regexp.MustCompile(`^[A-Z]{2}$`)
	isoAlpha3Regex   = regexp.MustCompile(`^[A-Z]{3}$`)
	isoNumericRegex  = regexp.MustCompile(`^\d{3}$`)
	iso31662Regex    = regexp.MustCompile(`^[A-Z]{2}-[A-Z0-9]{1,3}$`)
	iso4217Regex     = regexp.MustCompile(`^[A-Z]{3}$`)
	cronRegex        = regexp.MustCompile(`^\S+(\s+\S+){4,5}$`)
	htmlTagRegex     = regexp.MustCompile(`<[^>]+>`)
	htmlEncodedRegex = regexp.MustCompile(`&(?:[a-zA-Z]+|#\d+|#x[0-9a-fA-F]+);`)
	bcp47Regex       = regexp.MustCompile(`^[a-z]{2,3}(?:-[A-Za-z]{4})?(?:-[A-Z]{2})?(?:-[a-zA-Z0-9]{2,8})*$`)
	spicedbRegex     = regexp.MustCompile(`^[a-z][a-z0-9_]*(?:/[a-z][a-z0-9_]*)*#?[a-z][a-z0-9_]*$`)
	postcodeRegex    = regexp.MustCompile(`^[A-Z0-9 -]{2,10}$`)
	mongodbOIDRegex  = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)
	mongoConnRegex   = regexp.MustCompile(`^mongodb(\+srv)?://\S+`)
)

// isHexString 验证十六进制字符串的长度和字符
func isHexString(s string, length int) bool {
	if len(s) != length {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

// luhnCheck Luhn 算法校验
func luhnCheck(s string) bool {
	sum := 0
	nDigits := 0
	alternate := false
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if c < '0' || c > '9' {
			continue
		}
		n := int(c - '0')
		if alternate {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alternate = !alternate
		nDigits++
	}
	return nDigits >= 2 && sum%10 == 0
}

// ===== 哈希验证器 =====

func validateMD4(fl FieldLevel) bool         { return isHexString(fl.Field().String(), 32) }
func validateMD5(fl FieldLevel) bool         { return isHexString(fl.Field().String(), 32) }
func validateSHA256(fl FieldLevel) bool      { return isHexString(fl.Field().String(), 64) }
func validateSHA384(fl FieldLevel) bool      { return isHexString(fl.Field().String(), 96) }
func validateSHA512(fl FieldLevel) bool      { return isHexString(fl.Field().String(), 128) }
func validateRIPEMD128(fl FieldLevel) bool   { return isHexString(fl.Field().String(), 32) }
func validateRIPEMD160(fl FieldLevel) bool   { return isHexString(fl.Field().String(), 40) }
func validateTiger128(fl FieldLevel) bool    { return isHexString(fl.Field().String(), 32) }
func validateTiger160(fl FieldLevel) bool    { return isHexString(fl.Field().String(), 40) }
func validateTiger192(fl FieldLevel) bool    { return isHexString(fl.Field().String(), 48) }
func validateHexadecimal(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

// ===== UUID 变体验证器 =====

func validateUUIDRFC4122(fl FieldLevel) bool {
	s := fl.Field().String()
	if !validateUUID(fl) {
		return false
	}
	// RFC 4122 variant: position 19 must be 8, 9, a, b (or A, B)
	c := s[19]
	return c == '8' || c == '9' || c == 'a' || c == 'b' || c == 'A' || c == 'B'
}

func validateUUIDVersion(fl FieldLevel, version byte) bool {
	s := fl.Field().String()
	if !validateUUID(fl) {
		return false
	}
	return s[14] == version
}

func validateUUIDVersionRFC4122(fl FieldLevel, version byte) bool {
	s := fl.Field().String()
	if !validateUUID(fl) {
		return false
	}
	if s[14] != version {
		return false
	}
	c := s[19]
	return c == '8' || c == '9' || c == 'a' || c == 'b' || c == 'A' || c == 'B'
}

func validateUUID3(fl FieldLevel) bool          { return validateUUIDVersion(fl, '3') }
func validateUUID3RFC4122(fl FieldLevel) bool   { return validateUUIDVersionRFC4122(fl, '3') }
func validateUUID4(fl FieldLevel) bool          { return validateUUIDVersion(fl, '4') }
func validateUUID4RFC4122(fl FieldLevel) bool   { return validateUUIDVersionRFC4122(fl, '4') }
func validateUUID5(fl FieldLevel) bool          { return validateUUIDVersion(fl, '5') }
func validateUUID5RFC4122(fl FieldLevel) bool   { return validateUUIDVersionRFC4122(fl, '5') }

// ===== Base64 验证器 =====

func validateBase64(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

func validateBase64URL(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	_, err := base64.URLEncoding.DecodeString(s)
	return err == nil
}

func validateBase64RawURL(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	_, err := base64.RawURLEncoding.DecodeString(s)
	return err == nil
}

// ===== 颜色验证器 =====

func validateHexColor(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	return hexcolorRegex.MatchString(s)
}

func validateRGB(fl FieldLevel) bool {
	s := fl.Field().String()
	matches := rgbRegex.FindStringSubmatch(s)
	if matches == nil {
		return false
	}
	for i := 1; i <= 3; i++ {
		v, _ := strconv.Atoi(matches[i])
		if v > 255 {
			return false
		}
	}
	return true
}

func validateRGBA(fl FieldLevel) bool {
	s := fl.Field().String()
	matches := rgbaRegex.FindStringSubmatch(s)
	if matches == nil {
		return false
	}
	for i := 1; i <= 3; i++ {
		v, _ := strconv.Atoi(matches[i])
		if v > 255 {
			return false
		}
	}
	a, err := strconv.ParseFloat(matches[4], 64)
	if err != nil || a < 0 || a > 1 {
		return false
	}
	return true
}

func validateHSL(fl FieldLevel) bool {
	s := fl.Field().String()
	matches := hslRegex.FindStringSubmatch(s)
	if matches == nil {
		return false
	}
	h, _ := strconv.Atoi(matches[1])
	sl, _ := strconv.Atoi(matches[2])
	l, _ := strconv.Atoi(matches[3])
	return h <= 360 && sl <= 100 && l <= 100
}

func validateHSLA(fl FieldLevel) bool {
	s := fl.Field().String()
	matches := hslaRegex.FindStringSubmatch(s)
	if matches == nil {
		return false
	}
	h, _ := strconv.Atoi(matches[1])
	sl, _ := strconv.Atoi(matches[2])
	l, _ := strconv.Atoi(matches[3])
	a, err := strconv.ParseFloat(matches[4], 64)
	return h <= 360 && sl <= 100 && l <= 100 && err == nil && a >= 0 && a <= 1
}

func validateCMYK(fl FieldLevel) bool {
	s := fl.Field().String()
	matches := cmykRegex.FindStringSubmatch(s)
	if matches == nil {
		return false
	}
	for i := 1; i <= 4; i++ {
		v, _ := strconv.Atoi(matches[i])
		if v > 100 {
			return false
		}
	}
	return true
}

// ===== 证件/编号验证器 =====

func validateISBN10(fl FieldLevel) bool {
	s := strings.ReplaceAll(fl.Field().String(), "-", "")
	if len(s) != 10 {
		return false
	}
	sum := 0
	for i := 0; i < 9; i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
		sum += int(c-'0') * (10 - i)
	}
	last := s[9]
	var check int
	if last == 'X' || last == 'x' {
		check = 10
	} else if last >= '0' && last <= '9' {
		check = int(last - '0')
	} else {
		return false
	}
	sum += check
	return sum%11 == 0
}

func validateISBN13(fl FieldLevel) bool {
	s := strings.ReplaceAll(fl.Field().String(), "-", "")
	if len(s) != 13 {
		return false
	}
	sum := 0
	for i := 0; i < 12; i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
		weight := 1
		if i%2 == 1 {
			weight = 3
		}
		sum += int(c-'0') * weight
	}
	last := s[12]
	if last < '0' || last > '9' {
		return false
	}
	check := (10 - sum%10) % 10
	return int(last-'0') == check
}

func validateISBN(fl FieldLevel) bool {
	s := strings.ReplaceAll(fl.Field().String(), "-", "")
	if len(s) == 10 {
		return validateISBN10(fl)
	}
	if len(s) == 13 {
		return validateISBN13(fl)
	}
	return false
}

func validateISSN(fl FieldLevel) bool {
	s := fl.Field().String()
	if !issnRegex.MatchString(s) {
		return false
	}
	sum := 0
	// 前 4 位 (position 0-3)
	for i := 0; i < 4; i++ {
		sum += int(s[i]-'0') * (8 - i)
	}
	// 后 3 位 (position 5-7，跳过 position 4 的 dash)
	for i := 0; i < 3; i++ {
		sum += int(s[i+5]-'0') * (4 - i)
	}
	last := s[8]
	var check int
	if last == 'X' || last == 'x' {
		check = 10
	} else {
		check = int(last - '0')
	}
	return (sum+check)%11 == 0
}

func validateCreditCard(fl FieldLevel) bool {
	s := fl.Field().String()
	digits := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits = append(digits, s[i])
		}
	}
	return len(digits) >= 13 && luhnCheck(string(digits))
}

func validateLuhnChecksum(fl FieldLevel) bool {
	field := fl.Field()
	switch field.Kind() {
	case reflect.String:
		return luhnCheck(field.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return luhnCheck(fmt.Sprintf("%d", field.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return luhnCheck(fmt.Sprintf("%d", field.Uint()))
	default:
		return luhnCheck(field.String())
	}
}

func validateEIN(fl FieldLevel) bool {
	return einRegex.MatchString(fl.Field().String())
}

func validateSSN(fl FieldLevel) bool {
	return ssnRegex.MatchString(fl.Field().String())
}

// ===== 地址/加密验证器 =====

func validateBTCAddr(fl FieldLevel) bool {
	return btcAddrRegex.MatchString(fl.Field().String())
}

func validateBTCAddrBech32(fl FieldLevel) bool {
	return btcBech32Regex.MatchString(fl.Field().String())
}

func validateEthAddr(fl FieldLevel) bool {
	return ethAddrRegex.MatchString(fl.Field().String())
}

// ===== 地理/ISO 验证器 =====

func validateLatitude(fl FieldLevel) bool {
	lat, err := strconv.ParseFloat(fl.Field().String(), 64)
	return err == nil && lat >= -90 && lat <= 90
}

func validateLongitude(fl FieldLevel) bool {
	lon, err := strconv.ParseFloat(fl.Field().String(), 64)
	return err == nil && lon >= -180 && lon <= 180
}

func validateTimezone(fl FieldLevel) bool {
	_, err := time.LoadLocation(fl.Field().String())
	return err == nil
}

func validateISO3166Alpha2(fl FieldLevel) bool {
	return isoAlpha2Regex.MatchString(fl.Field().String())
}

func validateISO3166Alpha3(fl FieldLevel) bool {
	return isoAlpha3Regex.MatchString(fl.Field().String())
}

func validateISO3166Numeric(fl FieldLevel) bool {
	return isoNumericRegex.MatchString(fl.Field().String())
}

func validateISO31662(fl FieldLevel) bool {
	return iso31662Regex.MatchString(fl.Field().String())
}

func validateISO4217(fl FieldLevel) bool {
	return iso4217Regex.MatchString(fl.Field().String())
}

func validatePostcode(fl FieldLevel) bool {
	return postcodeRegex.MatchString(fl.Field().String())
}

func validatePostcodeField(fl FieldLevel) bool {
	countryField := fl.GetFieldByName(fl.Param())
	if !countryField.IsValid() {
		return false
	}
	country := countryField.String()
	postcode := fl.Field().String()
	if !postcodeRegex.MatchString(postcode) {
		return false
	}
	// 基础国家代码校验，具体国家邮编规则可扩展
	return isoAlpha2Regex.MatchString(strings.ToUpper(country))
}

// ===== 其他格式验证器 =====

func validateSemver(fl FieldLevel) bool {
	return semverRegex.MatchString(fl.Field().String())
}

func validateULID(fl FieldLevel) bool {
	return ulidRegex.MatchString(fl.Field().String())
}

func validateCVE(fl FieldLevel) bool {
	return cveRegex.MatchString(fl.Field().String())
}

func validateJWT(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	return jwtRegex.MatchString(s)
}

func validateHTML(fl FieldLevel) bool {
	return htmlTagRegex.MatchString(fl.Field().String())
}

func validateHTMLEncoded(fl FieldLevel) bool {
	return htmlEncodedRegex.MatchString(fl.Field().String())
}

func validateMongoDB(fl FieldLevel) bool {
	return mongodbOIDRegex.MatchString(fl.Field().String())
}

func validateMongoDBConnStr(fl FieldLevel) bool {
	return mongoConnRegex.MatchString(fl.Field().String())
}

func validateCron(fl FieldLevel) bool {
	return cronRegex.MatchString(strings.TrimSpace(fl.Field().String()))
}

func validateSpiceDb(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	return spicedbRegex.MatchString(s)
}

func validateDatetime(fl FieldLevel) bool {
	layout := fl.Param()
	if layout == "" {
		return false
	}
	_, err := time.Parse(layout, fl.Field().String())
	return err == nil
}

func validateE164(fl FieldLevel) bool {
	return e164Regex.MatchString(fl.Field().String())
}

func validateBIC(fl FieldLevel) bool {
	return bicRegex.MatchString(fl.Field().String())
}

func validateBICISO93622014(fl FieldLevel) bool {
	return bicRegex.MatchString(fl.Field().String())
}

func validateBCP47(fl FieldLevel) bool {
	return bcp47Regex.MatchString(fl.Field().String())
}

func validateBCP47Strict(fl FieldLevel) bool {
	return bcp47Regex.MatchString(fl.Field().String())
}

// FormatValidators 返回所有格式验证器注册表
func FormatValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		// 哈希
		"md4":         validateMD4,
		"md5":         validateMD5,
		"sha256":      validateSHA256,
		"sha384":      validateSHA384,
		"sha512":      validateSHA512,
		"ripemd128":   validateRIPEMD128,
		"ripemd160":   validateRIPEMD160,
		"tiger128":    validateTiger128,
		"tiger160":    validateTiger160,
		"tiger192":    validateTiger192,
		"hexadecimal": validateHexadecimal,
		// UUID 变体
		"uuid_rfc4122":    validateUUIDRFC4122,
		"uuid3":           validateUUID3,
		"uuid3_rfc4122":   validateUUID3RFC4122,
		"uuid4":           validateUUID4,
		"uuid4_rfc4122":   validateUUID4RFC4122,
		"uuid5":           validateUUID5,
		"uuid5_rfc4122":   validateUUID5RFC4122,
		// Base64
		"base64":       validateBase64,
		"base64url":    validateBase64URL,
		"base64rawurl": validateBase64RawURL,
		// 颜色
		"hexcolor": validateHexColor,
		"rgb":      validateRGB,
		"rgba":     validateRGBA,
		"hsl":      validateHSL,
		"hsla":     validateHSLA,
		"cmyk":     validateCMYK,
		// 证件/编号
		"isbn":           validateISBN,
		"isbn10":         validateISBN10,
		"isbn13":         validateISBN13,
		"issn":           validateISSN,
		"credit_card":    validateCreditCard,
		"luhn_checksum":  validateLuhnChecksum,
		"ein":            validateEIN,
		"ssn":            validateSSN,
		// 地址/加密
		"btc_addr":        validateBTCAddr,
		"btc_addr_bech32": validateBTCAddrBech32,
		"eth_addr":        validateEthAddr,
		// 地理/ISO
		"latitude":                     validateLatitude,
		"longitude":                    validateLongitude,
		"timezone":                     validateTimezone,
		"iso3166_1_alpha2":             validateISO3166Alpha2,
		"iso3166_1_alpha3":             validateISO3166Alpha3,
		"iso3166_1_alpha_numeric":      validateISO3166Numeric,
		"iso3166_2":                    validateISO31662,
		"iso4217":                      validateISO4217,
		"postcode_iso3166_alpha2":      validatePostcode,
		"postcode_iso3166_alpha2_field": validatePostcodeField,
		// 其他格式
		"semver":                    validateSemver,
		"ulid":                      validateULID,
		"cve":                       validateCVE,
		"jwt":                       validateJWT,
		"html":                      validateHTML,
		"html_encoded":              validateHTMLEncoded,
		"mongodb":                   validateMongoDB,
		"mongodb_connection_string": validateMongoDBConnStr,
		"cron":                      validateCron,
		"spicedb":                   validateSpiceDb,
		"datetime":                  validateDatetime,
		"e164":                      validateE164,
		"bic":                       validateBIC,
		"bic_iso_9362_2014":         validateBICISO93622014,
		"bcp47_language_tag":        validateBCP47,
		"bcp47_strict_language_tag": validateBCP47Strict,
	}
}
