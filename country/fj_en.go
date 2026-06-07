package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFiji.RegisterName(xlanguage.English, "Fiji")
	dataFiji.RegisterOfficialName(xlanguage.English, "Republic of Fiji")
	dataFiji.RegisterCapital(xlanguage.English, "Suva")
}
