//go:build country_all || country_ba || country_europe || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBosniaAndHerzegovina.RegisterName(xlanguage.English, "Bosnia and Herzegovina")
	dataBosniaAndHerzegovina.RegisterOfficialName(xlanguage.English, "Bosnia and Herzegovina")
	dataBosniaAndHerzegovina.RegisterCapital(xlanguage.English, "Sarajevo")
}
