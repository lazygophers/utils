//go:build (lang_es || lang_all) && (country_all || country_eastern_europe || country_europe || country_hu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.Spanish, "Hungría")
	dataHungary.RegisterOfficialName(xlanguage.Spanish, "Hungría")
	dataHungary.RegisterCapital(xlanguage.Spanish, "Budapest")
}
