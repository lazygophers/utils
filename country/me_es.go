//go:build (lang_es || lang_all) && (country_all || country_europe || country_me || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontenegro.RegisterName(xlanguage.Spanish, "Montenegro")
	dataMontenegro.RegisterOfficialName(xlanguage.Spanish, "Montenegro")
	dataMontenegro.RegisterCapital(xlanguage.Spanish, "Podgorica")
}
