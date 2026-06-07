//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAndorra.RegisterName(xlanguage.Russian, "Андорра")
	dataAndorra.RegisterOfficialName(xlanguage.Russian, "Княжество Андорра")
	dataAndorra.RegisterCapital(xlanguage.Russian, "Андорра-ла-Велья")
}
