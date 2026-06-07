//go:build (lang_es || lang_all) && (country_all || country_asia || country_eastern_asia || country_mo)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMacao.RegisterName(xlanguage.Spanish, "Macao")
	dataMacao.RegisterOfficialName(xlanguage.Spanish, "Región Administrativa Especial de Macao")
	dataMacao.RegisterCapital(xlanguage.Spanish, "Macao")
}
