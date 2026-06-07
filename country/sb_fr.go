//go:build (lang_fr || lang_all) && (country_all || country_melanesia || country_oceania || country_sb)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSolomonIslands.RegisterName(xlanguage.French, "Îles Salomon")
	dataSolomonIslands.RegisterOfficialName(xlanguage.French, "Îles Salomon")
	dataSolomonIslands.RegisterCapital(xlanguage.French, "Honiara")
}
