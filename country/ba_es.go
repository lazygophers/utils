//go:build (lang_es || lang_all) && (country_all || country_ba || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBosniaAndHerzegovina.RegisterName(xlanguage.Spanish, "Bosnia y Herzegovina")
	dataBosniaAndHerzegovina.RegisterOfficialName(xlanguage.Spanish, "Bosnia y Herzegovina")
	dataBosniaAndHerzegovina.RegisterCapital(xlanguage.Spanish, "Sarajevo")
}
