//go:build (lang_ko || lang_all) && (country_all || country_asia || country_tr || country_western_asia || currency_all || currency_try)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Try.RegisterName(xlanguage.Korean, "터키 리라")
}
