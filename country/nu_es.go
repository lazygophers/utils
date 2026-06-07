//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiue.RegisterName(xlanguage.Spanish, "Niue")
	dataNiue.RegisterOfficialName(xlanguage.Spanish, "Niue")
	dataNiue.RegisterCapital(xlanguage.Spanish, "Alofi")
}
