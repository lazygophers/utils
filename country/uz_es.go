//go:build (lang_es || lang_all) && (country_all || country_asia || country_central_asia || country_uz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUzbekistan.RegisterName(xlanguage.Spanish, "Uzbekistán")
	dataUzbekistan.RegisterOfficialName(xlanguage.Spanish, "República de Uzbekistán")
	dataUzbekistan.RegisterCapital(xlanguage.Spanish, "Taskent")
}
