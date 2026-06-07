//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.Spanish, "Rumania")
	dataRomania.RegisterOfficialName(xlanguage.Spanish, "Rumania")
	dataRomania.RegisterCapital(xlanguage.Spanish, "Bucarest")
}
