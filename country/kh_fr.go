//go:build (lang_fr || lang_all) && (country_all || country_asia || country_kh || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.French, "Cambodge")
	dataCambodia.RegisterOfficialName(xlanguage.French, "Royaume du Cambodge")
	dataCambodia.RegisterCapital(xlanguage.French, "Phnom Penh")
}
