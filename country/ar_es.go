//go:build country_all || country_americas || country_ar || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.Spanish, "Argentina")
	dataArgentina.RegisterOfficialName(xlanguage.Spanish, "República Argentina")
	dataArgentina.RegisterCapital(xlanguage.Spanish, "Buenos Aires")
}
