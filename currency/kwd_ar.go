//go:build (lang_ar || lang_all) && (country_all || country_asia || country_kw || country_western_asia || currency_all || currency_kwd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kwd.RegisterName(xlanguage.Arabic, "دينار كويتي")
}
