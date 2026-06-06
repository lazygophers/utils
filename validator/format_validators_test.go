package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatValidators(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	// 哈希
	t.Run("md5", func(t *testing.T) {
		assert.NoError(t, v.Var("d41d8cd98f00b204e9800998ecf8427e", "md5"))
		assert.Error(t, v.Var("not-a-hash", "md5"))
	})
	t.Run("sha256", func(t *testing.T) {
		assert.NoError(t, v.Var("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", "sha256"))
		assert.Error(t, v.Var("not-a-hash", "sha256"))
	})
	t.Run("hexadecimal", func(t *testing.T) {
		assert.NoError(t, v.Var("deadbeef", "hexadecimal"))
		assert.Error(t, v.Var("nothex", "hexadecimal"))
	})

	// UUID 变体
	t.Run("uuid_rfc4122", func(t *testing.T) {
		assert.NoError(t, v.Var("550e8400-e29b-41d4-a716-446655440000", "uuid_rfc4122"))
		assert.Error(t, v.Var("550e8400-e29b-41d4-1716-446655440000", "uuid_rfc4122"))
	})
	t.Run("uuid4", func(t *testing.T) {
		assert.NoError(t, v.Var("550e8400-e29b-44d4-a716-446655440000", "uuid4"))
		assert.Error(t, v.Var("550e8400-e29b-11d4-a716-446655440000", "uuid4"))
	})
	t.Run("uuid5", func(t *testing.T) {
		assert.NoError(t, v.Var("550e8400-e29b-55d4-a716-446655440000", "uuid5"))
		assert.Error(t, v.Var("550e8400-e29b-44d4-a716-446655440000", "uuid5"))
	})

	// Base64
	t.Run("base64", func(t *testing.T) {
		assert.NoError(t, v.Var("SGVsbG8gV29ybGQ=", "base64"))
		assert.Error(t, v.Var("not!base64!", "base64"))
	})
	t.Run("base64url", func(t *testing.T) {
		assert.NoError(t, v.Var("SGVsbG8gV29ybGQ=", "base64url"))
		assert.Error(t, v.Var("not!valid!", "base64url"))
	})
	t.Run("base64rawurl", func(t *testing.T) {
		assert.NoError(t, v.Var("SGVsbG8gV29ybGQ", "base64rawurl"))
		assert.Error(t, v.Var("not!valid!", "base64rawurl"))
	})

	// 颜色
	t.Run("hexcolor", func(t *testing.T) {
		assert.NoError(t, v.Var("#ff0000", "hexcolor"))
		assert.NoError(t, v.Var("#f00", "hexcolor"))
		assert.Error(t, v.Var("ff0000", "hexcolor"))
	})
	t.Run("rgb", func(t *testing.T) {
		assert.NoError(t, v.Var("rgb(255, 128, 0)", "rgb"))
		assert.Error(t, v.Var("rgb(300, 0, 0)", "rgb"))
		assert.Error(t, v.Var("not-rgb", "rgb"))
	})
	t.Run("rgba", func(t *testing.T) {
		assert.NoError(t, v.Var("rgba(255, 128, 0, 0.5)", "rgba"))
		assert.Error(t, v.Var("rgba(255, 128, 0, 2)", "rgba"))
	})
	t.Run("hsl", func(t *testing.T) {
		assert.NoError(t, v.Var("hsl(180, 50%, 50%)", "hsl"))
		assert.Error(t, v.Var("hsl(370, 50%, 50%)", "hsl"))
	})
	t.Run("cmyk", func(t *testing.T) {
		assert.NoError(t, v.Var("cmyk(100%, 50%, 0%, 25%)", "cmyk"))
		assert.Error(t, v.Var("cmyk(110%, 0%, 0%, 0%)", "cmyk"))
	})

	// 证件/编号
	t.Run("isbn10", func(t *testing.T) {
		assert.NoError(t, v.Var("0471958697", "isbn10"))
		assert.NoError(t, v.Var("0-321-14653-0", "isbn10"))
		assert.Error(t, v.Var("1234567890", "isbn10"))
	})
	t.Run("isbn13", func(t *testing.T) {
		assert.NoError(t, v.Var("978-0-321-14653-3", "isbn13"))
		assert.Error(t, v.Var("978-0-321-14653-0", "isbn13"))
	})
	t.Run("isbn", func(t *testing.T) {
		assert.NoError(t, v.Var("0471958697", "isbn"))
		assert.Error(t, v.Var("not-isbn", "isbn"))
	})
	t.Run("issn", func(t *testing.T) {
		assert.NoError(t, v.Var("0317-8471", "issn"))
		assert.Error(t, v.Var("0000-0001", "issn"))
	})
	t.Run("credit_card", func(t *testing.T) {
		assert.NoError(t, v.Var("4111111111111111", "credit_card"))
		assert.Error(t, v.Var("4111111111111112", "credit_card"))
	})
	t.Run("luhn_checksum", func(t *testing.T) {
		assert.NoError(t, v.Var("79927398713", "luhn_checksum"))
		assert.Error(t, v.Var("79927398710", "luhn_checksum"))
	})
	t.Run("ein", func(t *testing.T) {
		assert.NoError(t, v.Var("12-3456789", "ein"))
		assert.Error(t, v.Var("123456789", "ein"))
	})
	t.Run("ssn", func(t *testing.T) {
		assert.NoError(t, v.Var("123-45-6789", "ssn"))
		assert.Error(t, v.Var("123456789", "ssn"))
	})

	// 地址/加密
	t.Run("btc_addr", func(t *testing.T) {
		assert.NoError(t, v.Var("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", "btc_addr"))
		assert.Error(t, v.Var("not-btc", "btc_addr"))
	})
	t.Run("eth_addr", func(t *testing.T) {
		assert.NoError(t, v.Var("0x71C7656EC7ab88b098defB751B7401B5f6d8976F", "eth_addr"))
		assert.Error(t, v.Var("0x123", "eth_addr"))
	})

	// 地理
	t.Run("latitude", func(t *testing.T) {
		assert.NoError(t, v.Var("37.7749", "latitude"))
		assert.Error(t, v.Var("91", "latitude"))
	})
	t.Run("longitude", func(t *testing.T) {
		assert.NoError(t, v.Var("-122.4194", "longitude"))
		assert.Error(t, v.Var("181", "longitude"))
	})
	t.Run("timezone", func(t *testing.T) {
		assert.NoError(t, v.Var("America/New_York", "timezone"))
		assert.Error(t, v.Var("Invalid/Timezone", "timezone"))
	})
	t.Run("iso3166_1_alpha2", func(t *testing.T) {
		assert.NoError(t, v.Var("US", "iso3166_1_alpha2"))
		assert.Error(t, v.Var("USA", "iso3166_1_alpha2"))
	})
	t.Run("iso3166_1_alpha3", func(t *testing.T) {
		assert.NoError(t, v.Var("USA", "iso3166_1_alpha3"))
		assert.Error(t, v.Var("US", "iso3166_1_alpha3"))
	})
	t.Run("iso4217", func(t *testing.T) {
		assert.NoError(t, v.Var("USD", "iso4217"))
		assert.Error(t, v.Var("US", "iso4217"))
	})

	// 其他格式
	t.Run("semver", func(t *testing.T) {
		assert.NoError(t, v.Var("1.2.3", "semver"))
		assert.NoError(t, v.Var("v1.0.0-alpha", "semver"))
		assert.Error(t, v.Var("1.2", "semver"))
	})
	t.Run("ulid", func(t *testing.T) {
		assert.NoError(t, v.Var("01H4G12PGZSVKJHF6A4RE5R6TY", "ulid"))
		assert.Error(t, v.Var("not-ulid", "ulid"))
	})
	t.Run("cve", func(t *testing.T) {
		assert.NoError(t, v.Var("CVE-2024-12345", "cve"))
		assert.Error(t, v.Var("2024-12345", "cve"))
	})
	t.Run("jwt", func(t *testing.T) {
		assert.NoError(t, v.Var("eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U", "jwt"))
		assert.Error(t, v.Var("not.jwt", "jwt"))
	})
	t.Run("html", func(t *testing.T) {
		assert.NoError(t, v.Var("<p>Hello</p>", "html"))
		assert.Error(t, v.Var("plain text", "html"))
	})
	t.Run("html_encoded", func(t *testing.T) {
		assert.NoError(t, v.Var("hello&amp;world", "html_encoded"))
		assert.Error(t, v.Var("hello world", "html_encoded"))
	})
	t.Run("mongodb", func(t *testing.T) {
		assert.NoError(t, v.Var("507f1f77bcf86cd799439011", "mongodb"))
		assert.Error(t, v.Var("not-mongo", "mongodb"))
	})
	t.Run("cron", func(t *testing.T) {
		assert.NoError(t, v.Var("*/5 * * * *", "cron"))
		assert.Error(t, v.Var("not-cron", "cron"))
	})
	t.Run("datetime", func(t *testing.T) {
		assert.NoError(t, v.Var("2024-01-15", "datetime=2006-01-02"))
		assert.Error(t, v.Var("not-a-date", "datetime=2006-01-02"))
	})
	t.Run("e164", func(t *testing.T) {
		assert.NoError(t, v.Var("+14155552671", "e164"))
		assert.Error(t, v.Var("14155552671", "e164"))
	})
	t.Run("bic", func(t *testing.T) {
		assert.NoError(t, v.Var("DEUTDEFF", "bic"))
		assert.NoError(t, v.Var("DEUTDEFF500", "bic"))
		assert.Error(t, v.Var("DE", "bic"))
	})
	t.Run("bcp47_language_tag", func(t *testing.T) {
		assert.NoError(t, v.Var("en", "bcp47_language_tag"))
		assert.NoError(t, v.Var("zh-Hans-CN", "bcp47_language_tag"))
		assert.Error(t, v.Var("123", "bcp47_language_tag"))
	})
	t.Run("spicedb", func(t *testing.T) {
		assert.NoError(t, v.Var("document/doc1#viewer", "spicedb"))
		assert.Error(t, v.Var("not valid!", "spicedb"))
	})
	t.Run("mongodb_connection_string", func(t *testing.T) {
		assert.NoError(t, v.Var("mongodb://localhost:27017", "mongodb_connection_string"))
		assert.NoError(t, v.Var("mongodb+srv://cluster.example.com", "mongodb_connection_string"))
		assert.Error(t, v.Var("mysql://localhost", "mongodb_connection_string"))
	})
	t.Run("postcode_iso3166_alpha2", func(t *testing.T) {
		assert.NoError(t, v.Var("12345", "postcode_iso3166_alpha2"))
		assert.Error(t, v.Var("X", "postcode_iso3166_alpha2"))
	})
	t.Run("sha384", func(t *testing.T) {
		assert.NoError(t, v.Var("38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b", "sha384"))
	})
	t.Run("sha512", func(t *testing.T) {
		assert.NoError(t, v.Var("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", "sha512"))
	})
	t.Run("btc_addr_bech32", func(t *testing.T) {
		assert.NoError(t, v.Var("bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4", "btc_addr_bech32"))
	})
	t.Run("iso3166_1_alpha_numeric", func(t *testing.T) {
		assert.NoError(t, v.Var("840", "iso3166_1_alpha_numeric"))
		assert.Error(t, v.Var("84", "iso3166_1_alpha_numeric"))
	})
	t.Run("iso3166_2", func(t *testing.T) {
		assert.NoError(t, v.Var("US-CA", "iso3166_2"))
		assert.Error(t, v.Var("US", "iso3166_2"))
	})
}
