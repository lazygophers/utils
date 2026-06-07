//go:build (lang_es || lang_all) && (country_all || country_europe || country_mk || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthMacedonia.RegisterName(xlanguage.Spanish, "Macedonia del Norte")
	dataNorthMacedonia.RegisterOfficialName(xlanguage.Spanish, "República de Macedonia del Norte")
	dataNorthMacedonia.RegisterCapital(xlanguage.Spanish, "Skopie")
}
