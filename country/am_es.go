//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.Spanish, "Armenia")
	dataArmenia.RegisterOfficialName(xlanguage.Spanish, "República de Armenia")
	dataArmenia.RegisterCapital(xlanguage.Spanish, "Ereván")
}
