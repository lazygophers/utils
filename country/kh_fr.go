//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.French, "Cambodge")
	dataCambodia.RegisterOfficialName(xlanguage.French, "Royaume du Cambodge")
	dataCambodia.RegisterCapital(xlanguage.French, "Phnom Penh")
}
