//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.Russian, "Эритрея")
	dataEritrea.RegisterOfficialName(xlanguage.Russian, "Государство Эритрея")
	dataEritrea.RegisterCapital(xlanguage.Russian, "Асмэра")
}
