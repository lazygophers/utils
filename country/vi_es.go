//go:build (lang_es || lang_all) && (country_all || country_americas || country_caribbean || country_vi)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsVirginIslands.RegisterName(xlanguage.Spanish, "Islas Vírgenes de los Estados Unidos")
	dataUsVirginIslands.RegisterOfficialName(xlanguage.Spanish, "Islas Vírgenes de los Estados Unidos")
	dataUsVirginIslands.RegisterCapital(xlanguage.Spanish, "Charlotte Amalie")
}
