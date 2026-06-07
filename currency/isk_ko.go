//go:build (lang_ko || lang_all) && (country_all || country_europe || country_is || country_northern_europe || currency_all || currency_isk)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Isk.RegisterName(xlanguage.Korean, "아이슬란드 크로나")
}
