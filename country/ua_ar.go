//go:build (lang_ar || lang_all) && (country_all || country_eastern_europe || country_europe || country_ua)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.Arabic, "أوكرانيا")
	dataUkraine.RegisterOfficialName(xlanguage.Arabic, "أوكرانيا")
	dataUkraine.RegisterCapital(xlanguage.Arabic, "كييف")
}
