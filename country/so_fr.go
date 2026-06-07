//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSomalia.RegisterName(xlanguage.French, "Somalie")
	dataSomalia.RegisterOfficialName(xlanguage.French, "République fédérale de Somalie")
	dataSomalia.RegisterCapital(xlanguage.French, "Mogadiscio")
}
