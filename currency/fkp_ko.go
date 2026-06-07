//go:build (lang_ko || lang_all) && (country_all || country_americas || country_fk || country_south_america || currency_all || currency_fkp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	FKP.RegisterName(xlanguage.Korean, "포클랜드 제도 파운드")
}
