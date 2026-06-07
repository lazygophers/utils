//go:build (lang_ar || lang_all) && (country_all || country_americas || country_bo || country_south_america || currency_all || currency_bob)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BOB.RegisterName(xlanguage.Arabic, "بوليفيانو")
}
