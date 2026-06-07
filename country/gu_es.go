//go:build (lang_es || lang_all) && (country_all || country_gu || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuam.RegisterName(xlanguage.Spanish, "Guam")
	dataGuam.RegisterOfficialName(xlanguage.Spanish, "Territorio de Guam")
	dataGuam.RegisterCapital(xlanguage.Spanish, "Hagåtña")
}
