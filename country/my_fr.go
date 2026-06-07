//go:build (lang_fr || lang_all) && (country_all || country_asia || country_my || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalaysia.RegisterName(xlanguage.French, "Malaisie")
	dataMalaysia.RegisterOfficialName(xlanguage.French, "Malaisie")
	dataMalaysia.RegisterCapital(xlanguage.French, "Kuala Lumpur")
}
