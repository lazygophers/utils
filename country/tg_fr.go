package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.French, "Togo")
	dataTogo.RegisterOfficialName(xlanguage.French, "République togolaise")
	dataTogo.RegisterCapital(xlanguage.French, "Lomé")
}
