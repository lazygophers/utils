//go:build country_all || country_europe || country_rs || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSerbia.RegisterName(xlanguage.English, "Serbia")
	dataSerbia.RegisterOfficialName(xlanguage.English, "Republic of Serbia")
	dataSerbia.RegisterCapital(xlanguage.English, "Belgrade")
}
