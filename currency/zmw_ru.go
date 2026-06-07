//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_zm || currency_all || currency_zmw)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Zmw.RegisterName(xlanguage.Russian, "Замбийская квача")
}
