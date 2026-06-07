package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelarus.RegisterName(xlanguage.Russian, "Беларусь")
	dataBelarus.RegisterOfficialName(xlanguage.Russian, "Республика Беларусь")
	dataBelarus.RegisterCapital(xlanguage.Russian, "Минск")
}
