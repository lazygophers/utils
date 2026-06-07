//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKiribati.RegisterName(xlanguage.Arabic, "كيريباتي")
	dataKiribati.RegisterOfficialName(xlanguage.Arabic, "جمهورية كيريباتي")
	dataKiribati.RegisterCapital(xlanguage.Arabic, "تاراوا الجنوبية")
}
