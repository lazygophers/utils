//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_et)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEthiopia.RegisterName(xlanguage.French, "Éthiopie")
	dataEthiopia.RegisterOfficialName(xlanguage.French, "République fédérale démocratique d'Éthiopie")
	dataEthiopia.RegisterCapital(xlanguage.French, "Addis-Abeba")
}
