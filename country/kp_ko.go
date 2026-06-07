package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthKorea.RegisterName(xlanguage.Korean, "조선민주주의인민공화국")
	dataNorthKorea.RegisterOfficialName(xlanguage.Korean, "조선민주주의인민공화국")
	dataNorthKorea.RegisterCapital(xlanguage.Korean, "평양")
}
