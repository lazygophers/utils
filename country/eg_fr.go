//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.French, "Égypte")
	dataEgypt.RegisterOfficialName(xlanguage.French, "République arabe d'Égypte")
	dataEgypt.RegisterCapital(xlanguage.French, "Le Caire")
}
