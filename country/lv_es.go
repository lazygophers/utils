//go:build (lang_es || lang_all) && (country_all || country_europe || country_lv || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.Spanish, "Letonia")
	dataLatvia.RegisterOfficialName(xlanguage.Spanish, "República de Letonia")
	dataLatvia.RegisterCapital(xlanguage.Spanish, "Riga")
}
