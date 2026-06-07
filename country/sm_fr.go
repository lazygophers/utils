//go:build (lang_fr || lang_all) && (country_all || country_europe || country_sm || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.French, "Saint-Marin")
	dataSanMarino.RegisterOfficialName(xlanguage.French, "République de Saint-Marin")
	dataSanMarino.RegisterCapital(xlanguage.French, "Saint-Marin")
}
