//go:build country_all || country_americas || country_gl || country_northern_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreenland.RegisterName(xlanguage.English, "Greenland")
	dataGreenland.RegisterOfficialName(xlanguage.English, "Greenland")
	dataGreenland.RegisterCapital(xlanguage.English, "Nuuk")
}
