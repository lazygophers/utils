//go:build (lang_es || lang_all) && (country_all || country_americas || country_caribbean || country_cw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuracao.RegisterName(xlanguage.Spanish, "Curazao")
	dataCuracao.RegisterOfficialName(xlanguage.Spanish, "País de Curazao")
	dataCuracao.RegisterCapital(xlanguage.Spanish, "Willemstad")
}
