package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLuxembourg.RegisterName(xlanguage.French, "Luxembourg")
	dataLuxembourg.RegisterOfficialName(xlanguage.French, "Grand-Duché de Luxembourg")
	dataLuxembourg.RegisterCapital(xlanguage.French, "Luxembourg")
}
