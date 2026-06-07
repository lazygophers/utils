//go:build (lang_es || lang_all) && (country_all || country_asia || country_bt || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBhutan.RegisterName(xlanguage.Spanish, "Bután")
	dataBhutan.RegisterOfficialName(xlanguage.Spanish, "Reino de Bután")
	dataBhutan.RegisterCapital(xlanguage.Spanish, "Timbu")
}
