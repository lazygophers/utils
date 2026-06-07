//go:build country_all || country_antarctic || country_gs

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthGeorgiaAndSouthSandwich.RegisterName(xlanguage.English, "South Georgia and the South Sandwich Islands")
	dataSouthGeorgiaAndSouthSandwich.RegisterOfficialName(xlanguage.English, "South Georgia and the South Sandwich Islands")
	dataSouthGeorgiaAndSouthSandwich.RegisterCapital(xlanguage.English, "King Edward Point")
}
