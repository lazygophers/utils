//go:build country_all || country_europe || country_mk || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthMacedonia.RegisterName(xlanguage.English, "North Macedonia")
	dataNorthMacedonia.RegisterOfficialName(xlanguage.English, "Republic of North Macedonia")
	dataNorthMacedonia.RegisterCapital(xlanguage.English, "Skopje")
}
