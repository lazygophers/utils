//go:build (lang_es || lang_all) && (country_all || country_melanesia || country_oceania || country_vu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVanuatu.RegisterName(xlanguage.Spanish, "Vanuatu")
	dataVanuatu.RegisterOfficialName(xlanguage.Spanish, "República de Vanuatu")
	dataVanuatu.RegisterCapital(xlanguage.Spanish, "Port Vila")
}
