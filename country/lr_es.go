//go:build (lang_es || lang_all) && (country_africa || country_all || country_lr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiberia.RegisterName(xlanguage.Spanish, "Liberia")
	dataLiberia.RegisterOfficialName(xlanguage.Spanish, "República de Liberia")
	dataLiberia.RegisterCapital(xlanguage.Spanish, "Monrovia")
}
