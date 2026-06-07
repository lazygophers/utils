package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.French, "Suisse")
	dataSwitzerland.RegisterOfficialName(xlanguage.French, "Confédération suisse")
	dataSwitzerland.RegisterCapital(xlanguage.French, "Berne")
}
