//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndia.RegisterName(xlanguage.French, "Inde")
	dataIndia.RegisterOfficialName(xlanguage.French, "République de l'Inde")
	dataIndia.RegisterCapital(xlanguage.French, "New Delhi")
}
