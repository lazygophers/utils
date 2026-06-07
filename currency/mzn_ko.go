//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_mz || currency_all || currency_mzn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mzn.RegisterName(xlanguage.Korean, "모잠비크 메티칼")
}
