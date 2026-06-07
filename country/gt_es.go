//go:build country_all || country_americas || country_central_america || country_gt

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.Spanish, "Guatemala")
	dataGuatemala.RegisterOfficialName(xlanguage.Spanish, "República de Guatemala")
	dataGuatemala.RegisterCapital(xlanguage.Spanish, "Ciudad de Guatemala")
}
