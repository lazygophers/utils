package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.English, "Georgia")
	dataGeorgia.RegisterOfficialName(xlanguage.English, "Georgia")
	dataGeorgia.RegisterCapital(xlanguage.English, "Tbilisi")
}
