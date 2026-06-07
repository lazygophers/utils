//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.Korean, "투르크메니스탄")
	dataTurkmenistan.RegisterOfficialName(xlanguage.Korean, "투르크메니스탄")
	dataTurkmenistan.RegisterCapital(xlanguage.Korean, "아시가바트")
}
