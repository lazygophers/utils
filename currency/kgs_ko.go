//go:build (lang_ko || lang_all) && (country_all || country_asia || country_central_asia || country_kg || currency_all || currency_kgs)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kgs.RegisterName(xlanguage.Korean, "키르기스스탄 솜")
}
