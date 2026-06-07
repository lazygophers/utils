//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahamas.RegisterName(xlanguage.French, "Bahamas")
	dataBahamas.RegisterOfficialName(xlanguage.French, "Commonwealth des Bahamas")
	dataBahamas.RegisterCapital(xlanguage.French, "Nassau")
}
