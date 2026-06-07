package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMartinique.RegisterName(xlanguage.French, "Martinique")
	dataMartinique.RegisterOfficialName(xlanguage.French, "Martinique")
	dataMartinique.RegisterCapital(xlanguage.French, "Fort-de-France")
}
