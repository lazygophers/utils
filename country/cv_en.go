package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaboVerde.RegisterName(xlanguage.English, "Cabo Verde")
	dataCaboVerde.RegisterOfficialName(xlanguage.English, "Republic of Cabo Verde")
	dataCaboVerde.RegisterCapital(xlanguage.English, "Praia")
}
