//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKuwait.RegisterName(xlanguage.Spanish, "Kuwait")
	dataKuwait.RegisterOfficialName(xlanguage.Spanish, "Estado de Kuwait")
	dataKuwait.RegisterCapital(xlanguage.Spanish, "Kuwait")
}
