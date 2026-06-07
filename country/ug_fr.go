//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.French, "Ouganda")
	dataUganda.RegisterOfficialName(xlanguage.French, "République d'Ouganda")
	dataUganda.RegisterCapital(xlanguage.French, "Kampala")
}
