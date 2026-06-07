//go:build (lang_fr || lang_all) && (country_all || country_europe || country_pt || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPortugal.RegisterName(xlanguage.French, "Portugal")
	dataPortugal.RegisterOfficialName(xlanguage.French, "République portugaise")
	dataPortugal.RegisterCapital(xlanguage.French, "Lisbonne")
}
