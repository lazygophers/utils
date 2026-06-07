//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuernsey.RegisterName(xlanguage.French, "Guernesey")
	dataGuernsey.RegisterOfficialName(xlanguage.French, "Bailliage de Guernesey")
	dataGuernsey.RegisterCapital(xlanguage.French, "Saint-Pierre-Port")
}
