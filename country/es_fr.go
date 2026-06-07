//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSpain.RegisterName(xlanguage.French, "Espagne")
	dataSpain.RegisterOfficialName(xlanguage.French, "Royaume d'Espagne")
	dataSpain.RegisterCapital(xlanguage.French, "Madrid")
}
