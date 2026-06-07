package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuracao.RegisterName(xlanguage.English, "Curacao")
	dataCuracao.RegisterOfficialName(xlanguage.English, "Country of Curacao")
	dataCuracao.RegisterCapital(xlanguage.English, "Willemstad")
}
