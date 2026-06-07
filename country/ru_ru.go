package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRussia.RegisterName(xlanguage.Russian, "Россия")
	dataRussia.RegisterOfficialName(xlanguage.Russian, "Российская Федерация")
	dataRussia.RegisterCapital(xlanguage.Russian, "Москва")
}
