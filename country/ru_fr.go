//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRussia.RegisterName(xlanguage.French, "Russie")
	dataRussia.RegisterOfficialName(xlanguage.French, "Fédération de Russie")
	dataRussia.RegisterCapital(xlanguage.French, "Moscou")
}
