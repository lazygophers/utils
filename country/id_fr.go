//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndonesia.RegisterName(xlanguage.French, "Indonésie")
	dataIndonesia.RegisterOfficialName(xlanguage.French, "République d'Indonésie")
	dataIndonesia.RegisterCapital(xlanguage.French, "Jakarta")
}
