//go:build country_all || country_melanesia || country_oceania || country_sb

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSolomonIslands.RegisterName(xlanguage.English, "Solomon Islands")
	dataSolomonIslands.RegisterOfficialName(xlanguage.English, "Solomon Islands")
	dataSolomonIslands.RegisterCapital(xlanguage.English, "Honiara")
}
