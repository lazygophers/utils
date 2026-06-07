//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.French, "Uruguay")
	dataUruguay.RegisterOfficialName(xlanguage.French, "République orientale de l'Uruguay")
	dataUruguay.RegisterCapital(xlanguage.French, "Montevideo")
}
