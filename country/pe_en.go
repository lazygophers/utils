//go:build country_all || country_americas || country_pe || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPeru.RegisterName(xlanguage.English, "Peru")
	dataPeru.RegisterOfficialName(xlanguage.English, "Republic of Peru")
	dataPeru.RegisterCapital(xlanguage.English, "Lima")
}
