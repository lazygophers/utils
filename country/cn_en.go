package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChina.RegisterName(xlanguage.English, "China")
	dataChina.RegisterOfficialName(xlanguage.English, "People's Republic of China")
	dataChina.RegisterCapital(xlanguage.English, "Beijing")
}
