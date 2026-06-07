//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMartinique.RegisterName(xlanguage.Spanish, "Martinica")
	dataMartinique.RegisterOfficialName(xlanguage.Spanish, "Martinica")
	dataMartinique.RegisterCapital(xlanguage.Spanish, "Fort-de-France")
}
