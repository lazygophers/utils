//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.French, "Érythrée")
	dataEritrea.RegisterOfficialName(xlanguage.French, "État d'Érythrée")
	dataEritrea.RegisterCapital(xlanguage.French, "Asmara")
}
