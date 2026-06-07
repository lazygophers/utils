//go:build (lang_ru || lang_all) && (country_all || country_asia || country_az || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAzerbaijan.RegisterName(xlanguage.Russian, "Азербайджан")
	dataAzerbaijan.RegisterOfficialName(xlanguage.Russian, "Азербайджанская Республика")
	dataAzerbaijan.RegisterCapital(xlanguage.Russian, "Баку")
}
