//go:build (lang_ru || lang_all) && (country_all || country_americas || country_co || country_south_america || currency_all || currency_cop)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	COP.RegisterName(xlanguage.Russian, "Колумбийское песо")
}
