//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.Arabic, "أوكرانيا")
	dataUkraine.RegisterOfficialName(xlanguage.Arabic, "أوكرانيا")
	dataUkraine.RegisterCapital(xlanguage.Arabic, "كييف")
}
