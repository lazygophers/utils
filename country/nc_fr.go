package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewCaledonia.RegisterName(xlanguage.French, "Nouvelle-Calédonie")
	dataNewCaledonia.RegisterOfficialName(xlanguage.French, "Nouvelle-Calédonie")
	dataNewCaledonia.RegisterCapital(xlanguage.French, "Nouméa")
}
