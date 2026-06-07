package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLuxembourg.RegisterName(xlanguage.English, "Luxembourg")
	dataLuxembourg.RegisterOfficialName(xlanguage.English, "Grand Duchy of Luxembourg")
	dataLuxembourg.RegisterCapital(xlanguage.English, "Luxembourg")
}
