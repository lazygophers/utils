//go:build country_all || country_americas || country_cl || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChile.RegisterName(xlanguage.English, "Chile")
	dataChile.RegisterOfficialName(xlanguage.English, "Republic of Chile")
	dataChile.RegisterCapital(xlanguage.English, "Santiago")
}
