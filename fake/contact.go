package fake

import (
	"fmt"
	"strings"
)

// emailDomains is the pool of public email-service hostnames used by
// [Faker.Email] when assembling synthetic addresses. The list mixes western
// and Greater-China consumer providers plus the IANA-reserved example.com so
// that fixtures stay non-routable in tests that accidentally send mail.
var emailDomains = []string{
	"gmail.com",
	"yahoo.com",
	"outlook.com",
	"hotmail.com",
	"icloud.com",
	"qq.com",
	"163.com",
	"126.com",
	"sina.com",
	"mail.com",
	"example.com",
}

// CallingCode returns the country's primary ITU-T E.164 calling code with the
// leading "+" included (e.g. "+86", "+1", "+81"). Returns an empty string when
// the country has no registered calling code.
func (f *Faker) CallingCode() string {
	codes := f.country.CallingCodes()
	if len(codes) == 0 {
		return ""
	}
	return codes[0]
}

// Phone returns a synthetic mobile phone number formatted according to the
// faker country's local convention. Supported country-specific shapes:
//
//   - CN: "+86 1XX-XXXX-XXXX"
//   - US: "+1 (XXX) XXX-XXXX"
//   - JP: "+81 X[X]-XXXX-XXXX"
//
// All other countries fall back to "<callingCode> XXXXXXXXXX" (ten random
// digits), and a country with no calling code yields a bare ten-digit string
// prefixed by a single space.
func (f *Faker) Phone() string {
	switch f.country.Alpha2() {
	case "CN":
		prefix := f.pickString(f.locale.PhonePrefixes)
		return fmt.Sprintf("+86 %s-%s-%s", prefix, f.randomDigits(4), f.randomDigits(4))
	case "US":
		area := f.pickString(f.locale.PhonePrefixes)
		return fmt.Sprintf("+1 (%s) %s-%s", area, f.randomDigits(3), f.randomDigits(4))
	case "JP":
		prefix := f.pickString(f.locale.PhonePrefixes)
		return fmt.Sprintf("+81 %s-%s-%s", prefix, f.randomDigits(4), f.randomDigits(4))
	default:
		return f.CallingCode() + " " + f.randomDigits(10)
	}
}

// Tel returns a synthetic fixed-line ("landline") phone number formatted in
// the faker country's local convention. Supported country-specific shapes:
//
//   - CN: "+86 0XX-XXXX-XXXX" using a major-city trunk prefix.
//   - US: "+1 (XXX) XXX-XXXX" reusing the mobile area-code pool, since the
//     North American Numbering Plan does not separate mobile and fixed lines.
//   - JP: "+81 X[X]-XXXX-XXXX" using a trunk area code.
//
// Other countries fall back to "<callingCode> XXXXXXXX" (eight random digits).
func (f *Faker) Tel() string {
	switch f.country.Alpha2() {
	case "CN":
		prefix := f.pickString(f.locale.LandlinePrefix)
		return fmt.Sprintf("+86 %s-%s-%s", prefix, f.randomDigits(4), f.randomDigits(4))
	case "US":
		area := f.pickString(f.locale.PhonePrefixes)
		return fmt.Sprintf("+1 (%s) %s-%s", area, f.randomDigits(3), f.randomDigits(4))
	case "JP":
		prefix := f.pickString(f.locale.LandlinePrefix)
		return fmt.Sprintf("+81 %s-%s-%s", prefix, f.randomDigits(4), f.randomDigits(4))
	default:
		return f.CallingCode() + " " + f.randomDigits(8)
	}
}

// Email returns a synthetic RFC 5322 mailbox composed of a lower-case local
// part — one to two english tokens optionally followed by a numeric suffix —
// joined to a domain drawn from [emailDomains]. The address is never
// guaranteed to be deliverable; "example.com" is included in the domain pool
// to satisfy callers that need RFC 2606 compliance.
func (f *Faker) Email() string {
	first := strings.ToLower(pickEnglishToken(f, true))
	last := strings.ToLower(pickEnglishToken(f, false))

	var local strings.Builder
	local.Grow(len(first) + len(last) + 6)
	local.WriteString(sanitizeEmailLocal(first))
	if last != "" {
		if local.Len() > 0 {
			local.WriteByte('.')
		}
		local.WriteString(sanitizeEmailLocal(last))
	}
	if local.Len() == 0 {
		local.WriteString("user")
	}
	if f.intN(2) == 0 {
		local.WriteByte('.')
		local.WriteString(f.randomDigits(2 + f.intN(3)))
	}

	domain := f.pickString(emailDomains)
	return local.String() + "@" + domain
}

// randomDigits returns a string of exactly n decimal digits ('0'-'9'). When
// n <= 0 the empty string is returned.
func (f *Faker) randomDigits(n int) string {
	if n <= 0 {
		return ""
	}
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < n; i++ {
		b.WriteByte(byte('0' + f.intN(10)))
	}
	return b.String()
}

// pickEnglishToken draws an ASCII-friendly word for use as an email local
// part. When first is true a male first name is sampled (any-gender pool
// reuse is intentional — case is normalised by the caller), otherwise a
// surname is drawn. Falls back to a lorem token when the english pools are
// somehow empty.
func pickEnglishToken(f *Faker, first bool) string {
	if first {
		if len(englishMaleFirstNames) > 0 {
			return f.pickString(englishMaleFirstNames)
		}
	} else {
		if len(englishLastNames) > 0 {
			return f.pickString(englishLastNames)
		}
	}
	return f.pickString(loremWords)
}

// sanitizeEmailLocal restricts the input to characters safe in the local part
// of an email address — lower-case letters and digits. All other code points
// are dropped. Callers are expected to pre-lowercase the input.
func sanitizeEmailLocal(s string) string {
	if s == "" {
		return ""
	}
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
		}
	}
	return b.String()
}
