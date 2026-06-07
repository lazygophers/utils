//go:build (lang_ko || lang_all) && (country_africa || country_all || country_middle_africa || country_st)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaoTomeAndPrincipe.RegisterName(xlanguage.Korean, "상투메 프린시페")
	dataSaoTomeAndPrincipe.RegisterOfficialName(xlanguage.Korean, "상투메 프린시페 민주 공화국")
	dataSaoTomeAndPrincipe.RegisterCapital(xlanguage.Korean, "상투메")
}
