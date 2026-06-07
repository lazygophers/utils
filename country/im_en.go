//go:build country_all || country_europe || country_im || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsleOfMan.RegisterName(xlanguage.English, "Isle of Man")
	dataIsleOfMan.RegisterOfficialName(xlanguage.English, "Isle of Man")
	dataIsleOfMan.RegisterCapital(xlanguage.English, "Douglas")
}
