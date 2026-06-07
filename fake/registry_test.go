package fake

import (
	"testing"

	"github.com/lazygophers/utils/country"
)

func TestRegister_NilPanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic on nil locale")
		}
	}()
	register(nil)
}

func TestRegister_NilCountryPanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic on nil country")
		}
	}()
	register(&Locale{Country: nil})
}

func TestRegister_DuplicatePanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic on duplicate register")
		}
	}()
	// CN is already registered by init(); re-registering must panic.
	register(&Locale{Country: country.China})
}

func TestLookupLocale_Nil(t *testing.T) {
	got := lookupLocale(nil)
	if got != nil {
		t.Fatalf("lookupLocale(nil) = %v, want nil", got)
	}
}

func TestLookupLocale_Known(t *testing.T) {
	l := lookupLocale(country.China)
	if l == nil {
		t.Fatal("CN locale missing")
	}
	if l.Country.Alpha2() != "CN" {
		t.Fatalf("locale country alpha2 = %s", l.Country.Alpha2())
	}
}
