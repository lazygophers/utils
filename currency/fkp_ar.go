//go:build (lang_ar || lang_all) && (country_all || country_americas || country_fk || country_south_america || currency_all || currency_fkp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	FKP.RegisterName(xlanguage.Arabic, "جنيه فوكلاندي")
}
