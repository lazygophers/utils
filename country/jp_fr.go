//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJapan.RegisterName(xlanguage.French, "Japon")
	dataJapan.RegisterOfficialName(xlanguage.French, "Japon")
	dataJapan.RegisterCapital(xlanguage.French, "Tokyo")
}
