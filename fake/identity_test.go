package fake_test

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/fake"
)

type idCardCase struct {
	name    string
	c       *country.Country
	pattern string
}

func TestIdCard_Formats(t *testing.T) {
	cases := []idCardCase{
		{name: "CN", c: country.China, pattern: `^\d{17}[\dX]$`},
		{name: "US", c: country.UnitedStates, pattern: `^\d{3}-\d{2}-\d{4}$`},
		{name: "JP", c: country.Japan, pattern: `^\d{12}$`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := fake.New(tc.c, fake.WithSeed(42))
			re := regexp.MustCompile(tc.pattern)
			for i := 0; i < 20; i++ {
				id := f.IdCard()
				if !re.MatchString(id) {
					t.Fatalf("IdCard %q doesn't match %s", id, tc.pattern)
				}
			}
		})
	}
}

// cnIdCardWeights / cnIdCardCheckCodes mirror the values in cn.go so the test
// is independent of the source implementation.
var (
	cnIdCardWeightsTest = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	cnIdCardCheckTest   = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
)

func TestCnIdCard_Checksum(t *testing.T) {
	f := fake.New(country.China, fake.WithSeed(1))
	for i := 0; i < 100; i++ {
		id := f.IdCard()
		if len(id) != 18 {
			t.Fatalf("len(%q) != 18", id)
		}
		sum := 0
		for k := 0; k < 17; k++ {
			sum += int(id[k]-'0') * cnIdCardWeightsTest[k]
		}
		want := cnIdCardCheckTest[sum%11]
		if id[17] != want {
			t.Fatalf("id %q checksum mismatch: got %c want %c", id, id[17], want)
		}
	}
}

func TestCnIdCard_GenderParity(t *testing.T) {
	birth := time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)
	fMale := fake.New(country.China, fake.WithSeed(1), fake.WithGender(fake.GenderMale))
	fFemale := fake.New(country.China, fake.WithSeed(1), fake.WithGender(fake.GenderFemale))

	for i := 0; i < 30; i++ {
		idMale := fMale.IdCardOf(fake.GenderMale, birth)
		// digit 16 (0-indexed) is last sequence digit; parity encodes gender.
		// Male = odd, Female = even.
		last := int(idMale[16] - '0')
		if last%2 == 0 {
			t.Fatalf("male id %q seq tail %d is even", idMale, last)
		}
		idFemale := fFemale.IdCardOf(fake.GenderFemale, birth)
		last2 := int(idFemale[16] - '0')
		if last2%2 != 0 {
			t.Fatalf("female id %q seq tail %d is odd", idFemale, last2)
		}
	}
}

func TestJpMyNumber_CheckDigit(t *testing.T) {
	weights1 := []int{6, 5, 4, 3, 2, 7}
	weights2 := []int{6, 5, 4, 3, 2}
	f := fake.New(country.Japan, fake.WithSeed(2))
	for i := 0; i < 100; i++ {
		id := f.IdCard()
		if len(id) != 12 {
			t.Fatalf("len(%q) != 12", id)
		}
		sum := 0
		for k := 0; k < 6; k++ {
			sum += int(id[k]-'0') * weights1[k]
		}
		for k := 0; k < 5; k++ {
			sum += int(id[6+k]-'0') * weights2[k]
		}
		check := 11 - (sum % 11)
		if check >= 10 {
			check = 0
		}
		got := int(id[11] - '0')
		if got != check {
			t.Fatalf("my-number %q check mismatch: got %d want %d", id, got, check)
		}
	}
}

func TestUsSsn_NoReservedRanges(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(3))
	re := regexp.MustCompile(`^(\d{3})-(\d{2})-(\d{4})$`)
	for i := 0; i < 100; i++ {
		ssn := f.IdCard()
		m := re.FindStringSubmatch(ssn)
		if m == nil {
			t.Fatalf("ssn malformed: %q", ssn)
		}
		// area must not be 000 or 666 or 900-999.
		var area int
		for _, c := range m[1] {
			area = area*10 + int(c-'0')
		}
		if area == 0 || area == 666 || area >= 900 {
			t.Fatalf("ssn area %d in reserved block: %q", area, ssn)
		}
		// group != 00, serial != 0000.
		if m[2] == "00" {
			t.Fatalf("ssn group 00: %q", ssn)
		}
		if m[3] == "0000" {
			t.Fatalf("ssn serial 0000: %q", ssn)
		}
	}
}

type passportCase struct {
	name   string
	c      *country.Country
	prefix string
}

func TestPassportNo_Prefix(t *testing.T) {
	cases := []passportCase{
		{name: "CN", c: country.China, prefix: "E"},
		{name: "US", c: country.UnitedStates, prefix: "A"},
		{name: "JP", c: country.Japan, prefix: "TR"},
		{name: "AD-default", c: country.Get("AD"), prefix: "P"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.c == nil {
				t.Skip("country not available in current build")
			}
			f := fake.New(tc.c, fake.WithSeed(4))
			p := f.PassportNo()
			if !strings.HasPrefix(p, tc.prefix) {
				t.Fatalf("expected prefix %q, got %q", tc.prefix, p)
			}
			rest := strings.TrimPrefix(p, tc.prefix)
			if len(rest) != 8 {
				t.Fatalf("expected 8 digits after prefix, got %q", rest)
			}
			for _, r := range rest {
				if r < '0' || r > '9' {
					t.Fatalf("non-digit in passport number: %q", p)
				}
			}
		})
	}
}

func TestBirthday_Range(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(5))
	now := time.Now()
	for i := 0; i < 50; i++ {
		bd := f.Birthday(18, 65)
		age := now.Year() - bd.Year()
		if bd.YearDay() > now.YearDay() {
			age--
		}
		if age < 17 || age > 66 {
			t.Fatalf("age %d outside [18,65] approx range; bd=%v", age, bd)
		}
	}
}

func TestBirthday_SwapBounds(t *testing.T) {
	f := fake.New(country.UnitedStates, fake.WithSeed(6))
	bd := f.Birthday(40, 20) // reversed, should still produce valid date
	if bd.IsZero() {
		t.Fatal("Birthday returned zero")
	}
}

func TestIdCardOf_NoGenerator(t *testing.T) {
	// AD locale has no IdCardGen, so generic 12-digit id should be returned.
	c := country.Get("AD")
	if c == nil {
		t.Skip("AD country requires country_all build tag")
	}
	f := fake.New(c, fake.WithSeed(7))
	id := f.IdCard()
	if len(id) != 12 {
		t.Fatalf("expected 12-digit generic id, got %q (len %d)", id, len(id))
	}
	for _, r := range id {
		if r < '0' || r > '9' {
			t.Fatalf("non-digit in generic id: %q", id)
		}
	}
}
