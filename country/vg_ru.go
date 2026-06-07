//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_vg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishVirginIslands.RegisterName(xlanguage.Russian, "Британские Виргинские острова")
	dataBritishVirginIslands.RegisterOfficialName(xlanguage.Russian, "Виргинские острова")
	dataBritishVirginIslands.RegisterCapital(xlanguage.Russian, "Род-Таун")
}
