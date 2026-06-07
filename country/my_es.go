//go:build (lang_es || lang_all) && (country_all || country_asia || country_my || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalaysia.RegisterName(xlanguage.Spanish, "Malasia")
	dataMalaysia.RegisterOfficialName(xlanguage.Spanish, "Malasia")
	dataMalaysia.RegisterCapital(xlanguage.Spanish, "Kuala Lumpur")
}
