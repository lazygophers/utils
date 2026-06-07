//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMongolia.RegisterName(xlanguage.French, "Mongolie")
	dataMongolia.RegisterOfficialName(xlanguage.French, "Mongolie")
	dataMongolia.RegisterCapital(xlanguage.French, "Oulan-Bator")
}
