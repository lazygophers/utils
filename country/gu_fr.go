//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuam.RegisterName(xlanguage.French, "Guam")
	dataGuam.RegisterOfficialName(xlanguage.French, "Guam")
	dataGuam.RegisterCapital(xlanguage.French, "Hagåtña")
}
