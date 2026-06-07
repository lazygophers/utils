//go:build (lang_ko || lang_all) && (country_all || country_americas || country_ca || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.Korean, "캐나다")
	dataCanada.RegisterOfficialName(xlanguage.Korean, "캐나다")
	dataCanada.RegisterCapital(xlanguage.Korean, "오타와")
}
