//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_ug)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.French, "Ouganda")
	dataUganda.RegisterOfficialName(xlanguage.French, "République d'Ouganda")
	dataUganda.RegisterCapital(xlanguage.French, "Kampala")
}
