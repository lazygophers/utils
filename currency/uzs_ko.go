//go:build (lang_ko || lang_all) && (country_all || country_asia || country_central_asia || country_uz || currency_all || currency_uzs)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Uzs.RegisterName(xlanguage.Korean, "우즈베키스탄 숨")
}
