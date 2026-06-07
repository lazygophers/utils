//go:build (lang_ar || lang_all) && (country_all || country_americas || country_aw || country_caribbean || currency_all || currency_awg)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Awg.RegisterName(xlanguage.Arabic, "فلورن أروبي")
}
