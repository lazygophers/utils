//go:build country_all || country_americas || country_aw || country_caribbean

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAruba.RegisterName(xlanguage.English, "Aruba")
	dataAruba.RegisterOfficialName(xlanguage.English, "Aruba")
	dataAruba.RegisterCapital(xlanguage.English, "Oranjestad")
}
