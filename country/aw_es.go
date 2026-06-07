//go:build (lang_es || lang_all) && (country_all || country_americas || country_aw || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAruba.RegisterName(xlanguage.Spanish, "Aruba")
	dataAruba.RegisterOfficialName(xlanguage.Spanish, "Aruba")
	dataAruba.RegisterCapital(xlanguage.Spanish, "Oranjestad")
}
