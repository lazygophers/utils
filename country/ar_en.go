//go:build country_all || country_americas || country_ar || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.English, "Argentina")
	dataArgentina.RegisterOfficialName(xlanguage.English, "Argentine Republic")
	dataArgentina.RegisterCapital(xlanguage.English, "Buenos Aires")
}
