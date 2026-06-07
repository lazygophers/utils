//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishVirginIslands.RegisterName(xlanguage.Spanish, "Islas Vírgenes Británicas")
	dataBritishVirginIslands.RegisterOfficialName(xlanguage.Spanish, "Islas Vírgenes Británicas")
	dataBritishVirginIslands.RegisterCapital(xlanguage.Spanish, "Road Town")
}
