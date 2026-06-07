//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTokelau.RegisterName(xlanguage.French, "Tokelau")
	dataTokelau.RegisterOfficialName(xlanguage.French, "Tokelau")
}
