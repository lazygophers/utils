//go:build country_all || country_americas || country_caribbean || country_mq

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMartinique.RegisterName(xlanguage.English, "Martinique")
	dataMartinique.RegisterOfficialName(xlanguage.English, "Martinique")
	dataMartinique.RegisterCapital(xlanguage.English, "Fort-de-France")
}
