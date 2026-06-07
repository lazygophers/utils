//go:build (lang_ko || lang_all) && (country_all || country_melanesia || country_oceania || country_vu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVanuatu.RegisterName(xlanguage.Korean, "바누아투")
	dataVanuatu.RegisterOfficialName(xlanguage.Korean, "바누아투 공화국")
	dataVanuatu.RegisterCapital(xlanguage.Korean, "포트빌라")
}
