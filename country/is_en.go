package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIceland.RegisterName(xlanguage.English, "Iceland")
	dataIceland.RegisterOfficialName(xlanguage.English, "Iceland")
	dataIceland.RegisterCapital(xlanguage.English, "Reykjavik")
}
