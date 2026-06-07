package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthKorea.RegisterName(xlanguage.Korean, "대한민국")
	dataSouthKorea.RegisterOfficialName(xlanguage.Korean, "대한민국")
	dataSouthKorea.RegisterCapital(xlanguage.Korean, "서울")
}
