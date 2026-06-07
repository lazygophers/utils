package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurundi.RegisterName(xlanguage.French, "Burundi")
	dataBurundi.RegisterOfficialName(xlanguage.French, "République du Burundi")
	dataBurundi.RegisterCapital(xlanguage.French, "Gitega")
}
