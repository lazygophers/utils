//go:build (lang_ko || lang_all) && (country_all || country_oceania || country_pf || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.Korean, "프랑스령 폴리네시아")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.Korean, "프랑스령 폴리네시아")
	dataFrenchPolynesia.RegisterCapital(xlanguage.Korean, "파페에테")
}
