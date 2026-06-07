//go:build (lang_es || lang_all) && (country_all || country_am || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.Spanish, "Armenia")
	dataArmenia.RegisterOfficialName(xlanguage.Spanish, "República de Armenia")
	dataArmenia.RegisterCapital(xlanguage.Spanish, "Ereván")
}
