//go:build country_all || country_americas || country_south_america || country_uy

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.Spanish, "Uruguay")
	dataUruguay.RegisterOfficialName(xlanguage.Spanish, "República Oriental del Uruguay")
	dataUruguay.RegisterCapital(xlanguage.Spanish, "Montevideo")
}
