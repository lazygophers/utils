//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.French, "Estonie")
	dataEstonia.RegisterOfficialName(xlanguage.French, "République d'Estonie")
	dataEstonia.RegisterCapital(xlanguage.French, "Tallinn")
}
