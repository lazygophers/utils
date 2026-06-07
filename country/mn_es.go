//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMongolia.RegisterName(xlanguage.Spanish, "Mongolia")
	dataMongolia.RegisterOfficialName(xlanguage.Spanish, "Mongolia")
	dataMongolia.RegisterCapital(xlanguage.Spanish, "Ulán Bator")
}
