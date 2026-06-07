//go:build (lang_fr || lang_all) && (country_all || country_asia || country_il || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsrael.RegisterName(xlanguage.French, "Israël")
	dataIsrael.RegisterOfficialName(xlanguage.French, "État d'Israël")
	dataIsrael.RegisterCapital(xlanguage.French, "Jérusalem")
}
