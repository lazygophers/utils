//go:build (lang_ko || lang_all) && (country_all || country_asia || country_id || country_south_eastern_asia || currency_all || currency_idr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Idr.RegisterName(xlanguage.Korean, "인도네시아 루피아")
}
