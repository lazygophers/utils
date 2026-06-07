//go:build country_africa || country_all || country_northern_africa || country_sd

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.English, "Sudan")
	dataSudan.RegisterOfficialName(xlanguage.English, "Republic of the Sudan")
	dataSudan.RegisterCapital(xlanguage.English, "Khartoum")
}
