//go:build (lang_ko || lang_all) && (country_africa || country_all || country_sh || country_western_africa || currency_all || currency_shp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SHP.RegisterName(xlanguage.Korean, "세인트헬레나 파운드")
}
