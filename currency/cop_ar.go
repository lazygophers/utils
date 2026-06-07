//go:build (lang_ar || lang_all) && (country_all || country_americas || country_co || country_south_america || currency_all || currency_cop)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	COP.RegisterName(xlanguage.Arabic, "بيزو كولومبي")
}
