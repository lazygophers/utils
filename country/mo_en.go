package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMacao.RegisterName(xlanguage.English, "Macao")
	dataMacao.RegisterOfficialName(xlanguage.English, "Macao Special Administrative Region of the People's Republic of China")
	dataMacao.RegisterCapital(xlanguage.English, "Macao")
}
