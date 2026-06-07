//go:build (lang_ar || lang_all) && (country_all || country_eastern_europe || country_europe || country_ro)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.Arabic, "رومانيا")
	dataRomania.RegisterOfficialName(xlanguage.Arabic, "رومانيا")
	dataRomania.RegisterCapital(xlanguage.Arabic, "بوخارست")
}
