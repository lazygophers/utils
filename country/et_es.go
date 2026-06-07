//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_et)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEthiopia.RegisterName(xlanguage.Spanish, "Etiopía")
	dataEthiopia.RegisterOfficialName(xlanguage.Spanish, "República Federal Democrática de Etiopía")
	dataEthiopia.RegisterCapital(xlanguage.Spanish, "Adís Abeba")
}
