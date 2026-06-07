//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.Korean, "가봉")
	dataGabon.RegisterOfficialName(xlanguage.Korean, "가봉 공화국")
	dataGabon.RegisterCapital(xlanguage.Korean, "리브르빌")
}
