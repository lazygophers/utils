//go:build country_all || country_europe || country_no || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.English, "Norway")
	dataNorway.RegisterOfficialName(xlanguage.English, "Kingdom of Norway")
	dataNorway.RegisterCapital(xlanguage.English, "Oslo")
}
