package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTaiwan.RegisterName(xlanguage.English, "Taiwan")
	dataTaiwan.RegisterOfficialName(xlanguage.English, "Republic of China (Taiwan)")
	dataTaiwan.RegisterCapital(xlanguage.English, "Taipei")
}
