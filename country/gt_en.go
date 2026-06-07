//go:build country_all || country_americas || country_central_america || country_gt

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.English, "Guatemala")
	dataGuatemala.RegisterOfficialName(xlanguage.English, "Republic of Guatemala")
	dataGuatemala.RegisterCapital(xlanguage.English, "Guatemala City")
}
