//go:build (lang_fr || lang_all) && (country_all || country_americas || country_caribbean || country_cw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuracao.RegisterName(xlanguage.French, "Curaçao")
	dataCuracao.RegisterOfficialName(xlanguage.French, "Pays de Curaçao")
	dataCuracao.RegisterCapital(xlanguage.French, "Willemstad")
}
