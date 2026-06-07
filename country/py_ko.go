//go:build (lang_ko || lang_all) && (country_all || country_americas || country_py || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataParaguay.RegisterName(xlanguage.Korean, "파라과이")
	dataParaguay.RegisterOfficialName(xlanguage.Korean, "파라과이 공화국")
	dataParaguay.RegisterCapital(xlanguage.Korean, "아순시온")
}
