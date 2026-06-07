//go:build (lang_ja || lang_all) && (country_af || country_all || country_asia || country_southern_asia || currency_afn || currency_all)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Afn.RegisterName(xlanguage.Japanese, "アフガニ")
}
