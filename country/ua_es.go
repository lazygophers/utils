//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.Spanish, "Ucrania")
	dataUkraine.RegisterOfficialName(xlanguage.Spanish, "Ucrania")
	dataUkraine.RegisterCapital(xlanguage.Spanish, "Kiev")
}
