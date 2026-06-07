//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthMacedonia.RegisterName(xlanguage.Korean, "북마케도니아")
	dataNorthMacedonia.RegisterOfficialName(xlanguage.Korean, "북마케도니아 공화국")
	dataNorthMacedonia.RegisterCapital(xlanguage.Korean, "스코페")
}
