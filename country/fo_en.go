//go:build country_all || country_europe || country_fo || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFaroeIslands.RegisterName(xlanguage.English, "Faroe Islands")
	dataFaroeIslands.RegisterOfficialName(xlanguage.English, "Faroe Islands")
	dataFaroeIslands.RegisterCapital(xlanguage.English, "Torshavn")
}
