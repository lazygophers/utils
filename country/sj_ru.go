//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSvalbardAndJanMayen.RegisterName(xlanguage.Russian, "Шпицберген и Ян-Майен")
	dataSvalbardAndJanMayen.RegisterOfficialName(xlanguage.Russian, "Шпицберген и Ян-Майен")
	dataSvalbardAndJanMayen.RegisterCapital(xlanguage.Russian, "Лонгйир")
}
