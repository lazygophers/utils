//go:build (lang_ja || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_vn || currency_all || currency_vnd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	VND.RegisterName(xlanguage.Japanese, "ドン")
}
