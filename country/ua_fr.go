//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.French, "Ukraine")
	dataUkraine.RegisterOfficialName(xlanguage.French, "Ukraine")
	dataUkraine.RegisterCapital(xlanguage.French, "Kiev")
}
