//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaymanIslands.RegisterName(xlanguage.Russian, "Острова Кайман")
	dataCaymanIslands.RegisterOfficialName(xlanguage.Russian, "Острова Кайман")
	dataCaymanIslands.RegisterCapital(xlanguage.Russian, "Джорджтаун")
}
