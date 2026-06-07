package country_test

import (
	"strings"
	"testing"

	xlanguage "golang.org/x/text/language"

	"github.com/lazygophers/utils/country"
	"github.com/lazygophers/utils/currency"
)

type getCase struct {
	name string
	in   string
	want *country.Country
}

func TestGet(t *testing.T) {
	cases := []getCase{
		{"upper-cn", "CN", country.China},
		{"lower-cn", "cn", country.China},
		{"empty", "", nil},
		{"unknown", "XX", nil},
		{"one-letter", "C", nil},
		{"three-letter", "CHN", nil},
		{"mixed-case", "Us", country.UnitedStates},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := country.Get(c.in)
			if got != c.want {
				t.Fatalf("Get(%q) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}

type alpha3Case struct {
	name string
	in   string
	want *country.Country
}

func TestGetByAlpha3(t *testing.T) {
	cases := []alpha3Case{
		{"upper", "CHN", country.China},
		{"lower", "chn", country.China},
		{"mixed", "Usa", country.UnitedStates},
		{"empty", "", nil},
		{"too-short", "CN", nil},
		{"too-long", "CHIN", nil},
		{"unknown", "ZZZ", nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := country.GetByAlpha3(c.in)
			if got != c.want {
				t.Fatalf("GetByAlpha3(%q) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}

func TestGetByNumeric(t *testing.T) {
	if got := country.GetByNumeric(156); got != country.China {
		t.Fatalf("GetByNumeric(156) = %v, want China", got)
	}
	if got := country.GetByNumeric(840); got != country.UnitedStates {
		t.Fatalf("GetByNumeric(840) = %v, want US", got)
	}
	if got := country.GetByNumeric(0); got != nil {
		t.Fatalf("GetByNumeric(0) = %v, want nil", got)
	}
	if got := country.GetByNumeric(9999); got != nil {
		t.Fatalf("GetByNumeric(9999) = %v, want nil", got)
	}
}

type nameLookupCase struct {
	name string
	in   string
	want *country.Country
}

func TestGetByName(t *testing.T) {
	cases := []nameLookupCase{
		{"en-common", "China", country.China},
		{"en-lower", "china", country.China},
		{"en-trim", "  China  ", country.China},
		{"en-official", "People's Republic of China", country.China},
		{"zh-common", "中国", country.China},
		{"empty", "", nil},
		{"unknown", "Atlantis", nil},
		{"whitespace-only", "   ", nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := country.GetByName(c.in)
			if got != c.want {
				t.Fatalf("GetByName(%q) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}

func TestFieldAccessors(t *testing.T) {
	cn := country.China
	if cn.Alpha2() != "CN" {
		t.Errorf("Alpha2: %q", cn.Alpha2())
	}
	if cn.Alpha3() != "CHN" {
		t.Errorf("Alpha3: %q", cn.Alpha3())
	}
	if cn.Numeric() != 156 {
		t.Errorf("Numeric: %d", cn.Numeric())
	}
	if cn.Continent() != "AS" {
		t.Errorf("Continent: %q", cn.Continent())
	}
	if cn.Region().RegionName() != "Asia" {
		t.Errorf("Region: %q", cn.Region().RegionName())
	}
	if cn.Subregion() != "Eastern Asia" {
		t.Errorf("Subregion: %q", cn.Subregion())
	}
	if cn.FlagEmoji() != "\U0001F1E8\U0001F1F3" {
		t.Errorf("FlagEmoji: %q", cn.FlagEmoji())
	}
	if cn.Currency() != currency.CNY {
		t.Errorf("Currency: %v", cn.Currency())
	}
	if cn.String() != cn.Alpha2() {
		t.Errorf("String != Alpha2")
	}
	if len(cn.CallingCodes()) == 0 {
		t.Errorf("CallingCodes empty")
	}
	if len(cn.Timezones()) == 0 {
		t.Errorf("Timezones empty")
	}
	if len(cn.Tlds()) == 0 {
		t.Errorf("Tlds empty")
	}
	if cn.OfficialLanguage() == (xlanguage.Tag{}) {
		t.Errorf("OfficialLanguage empty")
	}
	if len(cn.SpokenLanguages()) == 0 {
		t.Errorf("SpokenLanguages empty")
	}
}

func TestSliceFieldsReturnCopies(t *testing.T) {
	cn := country.China

	cc := cn.CallingCodes()
	if len(cc) > 0 {
		cc[0] = "MUTATED"
	}
	if country.China.CallingCodes()[0] == "MUTATED" {
		t.Errorf("CallingCodes leak")
	}

	tz := cn.Timezones()
	if len(tz) > 0 {
		tz[0] = "MUTATED"
	}
	if country.China.Timezones()[0] == "MUTATED" {
		t.Errorf("Timezones leak")
	}

	tlds := cn.Tlds()
	if len(tlds) > 0 {
		tlds[0] = "MUTATED"
	}
	if country.China.Tlds()[0] == "MUTATED" {
		t.Errorf("Tlds leak")
	}

	langs := cn.SpokenLanguages()
	if len(langs) > 0 {
		langs[0] = xlanguage.Make("xx")
	}
	if country.China.SpokenLanguages()[0] == xlanguage.Make("xx") {
		t.Errorf("SpokenLanguages leak")
	}
}

func TestStringEqualsAlpha2(t *testing.T) {
	for _, c := range country.List() {
		if c.String() != c.Alpha2() {
			t.Fatalf("%s: String=%q Alpha2=%q", c.Alpha2(), c.String(), c.Alpha2())
		}
	}
	if !strings.EqualFold(country.UnitedStates.String(), "US") {
		t.Errorf("US string mismatch")
	}
}
