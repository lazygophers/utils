package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.English, "Guatemala")
	dataGuatemala.RegisterOfficialName(xlanguage.English, "Republic of Guatemala")
	dataGuatemala.RegisterCapital(xlanguage.English, "Guatemala City")
}
