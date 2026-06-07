//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintHelena.RegisterName(xlanguage.Spanish, "Santa Elena, Ascensión y Tristán de Acuña")
	dataSaintHelena.RegisterOfficialName(xlanguage.Spanish, "Santa Elena, Ascensión y Tristán de Acuña")
	dataSaintHelena.RegisterCapital(xlanguage.Spanish, "Jamestown")
}
