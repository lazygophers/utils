//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedKingdom.RegisterName(xlanguage.French, "Royaume-Uni")
	dataUnitedKingdom.RegisterOfficialName(xlanguage.French, "Royaume-Uni de Grande-Bretagne et d'Irlande du Nord")
	dataUnitedKingdom.RegisterCapital(xlanguage.French, "Londres")
}
