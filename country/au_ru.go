//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustralia.RegisterName(xlanguage.Russian, "Австралия")
	dataAustralia.RegisterOfficialName(xlanguage.Russian, "Австралийский Союз")
	dataAustralia.RegisterCapital(xlanguage.Russian, "Канберра")
}
