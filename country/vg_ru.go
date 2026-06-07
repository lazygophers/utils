//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishVirginIslands.RegisterName(xlanguage.Russian, "Британские Виргинские острова")
	dataBritishVirginIslands.RegisterOfficialName(xlanguage.Russian, "Виргинские острова")
	dataBritishVirginIslands.RegisterCapital(xlanguage.Russian, "Род-Таун")
}
