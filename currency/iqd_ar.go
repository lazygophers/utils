//go:build (lang_ar || lang_all) && (country_all || country_asia || country_iq || country_western_asia || currency_all || currency_iqd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Iqd.RegisterName(xlanguage.Arabic, "دينار عراقي")
}
