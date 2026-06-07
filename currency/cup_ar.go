//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_cu || currency_all || currency_cup)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Cup.RegisterName(xlanguage.Arabic, "بيزو كوبي")
}
