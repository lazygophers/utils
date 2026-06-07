//go:build (lang_ko || lang_all) && (country_africa || country_all || country_northern_africa || country_sd || currency_all || currency_sdg)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sdg.RegisterName(xlanguage.Korean, "수단 파운드")
}
