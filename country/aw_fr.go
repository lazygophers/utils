//go:build (lang_fr || lang_all) && (country_all || country_americas || country_aw || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAruba.RegisterName(xlanguage.French, "Aruba")
	dataAruba.RegisterOfficialName(xlanguage.French, "Aruba")
	dataAruba.RegisterCapital(xlanguage.French, "Oranjestad")
}
