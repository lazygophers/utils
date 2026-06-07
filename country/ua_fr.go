//go:build (lang_fr || lang_all) && (country_all || country_eastern_europe || country_europe || country_ua)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.French, "Ukraine")
	dataUkraine.RegisterOfficialName(xlanguage.French, "Ukraine")
	dataUkraine.RegisterCapital(xlanguage.French, "Kiev")
}
