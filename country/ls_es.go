//go:build (lang_es || lang_all) && (country_africa || country_all || country_ls || country_southern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLesotho.RegisterName(xlanguage.Spanish, "Lesoto")
	dataLesotho.RegisterOfficialName(xlanguage.Spanish, "Reino de Lesoto")
	dataLesotho.RegisterCapital(xlanguage.Spanish, "Maseru")
}
