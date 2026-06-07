package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.English, "Gabon")
	dataGabon.RegisterOfficialName(xlanguage.English, "Gabonese Republic")
	dataGabon.RegisterCapital(xlanguage.English, "Libreville")
}
