//go:build (lang_es || lang_all) && (country_all || country_asia || country_sa || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaudiArabia.RegisterName(xlanguage.Spanish, "Arabia Saudita")
	dataSaudiArabia.RegisterOfficialName(xlanguage.Spanish, "Reino de Arabia Saudita")
	dataSaudiArabia.RegisterCapital(xlanguage.Spanish, "Riad")
}
