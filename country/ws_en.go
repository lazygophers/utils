//go:build country_all || country_oceania || country_polynesia || country_ws

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSamoa.RegisterName(xlanguage.English, "Samoa")
	dataSamoa.RegisterOfficialName(xlanguage.English, "Independent State of Samoa")
	dataSamoa.RegisterCapital(xlanguage.English, "Apia")
}
