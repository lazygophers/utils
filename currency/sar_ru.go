//go:build (lang_ru || lang_all) && (country_all || country_asia || country_sa || country_western_asia || currency_all || currency_sar)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sar.RegisterName(xlanguage.Russian, "Саудовский риял")
}
