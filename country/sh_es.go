//go:build (lang_es || lang_all) && (country_africa || country_all || country_sh || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintHelena.RegisterName(xlanguage.Spanish, "Santa Elena, Ascensión y Tristán de Acuña")
	dataSaintHelena.RegisterOfficialName(xlanguage.Spanish, "Santa Elena, Ascensión y Tristán de Acuña")
	dataSaintHelena.RegisterCapital(xlanguage.Spanish, "Jamestown")
}
