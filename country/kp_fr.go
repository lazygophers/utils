//go:build (lang_fr || lang_all) && (country_all || country_asia || country_eastern_asia || country_kp)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthKorea.RegisterName(xlanguage.French, "Corée du Nord")
	dataNorthKorea.RegisterOfficialName(xlanguage.French, "République populaire démocratique de Corée")
	dataNorthKorea.RegisterCapital(xlanguage.French, "Pyongyang")
}
