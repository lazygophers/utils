package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMonaco.RegisterName(xlanguage.English, "Monaco")
	dataMonaco.RegisterOfficialName(xlanguage.English, "Principality of Monaco")
	dataMonaco.RegisterCapital(xlanguage.English, "Monaco")
}
