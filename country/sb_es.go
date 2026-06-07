//go:build (lang_es || lang_all) && (country_all || country_melanesia || country_oceania || country_sb)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSolomonIslands.RegisterName(xlanguage.Spanish, "Islas Salomón")
	dataSolomonIslands.RegisterOfficialName(xlanguage.Spanish, "Islas Salomón")
	dataSolomonIslands.RegisterCapital(xlanguage.Spanish, "Honiara")
}
