//go:build country_all || country_americas || country_south_america || country_sr

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.English, "Suriname")
	dataSuriname.RegisterOfficialName(xlanguage.English, "Republic of Suriname")
	dataSuriname.RegisterCapital(xlanguage.English, "Paramaribo")
}
