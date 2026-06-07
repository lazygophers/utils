//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaoTomeAndPrincipe.RegisterName(xlanguage.Korean, "상투메 프린시페")
	dataSaoTomeAndPrincipe.RegisterOfficialName(xlanguage.Korean, "상투메 프린시페 민주 공화국")
	dataSaoTomeAndPrincipe.RegisterCapital(xlanguage.Korean, "상투메")
}
