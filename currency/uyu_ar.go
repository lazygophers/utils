//go:build (lang_ar || lang_all) && (country_all || country_americas || country_south_america || country_uy || currency_all || currency_uyu)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Uyu.RegisterName(xlanguage.Arabic, "بيزو أوروغواي")
}
