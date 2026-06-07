package country_test

import (
	"testing"

	"github.com/lazygophers/utils/country"
)

func TestListLength(t *testing.T) {
	got := len(country.List())
	if got < 12 {
		t.Fatalf("List length = %d, want >= 12 (default set)", got)
	}
}

func TestAllEntriesCrossLookup(t *testing.T) {
	for _, c := range country.List() {
		if got := country.Get(c.Alpha2()); got != c {
			t.Errorf("%s: Get(%q) mismatch", c.Alpha2(), c.Alpha2())
		}
		if got := country.GetByAlpha3(c.Alpha3()); got != c {
			t.Errorf("%s: GetByAlpha3(%q) mismatch", c.Alpha2(), c.Alpha3())
		}
		if got := country.GetByNumeric(c.Numeric()); got != c {
			t.Errorf("%s: GetByNumeric(%d) mismatch", c.Alpha2(), c.Numeric())
		}
	}
}

func TestConstantsMatchLookup(t *testing.T) {
	if country.China != country.Get("CN") {
		t.Error("China != Get(CN)")
	}
	if country.UnitedStates != country.Get("US") {
		t.Error("UnitedStates != Get(US)")
	}
	if country.Japan != country.Get("JP") {
		t.Error("Japan != Get(JP)")
	}
}

func TestListReturnsIndependentSlice(t *testing.T) {
	a := country.List()
	if len(a) == 0 {
		t.Fatal("empty")
	}
	first := a[0]
	a[0] = nil
	b := country.List()
	if b[0] != first {
		t.Errorf("List shares backing array; b[0]=%v want %v", b[0], first)
	}
}

func TestNoDuplicateCodes(t *testing.T) {
	alpha2 := make(map[string]bool, 249)
	alpha3 := make(map[string]bool, 249)
	numeric := make(map[int]bool, 249)
	for _, c := range country.List() {
		if alpha2[c.Alpha2()] {
			t.Errorf("duplicate alpha2: %s", c.Alpha2())
		}
		alpha2[c.Alpha2()] = true
		if alpha3[c.Alpha3()] {
			t.Errorf("duplicate alpha3: %s", c.Alpha3())
		}
		alpha3[c.Alpha3()] = true
		if numeric[c.Numeric()] {
			t.Errorf("duplicate numeric: %d", c.Numeric())
		}
		numeric[c.Numeric()] = true
	}
}

func TestListSortedByAlpha2(t *testing.T) {
	a := country.List()
	for i := 1; i < len(a); i++ {
		if a[i-1].Alpha2() >= a[i].Alpha2() {
			t.Fatalf("not sorted at %d: %s >= %s", i, a[i-1].Alpha2(), a[i].Alpha2())
		}
	}
}
