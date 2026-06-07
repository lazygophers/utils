//go:build (lang_es || lang_all) && (country_africa || country_all || country_cv || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaboVerde.RegisterName(xlanguage.Spanish, "Cabo Verde")
	dataCaboVerde.RegisterOfficialName(xlanguage.Spanish, "República de Cabo Verde")
	dataCaboVerde.RegisterCapital(xlanguage.Spanish, "Praia")
}
