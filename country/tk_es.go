//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTokelau.RegisterName(xlanguage.Spanish, "Tokelau")
	dataTokelau.RegisterOfficialName(xlanguage.Spanish, "Tokelau")
	dataTokelau.RegisterCapital(xlanguage.Spanish, "Atafu")
}
