//go:build (lang_ko || lang_all) && (country_all || country_americas || country_gf || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.Korean, "프랑스령 기아나")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.Korean, "프랑스령 기아나")
	dataFrenchGuiana.RegisterCapital(xlanguage.Korean, "카옌")
}
