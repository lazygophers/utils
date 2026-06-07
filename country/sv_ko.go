//go:build (lang_ko || lang_all) && (country_all || country_americas || country_central_america || country_sv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.Korean, "엘살바도르")
	dataElSalvador.RegisterOfficialName(xlanguage.Korean, "엘살바도르 공화국")
	dataElSalvador.RegisterCapital(xlanguage.Korean, "산살바도르")
}
