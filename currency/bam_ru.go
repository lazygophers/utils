//go:build (lang_ru || lang_all) && (country_all || country_ba || country_europe || country_southern_europe || currency_all || currency_bam)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bam.RegisterName(xlanguage.Russian, "Конвертируемая марка")
}
