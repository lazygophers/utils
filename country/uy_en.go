//go:build country_all || country_americas || country_south_america || country_uy

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.English, "Uruguay")
	dataUruguay.RegisterOfficialName(xlanguage.English, "Oriental Republic of Uruguay")
	dataUruguay.RegisterCapital(xlanguage.English, "Montevideo")
}
