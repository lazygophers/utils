//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.Arabic, "جامايكا")
	dataJamaica.RegisterOfficialName(xlanguage.Arabic, "جامايكا")
	dataJamaica.RegisterCapital(xlanguage.Arabic, "كينغستون")
}
