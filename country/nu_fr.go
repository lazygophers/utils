//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiue.RegisterName(xlanguage.French, "Niue")
	dataNiue.RegisterOfficialName(xlanguage.French, "Niue")
	dataNiue.RegisterCapital(xlanguage.French, "Alofi")
}
