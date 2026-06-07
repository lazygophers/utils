package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.English, "Belize")
	dataBelize.RegisterOfficialName(xlanguage.English, "Belize")
	dataBelize.RegisterCapital(xlanguage.English, "Belmopan")
}
