//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsrael.RegisterName(xlanguage.French, "Israël")
	dataIsrael.RegisterOfficialName(xlanguage.French, "État d'Israël")
	dataIsrael.RegisterCapital(xlanguage.French, "Jérusalem")
}
