//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKuwait.RegisterName(xlanguage.French, "Koweït")
	dataKuwait.RegisterOfficialName(xlanguage.French, "État du Koweït")
	dataKuwait.RegisterCapital(xlanguage.French, "Koweït")
}
