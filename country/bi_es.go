//go:build (lang_es || lang_all) && (country_africa || country_all || country_bi || country_eastern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurundi.RegisterName(xlanguage.Spanish, "Burundi")
	dataBurundi.RegisterOfficialName(xlanguage.Spanish, "República de Burundi")
	dataBurundi.RegisterCapital(xlanguage.Spanish, "Gitega")
}
