//go:build (lang_es || lang_all) && (country_africa || country_all || country_ga || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.Spanish, "Gabón")
	dataGabon.RegisterOfficialName(xlanguage.Spanish, "República Gabonesa")
	dataGabon.RegisterCapital(xlanguage.Spanish, "Libreville")
}
