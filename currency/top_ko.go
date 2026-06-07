//go:build (lang_ko || lang_all) && (country_all || country_oceania || country_polynesia || country_to || currency_all || currency_top)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Top.RegisterName(xlanguage.Korean, "통가 파앙가")
}
