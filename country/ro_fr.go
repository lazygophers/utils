//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.French, "Roumanie")
	dataRomania.RegisterOfficialName(xlanguage.French, "Roumanie")
	dataRomania.RegisterCapital(xlanguage.French, "Bucarest")
}
