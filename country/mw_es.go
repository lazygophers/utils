//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_mw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.Spanish, "Malaui")
	dataMalawi.RegisterOfficialName(xlanguage.Spanish, "República de Malaui")
	dataMalawi.RegisterCapital(xlanguage.Spanish, "Lilongüe")
}
