//go:build (lang_fr || lang_all) && (country_all || country_americas || country_caribbean || country_jm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.French, "Jamaïque")
	dataJamaica.RegisterOfficialName(xlanguage.French, "Jamaïque")
	dataJamaica.RegisterCapital(xlanguage.French, "Kingston")
}
