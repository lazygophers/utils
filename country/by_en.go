package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelarus.RegisterName(xlanguage.English, "Belarus")
	dataBelarus.RegisterOfficialName(xlanguage.English, "Republic of Belarus")
	dataBelarus.RegisterCapital(xlanguage.English, "Minsk")
}
