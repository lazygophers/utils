//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_cu || currency_all || currency_cup)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	CUP.RegisterName(xlanguage.Russian, "Кубинское песо")
}
