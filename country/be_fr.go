package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.French, "Belgique")
	dataBelgium.RegisterOfficialName(xlanguage.French, "Royaume de Belgique")
	dataBelgium.RegisterCapital(xlanguage.French, "Bruxelles")
}
